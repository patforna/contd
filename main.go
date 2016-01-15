package main

import (
	"fmt"
	"os"

	"github.com/patforna/splendid/shipper"
)

func main() {
	shipper := shipper.Shipper{
		Image: "java:8",
		Command: "javac -verbose Hello.java",
		InputDir: "/tmp/input/x",
		OutputDir: "/tmp/output/x",
	}
	
	status := shipper.Run()

	fmt.Println("Done.")
	os.Exit(status)

}