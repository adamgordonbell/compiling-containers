// +build !dfrunnetwork

package dockerfile2llb

import (
	"github.com/agbell/compiling-containers/dockerfile/instructions"
	"github.com/moby/buildkit/client/llb"
)

func dispatchRunNetwork(c *instructions.RunCommand) (llb.RunOption, error) {
	return nil, nil
}
