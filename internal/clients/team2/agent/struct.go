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

	// CalculateSocialCapitalOtherAgent: (trustworthiness - cosine distance, social networks - friends, institutions - num of rounds on a bike)
	SocialCapital      map[uuid.UUID]float64 // Social Captial of other agents
	Reputation         map[uuid.UUID]float64 // Reputation of other agents
	Institution        map[uuid.UUID]float64 // Institution of other agents
	Network            map[uuid.UUID]float64 // Network of other agents
	GameIterations     int32                 // Keep track of game iterations // TODO: WHAT IS THIS?
	forgivenessCounter int32                 // Keep track of how many rounds we have been forgiving an agent
	gameState          objects.IGameState    // updated by the server at every round
	megaBikeId         uuid.UUID
	actions            []Action
	bikeCounter        map[uuid.UUID]int32
	soughtColour       utils.Colour // the colour of the lootbox that the agent is currently seeking
	onBike             bool
	energyLevel        float64 // float between 0 and 1
	points             int
	forces             utils.Forces
	allocationParams   objects.ResourceAllocationParams
	votedDirection     uuid.UUID
}

func NewBaseTeam2Biker(agentId uuid.UUID) *AgentTwo {
	color := utils.GenerateRandomColour()
	baseBiker := objects.GetBaseBiker(color, agentId)
	return &AgentTwo{
		// BaseBiker
		BaseBiker: baseBiker,

		// Modules
		EnvironmentModule: modules.GetEnvironmentModule(baseBiker.GetID(), baseBiker.GetGameState(), baseBiker.GetMegaBikeId()),

		SocialCapital:      make(map[uuid.UUID]float64),
		Reputation:         make(map[uuid.UUID]float64),
		Institution:        make(map[uuid.UUID]float64),
		Network:            make(map[uuid.UUID]float64),
		GameIterations:     0,
		forgivenessCounter: 0,
		gameState:          nil,
		megaBikeId:         uuid.UUID{},
		bikeCounter:        make(map[uuid.UUID]int32),
		soughtColour:       color,
		onBike:             false,
		energyLevel:        1.0,
		points:             0,
		forces:             utils.Forces{},
		allocationParams:   objects.ResourceAllocationParams{},
		votedDirection:     uuid.UUID{},
	}
}
