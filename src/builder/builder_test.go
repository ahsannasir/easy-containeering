package builder

import (
	"context"
	"testing"

	utils "ml-cicd/src/utilities"
)

func TestConfigCreateError(t *testing.T) {
	buildID := "testBuildID"
	utils.PrepareTestSetup(utils.GetBuildPath(buildID), "testBuildID")
	cli := utils.GetDockerClient(context.Background())
	err := Build(cli, "testBuildID", "ahsannasir", "testimagename")

	if err != nil {
		utils.DestroyTestSetup("testBuildID")
		t.Fatal("Failed: ", err)
	}
}
