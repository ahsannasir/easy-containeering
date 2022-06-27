package builder

import (
	"context"
	"fmt"
	"testing"

	utils "ml-cicd/src/utilities"

	"github.com/stretchr/testify/assert"
)

// Test_Builder_Build: Intelligently creates different possible inputs for build function
// and tests by passing in expected output and error values on different use cases.
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
			cli, err := utils.GetDockerClient(context.Background())
			err = Build(cli, test.buildID, test.repositoryName, test.imageName)

			if err != nil {
				err = *new(error)
			}
			utils.DestroyTestSetup(utils.GetBuildPath(test.buildID))
			assert.Equal(t, test.errorExpected, err)

		})
		if result == false {
			break
		}
	}
}
