package artifacts

import (
	"io"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request, buildID string) error {
	file, handler, err := r.FormFile("Dockerfile")
	defer file.Close()

	if err != nil {
		panic(err) //dont do this
	}

	// Create a Directory for build
	os.MkdirAll("./data/"+buildID, os.ModePerm)

	// Create a dockerfile and replicate the one sent by user
	f, err := os.OpenFile("./data/"+buildID+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		panic(err) //please dont
	}
	defer f.Close()
	io.Copy(f, file)
	return err
}

func Log(filename string, text string) {
	f, err := os.OpenFile(filename+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func Getlog(buildID string) string {
	dat, err := os.ReadFile(buildID + ".txt")
	if err != nil {
		panic(err)
	}
	return string(dat)
}
