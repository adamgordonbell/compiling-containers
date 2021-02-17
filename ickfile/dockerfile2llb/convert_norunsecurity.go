// +build !dfrunsecurity

package dockerfile2llb

import (
	"github.com/agbell/compiling-containers/ickfile/instructions"
	"github.com/moby/buildkit/client/llb"
)

func dispatchRunSecurity(c *instructions.RunCommand) (llb.RunOption, error) {
	return nil, nil
}
