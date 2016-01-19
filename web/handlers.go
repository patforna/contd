package web

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/patforna/splendid/models"
	"github.com/patforna/splendid/runner"
)

// $ curl -v localhost:8080/
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi Pat!")
}

// $ curl -v -X POST localhost:8080/builds
func NewBuild(w http.ResponseWriter, r *http.Request) {
	build := models.NewBuild()
	models.SaveBuild(build)

	pipeline := models.GetPipeline()

	runner := runner.Runner{
		Image: pipeline.Image,
		Command: pipeline.Command,
		InputDir: "/tmp/input/x",
		OutputDir: "/tmp/output/x",
	}

	//status := runner.Run()
	runner.Run()


	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(build)
}

// $ curl -v localhost:8080/build/:id
func Build(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	build := models.GetBuild(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(build)
}