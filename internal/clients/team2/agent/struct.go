package agent

import (
	"SOMAS2023/internal/clients/team2/modules"
	"SOMAS2023/internal/common/objects"
	"SOMAS2023/internal/common/utils"

	"github.com/google/uuid"
)

type IBaseBiker interface {
	objects.IBaseBiker
}

type AgentModules struct {
	Environment   *modules.EnvironmentModule
	SocialCapital *modules.SocialCapital
	Decision      *modules.DecisionModule
	Utils         *modules.UtilsModule
}

// Honestly Have no idea what this is for ????
type Action struct {
	AgentID         uuid.UUID
	Action          string
	Force           utils.Forces
	GameLoop        int32
	lootBoxlocation modules.ForceVector //utils.Coordinates
}

type AgentState struct {
	Actions      []Action // Why is this Needed???
	SoughtColour utils.Colour
	OnBike       bool
	EnergyLevel  float64
	Points       int
	Forces       utils.Forces
}

type EnvironmentState struct {
	GameState      objects.IGameState
	MegaBikeId     uuid.UUID
	VotedDirection uuid.UUID
	BikeCounter    map[uuid.UUID]int32
}

type AgentTwo struct {
	*objects.BaseBiker // Embedding the BaseBiker
	Modules            AgentModules
	State              AgentState
	EnvironmentState   EnvironmentState
}

func NewBaseTeam2Biker(agentId uuid.UUID) *AgentTwo {
	color := utils.GenerateRandomColour()
	baseBiker := objects.GetBaseBiker(color, agentId)

	return &AgentTwo{
		BaseBiker: baseBiker,
		Modules: AgentModules{
			Environment:   modules.GetEnvironmentModule(baseBiker.GetID(), baseBiker.GetGameState(), baseBiker.GetBike()),
			SocialCapital: modules.NewSocialCapital(),
			Decision:      modules.NewDecisionModule(),
			Utils:         modules.NewUtilsModule(),
		},
		State: AgentState{
			Actions:      []Action{},
			SoughtColour: color,
			OnBike:       false,
			EnergyLevel:  1.0,
			Points:       0,
			Forces:       utils.Forces{},
		},
		EnvironmentState: EnvironmentState{
			GameState:      baseBiker.GetGameState(),
			MegaBikeId:     baseBiker.GetBike(),
			BikeCounter:    make(map[uuid.UUID]int32),
			VotedDirection: uuid.UUID{},
		},
	}
}
