package main

import (
	"context"
	"os"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/util/system"
)

func main() {

	dt, err := createState().Marshal(context.TODO(), llb.LinuxAmd64)
	if err != nil {
		panic(err)
	}
	llb.WriteTo(dt, os.Stdout)
}

func createState() llb.State {
	return llb.Image("docker.io/library/golang:1.13-alpine").
		File(llb.Copy(llb.Local("mylocal"), "file.txt", "file.txt")).
		AddEnv("PATH", "/usr/local/go/bin:"+system.DefaultPathEnvUnix).
		AddEnv("GOPATH", "/go").
		Run(llb.Shlex("/bin/sh go --version || true")).Root()
}
