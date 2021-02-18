package main

import (
	"context"
	"os"

	"github.com/moby/buildkit/client/llb"
)

func main() {

	dt, err := createLLBState().Marshal(context.TODO(), llb.LinuxAmd64)
	if err != nil {
		panic(err)
	}
	llb.WriteTo(dt, os.Stdout)
}

func createLLBState() llb.State {
	return llb.Image("docker.io/library/alpine").
		File(llb.Copy(llb.Local("context"), "README.md", "README.md")).
		Run(llb.Args([]string{"/bin/sh", "-c", "echo \"programmatically built\" > /built.txt"})).Root()
}
