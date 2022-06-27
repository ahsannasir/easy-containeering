package registry

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"ml-cicd/src/artifacts"
	"time"

	utils "ml-cicd/src/utilities"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var authConfig types.AuthConfig

func SetRegistryAuth(username string, password string) {
	authConfig = types.AuthConfig{
		Username:      username,
		Password:      password,
		ServerAddress: "https://hub.docker.com/",
	}
}

// ImagePush: pushes an image build to the repository defined by the user
func ImagePush(dockerClient *client.Client, registryUserID string, imagename string, buildID string) error {

	if dockerClient == nil {
		utils.SetBuildStatus(buildID, 2)
		return errors.New("Invalid Client Received!")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	// prepare auth strings to authenticate service to docker repository
	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	// define tags
	tag := registryUserID + "/" + imagename

	// define push options by encapsulating credentials
	opts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}

	// push image to repository
	rd, err := dockerClient.ImagePush(ctx, tag, opts)
	if err != nil {
		utils.SetBuildStatus(buildID, 2)
		return err
	}

	defer rd.Close()
	// maintain logs for image push operation
	err = artifacts.GenLog(rd, utils.GetBuildPath(buildID)+"/"+buildID)
	if err != nil {
		utils.SetBuildStatus(buildID, 2)
		return err
	}

	return nil
}
