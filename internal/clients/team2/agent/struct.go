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

type AgentTwo struct {
	// BaseBiker represents a basic biker agent.
	*objects.BaseBiker

	// Modules
	EnvironmentModule   *modules.EnvironmentModule
	SocialCapitalModule *modules.SocialCapital
	DecisionModule      *modules.DecisionModule

	gameState        objects.IGameState // updated by the server at every round
	megaBikeId       uuid.UUID
	actions          []Action
	bikeCounter      map[uuid.UUID]int32
	soughtColour     utils.Colour // the colour of the lootbox that the agent is currently seeking
	onBike           bool
	energyLevel      float64 // float between 0 and 1
	points           int
	forces           utils.Forces
	allocationParams objects.ResourceAllocationParams
	votedDirection   uuid.UUID
}

func NewBaseTeam2Biker(agentId uuid.UUID) *AgentTwo {
	color := utils.GenerateRandomColour()
	baseBiker := objects.GetBaseBiker(color, agentId)
	return &AgentTwo{
		// BaseBiker
		BaseBiker: baseBiker,

		// Modules
		EnvironmentModule:   modules.GetEnvironmentModule(baseBiker.GetID(), baseBiker.GetGameState(), baseBiker.GetMegaBikeId()),
		SocialCapitalModule: modules.NewSocialCapital(),
		DecisionModule:      modules.NewDecisionModule(),

		gameState:        nil,
		megaBikeId:       uuid.UUID{},
		bikeCounter:      make(map[uuid.UUID]int32),
		soughtColour:     color,
		onBike:           false,
		energyLevel:      1.0,
		points:           0,
		forces:           utils.Forces{},
		allocationParams: objects.ResourceAllocationParams{},
		votedDirection:   uuid.UUID{},
	}
}
