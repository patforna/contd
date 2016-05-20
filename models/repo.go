package models

var pipelines []*Pipeline
var builds []*Build
var containerIds = make(map[string]string)

func SavePipeline(p *Pipeline) {
	pipelines = append(pipelines, p)
}

func GetPipeline(id string) *Pipeline {
	for _, b := range pipelines {
		if b.Id == id {
			return b
		}
	}
	panic("not found")
}

func SaveBuild(b *Build) {
	builds = append(builds, b)
}

func GetBuild(id string) *Build {
	for _, b := range builds {
		if b.Id == id {
			return b
		}
	}
	panic("not found")
}

func SaveContainerId(buildId string, containerId string) {
	containerIds[buildId] = containerId
}

func GetContainerId(buildId string) string {
	return containerIds[buildId]
}