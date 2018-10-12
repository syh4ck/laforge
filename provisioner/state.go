package provisioner

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"

	"github.com/gen0cide/laforge/core"
)

// State is the global state of this server and all required information
type State struct {
	sync.RWMutex
	Host            *core.Host            `json:"host,omitempty"`
	Network         *core.Network         `json:"network,omitempty"`
	Environment     *core.Environment     `json:"environment,omitempty"`
	Competition     *core.Competition     `json:"competition,omitempty"`
	Team            *core.Team            `json:"team,omitempty"`
	ProvisionedHost *core.ProvisionedHost `json:"provisioned_host,omitempty"`
	RenderedAt      time.Time             `json:"rendered_at,omitempty"`
	InitializedAt   time.Time             `json:"initialized_at,omitempty"`
	CompletedAt     time.Time             `json:"completed_at,omitempty"`
	CurrentState    string                `json:"current_state,omitempty"`
	Steps           []*Step               `json:"steps,omitempty"`
	Pending         map[int]*Step         `json:"pending_steps"`
	Completed       map[int]*Step         `json:"completed_steps"`
	CurrentStep     *Step                 `json:"current_step,omitempty"`
}

// LoadStateFile parses a JSON state file into a state object
func LoadStateFile(location string) (*State, error) {
	fdata, err := ioutil.ReadFile(location)
	if err != nil {
		return nil, err
	}
	newState := &State{}
	err = json.Unmarshal(fdata, newState)
	if err != nil {
		return nil, err
	}
	return newState, nil
}