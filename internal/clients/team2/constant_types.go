package team2

import (
	"SOMAS2023/internal/common/objects"
	"SOMAS2023/internal/common/utils"

	"github.com/google/uuid"
)

// Constants related to the calculation of social capital and trustworthiness
const (
	TrustWeight       = 1.0 // Weight for trust in social capital calculation
	InstitutionWeight = 0.0 // Weight for institution affiliation in social capital calculation
	NetworkWeight     = 0.0 // Weight for network strength in social capital calculation
)

const (
	forgivenessFactor = 0.5 // Factor used in trustworthiness update calculations
)

type IBaseBiker interface {
	objects.IBaseBiker
}

type AgentTwo struct {
	// BaseBiker represents a basic biker agent.
	*objects.BaseBiker
	// CalculateSocialCapitalOtherAgent: (trustworthiness - cosine distance, social networks - friends, institutions - num of rounds on a bike)
	SocialCapital      map[uuid.UUID]float64 // Social Captial of other agents
	Trust              map[uuid.UUID]float64 // Trust of other agents
	Institution        map[uuid.UUID]float64 // Institution of other agents
	Network            map[uuid.UUID]float64 // Network of other agents
	GameIterations     int32                 // Keep track of game iterations // TODO: WHAT IS THIS?
	forgivenessCounter int32                 // Keep track of how many rounds we have been forgiving an agent
	gameState          objects.IGameState    // updated by the server at every round
	megaBikeId         uuid.UUID
	bikeCounter        map[uuid.UUID]int32
	actions            []Action
	soughtColour       utils.Colour // the colour of the lootbox that the agent is currently seeking
	onBike             bool
	energyLevel        float64 // float between 0 and 1
	points             int
	forces             utils.Forces
	allocationParams   objects.ResourceAllocationParams
}

type ForceVector struct {
	X float64
	Y float64
}

type Action struct {
	AgentID         uuid.UUID
	Action          string
	Force           utils.Forces
	GameLoop        int32
	lootBoxlocation ForceVector //utils.Coordinates
}
