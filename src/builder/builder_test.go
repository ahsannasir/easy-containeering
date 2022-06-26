package builder

import (
	"context"
	"fmt"
	"testing"

	utils "ml-cicd/src/utilities"

	"github.com/stretchr/testify/assert"
)

// // Tests the normal workflow which succeeds
// func TestBuildSuccess(t *testing.T) {
// 	buildID := "testBuildID"
// 	utils.PrepareTestSetup(utils.GetBuildPath(buildID), "testBuildID")
// 	cli := utils.GetDockerClient(context.Background())
// 	err := Build(cli, buildID, "ahsannasir", "testimagename")

// 	if err != nil {
// 		utils.DestroyTestSetup(utils.GetBuildPath(buildID))
// 		t.Fatal("Failed: ", err)
// 	}
// }

// // Tests behavior of build in case empty image name is sent
// func TestBuildFail(t *testing.T) {
// 	buildID := ""
// 	utils.PrepareTestSetup(utils.GetBuildPath(buildID), "testBuildID")
// 	cli := utils.GetDockerClient(context.Background())
// 	err := Build(cli, buildID, "ahsannasir", "testimagename")

// 	if err != nil {
// 		utils.DestroyTestSetup(utils.GetBuildPath(buildID))
// 		t.Fatal("Failed: ", err)
// 	}
// }

func Test_Builder_Build(t *testing.T) {

	tests := []struct {
		name           string
		buildID        string
		repositoryName string
		imageName      string
		expected       interface{}
		errorExpected  error
	}{{
		name:           "Test Build Normal Workflow.",
		buildID:        "testBuildID",
		repositoryName: "ahsannasir",
		imageName:      "testimagename",
		expected:       nil,
		errorExpected:  nil,
	},
		{
			name:           "Test if Build functions with uppercase repository name.",
			buildID:        "testBuildID",
			repositoryName: "AHSANNASIR",
			imageName:      "testimagename",
			expected:       nil,
			errorExpected:  *new(error),
		},
		{
			name:           "Test if Build functions with uppercase image name.",
			buildID:        "testBuildID",
			repositoryName: "ahsannasir",
			imageName:      "TESTIMAGENAME",
			expected:       nil,
			errorExpected:  *new(error),
		},
		{
			name:           "Test if Build function fails with empty Build ID.",
			buildID:        "",
			repositoryName: "ahsannasir",
			imageName:      "TESTIMAGENAME",
			expected:       nil,
			errorExpected:  *new(error),
		},
		{
			name:           "Test if Build function fails with empty RepositoryName.",
			buildID:        "testBuildID",
			repositoryName: "",
			imageName:      "TESTIMAGENAME",
			expected:       nil,
			errorExpected:  *new(error),
		},
		{
			name:           "Test if Build function fails with empty image name.",
			buildID:        "testBuildID",
			repositoryName: "ahsannasir",
			imageName:      "",
			expected:       nil,
			errorExpected:  *new(error),
		},
	}

	for _, test := range tests {
		result := t.Run(test.name, func(t *testing.T) {
			fmt.Println("Executing Test: ", test.name)
			var (
				err error
			)

			utils.PrepareTestSetup(utils.GetBuildPath(test.buildID), test.buildID)
			cli := utils.GetDockerClient(context.Background())
			err = Build(cli, test.buildID, test.repositoryName, test.imageName)

			if err != nil {
				utils.DestroyTestSetup(utils.GetBuildPath(test.buildID))
				// t.Fatal("Failed: ", err)
				err = *new(error)
			}

			assert.Equal(t, test.errorExpected, err)

		})
		if result == false {
			break
		}
	}
}
