package utilities

import (
	"io"
	"os"
)

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
	// fmt.Println(string(file))
	defer file.Close()
	io.Copy(f, file)
	return nil
}

func DestroyTestSetup(path string) {
	os.RemoveAll(path)
}
