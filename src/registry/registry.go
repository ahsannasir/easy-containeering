package registry

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"ml-cicd/src/artifacts"
	"time"

	utils "ml-cicd/src/utilities"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var authConfig = types.AuthConfig{
	Username:      "ahsannasir",
	Password:      "playstationxbox1",
	ServerAddress: "https://hub.docker.com/",
}

func ImagePush(dockerClient *client.Client, dockerRegistryUserID string, imagename string, buildID string) error {
	// dockerRegistryUserID := "ahsannasir"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	tag := dockerRegistryUserID + "/" + imagename
	opts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}
	rd, err := dockerClient.ImagePush(ctx, tag, opts)
	if err != nil {
		return err
	}

	defer rd.Close()

	err = artifacts.GenLog(rd, utils.GetBuildPath(buildID)+"/"+buildID)
	if err != nil {
		return err
	}

	return nil
}
