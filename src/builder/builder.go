package builder

import (
	"context"
	"errors"
	"fmt"
	"os"

	artifacts "ml-cicd/src/artifacts"
	registry "ml-cicd/src/registry"
	utils "ml-cicd/src/utilities"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/stdcopy"
)

// Build: Initiates a docker image build
func Build(cli *client.Client, buildID string, repository string, imagename string) error {

	if buildID == "" {
		return errors.New("No valid build ID provided")
	}
	if repository == "" {
		return errors.New("No valid repository ID provided")
	}
	if imagename == "" {
		return errors.New("No valid imagename ID provided")
	}

	ctx := context.Background()
	// maintain build status 0 means "running"
	utils.SetBuildStatus(buildID, 0)

	// create a tar of the files submitted to further create an image out of it
	tar, err := archive.TarWithOptions(utils.GetBuildPath(buildID)+"/", &archive.TarOptions{})
	if err != nil {

		return err
	}
	fmt.Println(utils.GetBuildPath(buildID))
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
		return err
	}

	defer res.Body.Close()

	// Route build output
	stdcopy.StdCopy(os.Stdout, os.Stderr, res.Body)
	err = artifacts.GenLog(res.Body, utils.GetBuildPath(buildID)+"/"+buildID)
	if err != nil {
		return err
	}

	// Finally push this image to the docker repository configured by the user
	err = registry.ImagePush(cli, repository, imagename, buildID)
	if err != nil {
		utils.SetBuildStatus(buildID, 2)
		return err
	}

	// set build status to "success" of everthing went as expected
	utils.SetBuildStatus(buildID, 1)

	return err
}
