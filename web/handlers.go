package web

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/patforna/contd/models"
	"github.com/patforna/contd/runner"
	"gopkg.in/olivere/elastic.v3"
"github.com/patforna/contd/models"
)

// $ curl -v localhost:8080/
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi Pat!")
}

/* $ curl -v -X POST localhost:8080/builds -d '
     { image: ..., command: ... }'
*/
func NewPipeline(w http.ResponseWriter, r *http.Request) {
	// FIXME read image/command from request
	pipeline := models.NewPipeline("java:8", "javac -verbose Hello.java")
	models.SavePipeline(pipeline)
}

// $ curl -v localhost:8080/pipeline/:id
func Build(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	build := models.GetPipeline(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(build)
}

/* $ curl -v -X POST localhost:8080/builds -d '
     { pipeline_id: ... }'
*/
func NewBuild(w http.ResponseWriter, r *http.Request) {
	build := models.NewBuild()
	models.SaveBuild(build)

	// FIXME read pipeline_id from request
	pipelineId := "x"
	pipeline := models.GetPipeline(pipelineId)

	// FIXME use: docker inspect -f '{{ range .Mounts }}{{ if eq .Destination "/input" }}{{ .Source }}{{ end }}{{ end }}' data
	inputDir := "/mnt/sda1/var/lib/docker/volumes/ca893c1973fff5361b2be780c786608abf1d374d7a11e21a4d0f4509d02c7629/_data"
	// FIXME use: docker inspect -f '{{ range .Mounts }}{{ if eq .Destination "/output" }}{{ .Source }}{{ end }}{{ end }}' data
	outputDir := "/mnt/sda1/var/lib/docker/volumes/1e2a78a33d8dd3e86592f1aa3cff5581b83f284165fb30fe063107e2a08a85e1/_data"

	runner := runner.Runner{
		Image: pipeline.Image,
		Command: pipeline.Command,
		InputDir: inputDir,
		OutputDir:fmt.Sprintf("%s/%s", outputDir, build.Id),
	}

	containerId := runner.Run()
	models.SaveContainerId(build.Id, containerId)

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

// $ curl -v localhost:8080/log/:build_id
func Output(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//containerId := models.GetContainerId(vars["build_id"])


	client, err := elastic.NewClient(
		elastic.SetURL("http://188.166.136.103:9200"),
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		panic(err)
	}

	termQuery := elastic.NewTermQuery("docker.cid", "49a06358b246")
	searchResult, err := client.Search().
	//Index("twitter").   // search in index "twitter"
	Query(termQuery).// specify the query
	Sort("@timestamp", true). // sort by "user" field, ascending
	From(0).Size(100).
	Do()
	if err != nil {
		panic(err)
	}

	if searchResult.Hits != nil {
		fmt.Fprintf(w, "Found a total of %d messages\n", searchResult.Hits.TotalHits)

		for _, hit := range searchResult.Hits.Hits {
			var m map[string]interface{}
			err := json.Unmarshal(*hit.Source, &m)
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, "Message: %s\n", m["message"])
		}
	} else {
		// No hits
		fmt.Fprintf(w, "Found no messages\n")
	}

	//fmt.Fprintf(w, "Output for container: %s", "xxx")
	//	fmt.Fprintf(w, "Output for container: %s", containerId)
}
