package main

import (
	"fmt"
	//"time"

	"github.com/fsouza/go-dockerclient"
)

func main() {
	//endpoint := "unix:///var/run/docker.sock"
	//client, _ := docker.NewClient(endpoint)
	client, _ := docker.NewClientFromEnv() // if using docker machine
	hostConfig := docker.HostConfig{PublishAllPorts: true}
	createOpts := docker.CreateContainerOptions{
		Config: &docker.Config{
			Image: "busybox",
			//Cmd:   []string{"echo", fmt.Sprintf("hello pat! it's %s", time.Now())},
			Cmd:   []string{"sleep", "10"},
		},
		HostConfig: &hostConfig,
	}
	container, err := client.CreateContainer(createOpts)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	err = client.StartContainer(container.ID, &hostConfig)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Container started successfully!")
	}
}