package registry

import (
	"context"
	"os"
	"testing"

	utils "ml-cicd/src/utilities"
)

// func Test_ImagePush_Gen(t *testing.T) {
// 	cli := utils.GetDockerClient(context.Background())

// 	repository := "ahsannasir"
// 	imagename := "heycoolimage"
// 	buildID := "testbuildid"
// 	utils.PrepareTestSetup(utils.GetBuildPath(buildID), buildID)

// 	err := Build(cli, buildID, repository, imagename)

// 	os.MkdirAll(utils.GetBuildPath(buildID)+"/"+buildID, os.ModePerm)
// 	// Finally push this image to the docker repository configured by the user
// 	err = ImagePush(cli, repository, imagename, buildID)
// 	if err != nil {
// 		t.Fatal("Failed: ", err)
// 	}

// }

func Test_ImagePush_Gen_Not_Exists(t *testing.T) {
	cli, _ := utils.GetDockerClient(context.Background())

	repository := "ahsannasir"
	imagename := "notexistsimage"
	buildID := "testbuildid"
	os.MkdirAll(utils.GetBuildPath(buildID)+"/"+buildID, os.ModePerm)
	// Finally push this image to the docker repository configured by the user
	err := ImagePush(cli, repository, imagename, buildID)
	if err == nil {
		t.Fatal("Failed: ", err)
	}

}

func Test_ImagePush_Gen_Wrong_Repo(t *testing.T) {
	cli, _ := utils.GetDockerClient(context.Background())
	repository := "CAPITALLETTERS"
	imagename := "heycoolimage"
	buildID := "testbuildid"
	// Finally push this image to the docker repository configured by the user
	err := ImagePush(cli, repository, imagename, buildID)
	if err == nil {
		t.Fatal("Test Failed: Image shouldn't be pushed with uppercase repository names.")
	}
}

func Test_ImagePush_Gen_Wrong_Image(t *testing.T) {
	cli, _ := utils.GetDockerClient(context.Background())
	repository := "ahsannasir"
	imagename := "HEYCOOLIMAGE"
	buildID := "testbuildid"
	// Finally push this image to the docker repository configured by the user
	err := ImagePush(cli, repository, imagename, buildID)
	if err == nil {
		t.Fatal("Test Failed: Image shouldn't be pushed with uppercase image names.")
	}
}

func Test_ImagePush_Wrong_client(t *testing.T) {
	repository := "ahsannasir"
	imagename := "HEYCOOLIMAGE"
	buildID := "testbuildid"
	// Finally push this image to the docker repository configured by the user
	err := ImagePush(nil, repository, imagename, buildID)
	if err == nil {
		t.Fatal("Test Failed: Client should be valid!")
	}
}
