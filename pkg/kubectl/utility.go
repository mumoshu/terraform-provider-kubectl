package kubectl

import (
	"github.com/mumoshu/terraform-provider-eksctl/pkg/sdk"
	"log"
	"os/exec"
)

// State is a wrapper around both the input and output attributes that are relavent for updates
type State struct {
	Output string
}

// NewState is the constructor for State
func NewState() *State {
	return &State{}
}

func runCommand(ctx *sdk.Context, cmd *exec.Cmd, state *State, diffMode bool) (*State, error) {
	res, err := ctx.Run(cmd)
	if err != nil {
		return nil, err
	}

	newState := NewState()
	if diffMode && res.ExitStatus == 0 {
		newState.Output = ""
	} else {
		newState.Output = res.Output
	}

	log.Printf("[DEBUG] helmfile command new state: \"%v\"", newState)

	return newState, nil
}
