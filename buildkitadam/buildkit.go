package main

import (
	"context"
	"os"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/util/system"
)

func main() {

	dt, err := adam().Marshal(context.TODO(), llb.LinuxAmd64)
	if err != nil {
		panic(err)
	}
	llb.WriteTo(dt, os.Stdout)
}

func adam() llb.State {
	return llb.Image("docker.io/library/golang:1.13-alpine").
		AddEnv("PATH", "/usr/local/go/bin:"+system.DefaultPathEnvUnix).
		AddEnv("GOPATH", "/go").
		Run(llb.Shlex("/bin/sh go --version || true")).Root()
}
