package shipper

import (
	"fmt"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

// Use to run docker containers.
type Shipper struct {
	Image   string
	Command string
	InputDir string
	OutputDir  string
}

func (s *Shipper) Run() int {
	//endpoint := "unix:///var/run/docker.sock"
	//client, _ := docker.NewClient(endpoint)

	client, _ := docker.NewClientFromEnv() // if using docker machine

	hostConfig := docker.HostConfig{
		Privileged: true,
		PublishAllPorts: true,
		Binds: []string{
			fmt.Sprintf("%s:/input:ro", s.InputDir),
			fmt.Sprintf("%s:/output", s.OutputDir),
		},
	}

	cmd := strings.Join([]string{
		"mount -t overlayfs none -o lowerdir=/input,upperdir=/output /work",
		"cd /work",
		s.Command,
	}, " && ")

	createOpts := docker.CreateContainerOptions{
		Config: &docker.Config{
			Image: "java:8",
			Cmd: append([]string{"/bin/bash", "-c"}, cmd),
			WorkingDir: "/work",
		},
		HostConfig: &hostConfig,
	}

	container, err := client.CreateContainer(createOpts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Running container: %s\n", container.ID)

	err = client.StartContainer(container.ID, nil)
	if err != nil {
		panic(err)
	}

	status, err := client.WaitContainer(container.ID)
	if err != nil {
		panic(err)
	}

	//client.RemoveContainer(docker.RemoveContainerOptions{Force:true})

	return status
}
