package builder

import (
	"context"
	"os"

	artifacts "ml-cicd/src/artifacts"
	registry "ml-cicd/src/registry"
	utils "ml-cicd/src/utilities"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/stdcopy"
)

// Build: Initiates a docker image build
func Build(buildID string, repository string, imagename string) error {
	ctx := context.Background()

	// get a client to docker daemon
	cli := utils.GetDockerClient(ctx)

	// maintain build status 0 means "running"
	utils.SetBuildStatus(buildID, 0)

	// create a tar of the files submitted to further create an image out of it
	tar, err := archive.TarWithOptions("./data/"+buildID+"/", &archive.TarOptions{})
	if err != nil {
		panic(err)
	}
	// io.Copy(os.Stdout, tar)
	// provider build options, image details
	opts := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{repository + "/" + imagename},
		Remove:     false,
	}

	// finally build our image
	// equivalent to "docker build -t {{image_name}} ."
	res, err := cli.ImageBuild(ctx, tar, opts)
	if err != nil {
		utils.SetBuildStatus(buildID, 2)
		panic(err)
	}

	defer res.Body.Close()

	// Route build output
	stdcopy.StdCopy(os.Stdout, os.Stderr, res.Body)
	err = artifacts.GenLog(res.Body, utils.GetBuildPath(buildID)+"/"+buildID)
	if err != nil {
		panic(err)
	}

	// Finally push this image to the docker repository configured by the user
	err = registry.ImagePush(cli, repository, imagename, buildID)
	if err != nil {
		utils.SetBuildStatus(buildID, 2)
		panic(err)
	}

	// set build status to "success" of everthing went as expected
	utils.SetBuildStatus(buildID, 1)

	return err
}
