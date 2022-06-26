package registry

import (
	"context"
	"testing"

	utils "ml-cicd/src/utilities"
)

func Test_ImagePush_Gen(t *testing.T) {
	cli := utils.GetDockerClient(context.Background())
	repository := "ahsannasir"
	imagename := "heycoolimage"
	buildID := "testbuildid"
	// Finally push this image to the docker repository configured by the user
	err := ImagePush(cli, repository, imagename, buildID)
	if err != nil {
		t.Fatal("Failed: ", err)
	}

}

func Test_ImagePush_Gen_wrong_repo(t *testing.T) {
	cli := utils.GetDockerClient(context.Background())
	repository := "CAPITALLETTERS"
	imagename := "heycoolimage"
	buildID := "testbuildid"
	// Finally push this image to the docker repository configured by the user
	err := ImagePush(cli, repository, imagename, buildID)
	if err != nil {
		t.Fatal("Failed: ", err)
	}
}
