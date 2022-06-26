package utilities

import (
	"context"
	"testing"
)

func Test_Set_Build_status(t *testing.T) {
	buildID := "coolbuild"
	SetBuildStatus(buildID, 0)

	if GetBuildStatus(buildID) != "running" {
		t.Fatal("Failed: Build Status set incorrectly")
	}

	SetBuildStatus(buildID, 2)

	if GetBuildStatus(buildID) != "failed" {
		t.Fatal("Failed: Build Status set incorrectly")
	}

	SetBuildStatus(buildID, 1)

	if GetBuildStatus(buildID) != "success" {
		t.Fatal("Failed: Build Status set incorrectly")
	}
}

func Test_Get_Build_Path(t *testing.T) {
	buildID := "coolbuild"
	if GetBuildPath(buildID) != "./data/"+buildID {
		t.Fatal("Failed: Build Path set incorrectly")
	}
}

func Test_Get_Docker_client(t *testing.T) {

	_, err := GetDockerClient(context.Background())
	if err != nil {
		t.Fatal("failed: ", err)
	}
}
