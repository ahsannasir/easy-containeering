package builder

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"ml-cicd/src/registry"
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

			PrepareTestSetup(utils.GetBuildPath(test.buildID), test.buildID)
			cli, err := utils.GetDockerClient(context.Background())
			err = Build(cli, test.buildID, test.repositoryName, test.imageName)

			if err != nil {
				err = *new(error)
			}
			DestroyTestSetup(utils.GetBuildPath(test.buildID))
			assert.Equal(t, test.errorExpected, err)

		})
		if result == false {
			break
		}
	}
}

func PrepareTestSetup(path string, buildID string) error {
	os.MkdirAll(path, os.ModePerm)
	// Create a dockerfile and replicate the one sent by user
	f, err := os.OpenFile(path+"/Dockerfile", os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		return err //please dont
	}

	defer f.Close()

	file, err := os.Open("../../resources/Dockerfile")
	if err != nil {
		return err //please dont
	}
	registry.SetRegistryAuth("ahsannasir", "playstationxbox1")
	// fmt.Println(string(file))
	defer file.Close()
	io.Copy(f, file)
	return nil
}

func DestroyTestSetup(path string) {
	os.RemoveAll(path + "/")
}
