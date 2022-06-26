package main

// “In your company there are many data scientists, who train deep learning models. They use different ML frameworks such as TensorFlow, PyTorch, scikit-learn, etc.
// The data scientists need to jointly use and share the limited ML-accelerated hardware such as GPUs. Containers seem like an awesome way to encapsulate their custom
// model training scripts and their library dependencies. The data scientists are comfortable working with Docker containers. You need to design and implement a simple
// RESTful service that accepts a Dockerfile from the user. The service builds a Docker image out of the Dockerfile and pushes it to a Docker registry. Note that the
// build might take some time. For the purpose of this assignment, it can be assumed that the service and Docker daemon run on the same machine.”
import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Welcome to the ML CICD Service!")

	// publishes a build job
	http.HandleFunc("/api/publish", publisher)

	// returns build logs of a completed or a running build
	http.HandleFunc("/api/logs", getlogger)

	// returns build statuses
	http.HandleFunc("/api/status", getstatus)
	initiateServer()
}

// initiateServer: Initiates the server
func initiateServer() {
	err := http.ListenAndServe(":5433", logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

// logRequest: logout API calls to the stdout
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		log.Println()
		handler.ServeHTTP(w, r)
	})
}
