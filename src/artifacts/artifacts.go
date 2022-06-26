package artifacts

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"

	er "ml-cicd/api/types"
	utils "ml-cicd/src/utilities"
)

// GenArtifacts: This function generates a file adapted to custom mechanisms
func GenArtifacts(file io.Reader, filename string, buildID string) error {

	// Create a Directory for build if it doesn't exist
	os.MkdirAll(utils.GetBuildPath(buildID), os.ModePerm)

	// Create a dockerfile and replicate the one sent by user
	f, err := os.OpenFile(utils.GetBuildPath(buildID)+"/"+filename, os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		return err //please dont
	}

	defer f.Close()
	io.Copy(f, file)
	return err
}

// GetLog: This function allows reading logs for several running or complete builds
func GenLog(rd io.Reader, path string) error {
	var lastLine string

	scanner := bufio.NewScanner(rd)
	f, err := os.OpenFile(path+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	for scanner.Scan() {
		lastLine = scanner.Text()

		if _, err = f.WriteString(scanner.Text() + "\n"); err != nil {
			return err
		}
	}

	errLine := &er.ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func FetchLog(buildID string) (string, error) {
	dat, err := os.ReadFile(utils.GetBuildPath(buildID) + ".txt")
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
