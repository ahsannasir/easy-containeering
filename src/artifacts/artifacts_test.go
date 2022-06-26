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

	err = GenArtifacts(file, "Dockerfile", "testBuildId")
	if err != nil {
		t.Fatal("Failed: ", err)
	}
	utils.DestroyTestSetup(utils.GetBuildPath("testBuildId"))
}

func Test_Artifacts_Gen_SameContent(t *testing.T) {

	file, err := os.Open("../../resources/Dockerfile")
	if err != nil {
		t.Fatal("fail")
	}
	defer file.Close()

	err = GenArtifacts(file, "Dockerfile", "testBuildId")
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
	utils.DestroyTestSetup(utils.GetBuildPath("testBuildId"))
}
func Test_Artifacts_Gen_Log(t *testing.T) {
	err := GenLog(strings.NewReader("hey cool logs!"), utils.GetBuildPath("testBuildId"))
	if err != nil {
		t.Fatal("Failed: ", err)
	}
	utils.DestroyTestSetup(utils.GetBuildPath("testBuildId"))
}

func Test_Artifacts_Fetch_Log(t *testing.T) {
	err := GenLog(strings.NewReader("hey cool logs!"), utils.GetBuildPath("testBuildId"))
	if err != nil {
		t.Fatal("Failed: ", err)
	}
	logs, err := FetchLog("testBuildId")
	if err != nil && logs == "" {
		t.Fatal("Failed: ", err)
	}
	if !strings.Contains(logs, "hey cool logs!") {
		t.Fatal("Failed: Unequal strings!")
	}
	utils.DestroyTestSetup(utils.GetBuildPath("testBuildId"))
}
