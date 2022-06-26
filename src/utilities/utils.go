package utilities

import (
	"context"

	"github.com/docker/docker/client"
)

var builds = map[string]string{}

func SetBuildStatus(buildID string, flag int) {
	switch flag {
	case 0:
		builds[buildID] = "running"
	case 1:
		builds[buildID] = "success"
	case 2:
		builds[buildID] = "failed"
	}
}

func GetBuildStatus(buildID string) string {
	return builds[buildID]
}

func GetBuildPath(buildID string) string {
	return "./data/" + buildID
}

func GetDockerClient(ctx context.Context) (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return cli, nil
}
