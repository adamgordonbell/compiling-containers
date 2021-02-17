// Package command contains the set of Dockerfile commands.
package command

// Define constants for the command strings
const (
	Add         = "add"
	Arg         = "var"
	Cmd         = "do"
	Copy        = "stash"
	Entrypoint  = "please_do"
	Env         = "e222"
	Expose      = "expose"
	From        = "come_from"
	Healthcheck = "are_you_ok"
	Label       = "e252"
	Maintainer  = "e256"
	Onbuild     = "abstain"
	Run         = "please"
	Shell       = "mystery"
	StopSignal  = "i_do_not_compute"
	User        = "writing"
	Volume      = "->"
	Workdir     = "<-"
)

// Commands is list of all Dockerfile commands
var Commands = map[string]struct{}{
	Add:         {},
	Arg:         {},
	Cmd:         {},
	Copy:        {},
	Entrypoint:  {},
	Env:         {},
	Expose:      {},
	From:        {},
	Healthcheck: {},
	Label:       {},
	Maintainer:  {},
	Onbuild:     {},
	Run:         {},
	Shell:       {},
	StopSignal:  {},
	User:        {},
	Volume:      {},
	Workdir:     {},
}
