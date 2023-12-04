package agent

import (
	"SOMAS2023/internal/clients/team2/modules"
	"SOMAS2023/internal/common/utils"
	"math"

	"github.com/google/uuid"
)

// TODO: Add an update reputation function.

//////
/// Institutions
//////

// Get the direction to the voted lootbox
func (a *AgentTwo) GetVotedLootboxForces(lootboxID uuid.UUID) utils.Forces {
	lootbox := a.gameState.GetLootBoxes()[lootboxID]
	lootboxPositionX, lootboxPositionY := lootbox.GetPosition().X, lootbox.GetPosition().Y
	agentPositionX, agentPositionY := a.GetLocation().X, a.GetLocation().Y
	deltaX := lootboxPositionX - agentPositionX
	deltaY := lootboxPositionY - agentPositionY
	angle := math.Atan2(deltaY, deltaX)
	normalisedAngle := angle / math.Pi
	turningDecision := utils.TurningDecision{
		SteerBike:     true,
		SteeringForce: normalisedAngle - a.gameState.GetMegaBikes()[a.GetBike()].GetOrientation(),
	}
	return utils.Forces{
		Pedal:   utils.BikerMaxForce,
		Brake:   0.0,
		Turning: turningDecision,
	}
}

// Called by Events to obtain Event Value for update Institution
// Assume what they broadcast is the truth
// TODO: Obtain actual action performed from messaging
// 1. Rule Adhereance (Follow leader biker/ dictator)
func (a *AgentTwo) RuleAdherenceValue(agentID uuid.UUID, expectedAction, actualAction utils.Forces) float64 {
	actualVec := modules.GetForceVector(actualAction)
	expectVec := modules.GetForceVector(expectedAction)
	return actualVec.CosineSimilarity(*expectVec) * actualAction.Pedal
}
