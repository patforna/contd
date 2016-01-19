package models

var builds []*Build

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

func GetPipeline() *Pipeline {
	// hardcoded for now
	return NewPipeline()
}