package artifacts

import (
	"bytes"
	"io/ioutil"
	"log"
	utils "ml-cicd/src/utilities"
	"os"
	"strings"
	"testing"
)

func Test_Artifacts_Gen(t *testing.T) {

	file, err := os.Open("../../resources/Dockerfile")
	if err != nil {
		t.Fatal("fail")
	}
	defer file.Close()

	err = GenArtifacts(file, "Dockerfile", "testBuildId", false)
	if err != nil {
		t.Fatal("Failed: ", err)
	}
	DestroyTestSetup(utils.GetBuildPath("testBuildId"))
}

func Test_Artifacts_Gen_SameContent(t *testing.T) {

	file, err := os.Open("../../resources/Dockerfile")
	if err != nil {
		t.Fatal("fail")
	}
	defer file.Close()

	err = GenArtifacts(file, "Dockerfile", "testBuildId", false)
	if err != nil {
		t.Fatal("Failed: ", err)
	}
	f1, err := ioutil.ReadFile("../../resources/Dockerfile")
	f2, err := ioutil.ReadFile(utils.GetBuildPath("testBuildId") + "/Dockerfile")

	if err != nil {
		log.Fatal(err)
	}

	if bytes.Equal(f1, f2) == false {
		t.Fatal("Failed: Different files were generated")
	}
	DestroyTestSetup(utils.GetBuildPath("testBuildId"))
}
func Test_Artifacts_Gen_Log(t *testing.T) {
	err := GenLog(strings.NewReader("hey cool logs!"), utils.GetBuildPath("testBuildId"))
	if err != nil {
		t.Fatal("Failed: ", err)
	}
	DestroyTestSetup(utils.GetBuildPath("testBuildId"))
}

func Test_Artifacts_Fetch_Log(t *testing.T) {
	buildID := "testBuildId"
	os.MkdirAll(utils.GetBuildPath(buildID)+"/"+buildID, os.ModePerm)
	err := GenLog(strings.NewReader("hey cool logs!"), utils.GetBuildPath(buildID)+"/"+buildID)
	if err != nil {
		t.Fatal("Failed: ", err)
	}
	logs, err := FetchLog(buildID)
	if err != nil && logs == "" {
		t.Fatal("Failed: ", err)
	}
	if !strings.Contains(logs, "hey cool logs!") {
		t.Fatal("Failed: Unequal strings!")
	}
	DestroyTestSetup(utils.GetBuildPath(buildID) + "/" + buildID)
}

func DestroyTestSetup(path string) {
	os.RemoveAll(path + "/")
}
