package builder

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	artifacts "ml-cicd/src/artifacts"
	registry "ml-cicd/src/registry"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/stdcopy"
)

var builds = map[string]string{}

func Build(buildID string) error {
	ctx := context.Background()
	cli := getClient(ctx)

	builds[buildID] = "running"

	tar, err := archive.TarWithOptions("./data/"+buildID+"/", &archive.TarOptions{})
	if err != nil {
		panic(err)
	}
	// io.Copy(os.Stdout, tar)

	opts := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{"ahsannasir" + "/heycoolimage"},
		Remove:     false,
	}
	res, err := cli.ImageBuild(ctx, tar, opts)
	if err != nil {
		builds[buildID] = "failed"
		panic(err)
	}

	defer res.Body.Close()
	stdcopy.StdCopy(os.Stdout, os.Stderr, res.Body)
	err = print(res.Body, buildID)
	if err != nil {
		panic(err)
	}

	err = registry.ImagePush(cli)
	if err != nil {
		builds[buildID] = "failed"
		panic(err)
	}

	builds[buildID] = "success"

	return err
}

func Status(buildID string) string {
	return builds[buildID]
}

func getClient(ctx context.Context) *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return cli
}

func print(rd io.Reader, imageName string) error {
	var lastLine string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()
		fmt.Println(scanner.Text())
		artifacts.Log(imageName, scanner.Text()+"\n")
	}

	errLine := &ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

type ErrorLine struct {
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}
