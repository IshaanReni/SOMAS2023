package agent

import (
	"SOMAS2023/internal/clients/team2/modules"
	"SOMAS2023/internal/common/objects"
	"SOMAS2023/internal/common/utils"
	"math/rand"

	"github.com/google/uuid"
)

// DecideAction() BikerAction                                      // ** determines what action the agent is going to take this round. (changeBike or Pedal)
// DecideForce(direction uuid.UUID)                                // ** defines the vector you pass to the bike: [pedal, brake, turning]
// DecideJoining(pendinAgents []uuid.UUID) map[uuid.UUID]bool      // ** decide whether to accept or not accept bikers, ranks the ones
// ChangeBike() uuid.UUID                                          // ** called when biker wants to change bike, it will choose which bike to try and join
// ProposeDirection() uuid.UUID                                    // ** returns the id of the desired lootbox based on internal strategy
// FinalDirectionVote(proposals []uuid.UUID) voting.LootboxVoteMap // ** stage 3 of direction voting
// DecideAllocation() voting.IdVoteMap                             // ** decide the allocation parameters

func (a *AgentTwo) ChangeBike() uuid.UUID {
	decisionInputs := modules.DecisionInputs{SocialCapital: a.Modules.SocialCapital, Enviornment: a.Modules.Environment, AgentID: a.GetID()}
	isChangeBike, bikeId := a.Modules.Decision.MakeBikeChangeDecision(decisionInputs)
	if isChangeBike {
		return bikeId
	} else {
		return uuid.Nil
	}
}

func (a *AgentTwo) DecideAction() objects.BikerAction {
	return objects.Pedal
	// TODO: Use SC in decide action.

	// fmt.Println("DecideAction entering")
	// // lootBoxlocation := Vector{X: 0.0, Y: 0.0} // need to change this later on (possibly need to alter the updateTrustworthiness function)
	// //update agent's trustworthiness every round pretty much at the start of each epoch
	// a.gameState = a.GetGameState()

	// // fmt.Println("DecideAction megabikes: ", a.gameState.GetMegaBikes())
	// for id := range a.Modules.Environment.GetBikerAgents() {
	// 	// get the force for the agent with agentID in actions
	// 	// fmt.Println("DecideAction agentID: ", agentID)
	// 	for _, action := range a.actions {
	// 		// fmt.Println("DecideAction action: ", action)
	// 		if action.AgentID == id {
	// 			// update trustworthiness
	// 			// Needs to be updated so that a.NearLootbox() is replaced with the lootbox location that the agent says that they're going for
	// 			a.updateReputation(id, a.GetOptimalLootbox(), a.nearestLoot())
	// 		}
	// 	}
	// 	// a.updateTrustworthiness(agent.GetID(), forcesToVectorConversion(), lootBoxlocation)
	// }
	// // a.gameState.GetMegaBikes()[a.GetBike()].GetAgents()[0].GetForces()
	// // Check energy level, if below threshold, don't change bike
	// // energyThreshold := 0.2
	// // fmt.Println("OUTSIDE FOR LOOP: ", a.GetEnergyLevel(), energyThreshold, a.ChooseOptimalBike(), a.GetBike())

	// // TODO: ChangeBike is broken in GameLoop
	// // if (a.GetEnergyLevel() < energyThreshold) || (a.ChooseOptimalBike() == a.GetBike()) {
	// // 	return objects.Pedal
	// // } else {
	// // 	// random for now, changeBike changes to a random uuid for now.
	// // 	return objects.ChangeBike
	// // }
	// return objects.Pedal

	// // TODO: When we have access to limbo/void then we can worry about these
	// // Utility = expected gain - cost of changing bike(no of rounds in the void * energy level drain)
	// // no of rounds in the void = 1 + (distance to lootbox / speed of bike)
}

func (a *AgentTwo) DecideForce(direction uuid.UUID) {

	a.Modules.VotedDirection = direction

	if a.Modules.Environment.IsAudiNear() {
		// Move in opposite direction to Audi in full force
		bikePos, audiPos := a.Modules.Environment.GetBike().GetPosition(), a.Modules.Environment.GetAudi().GetPosition()
		force := a.Modules.Utils.GetForcesToTargetWithDirectionOffset(utils.BikerMaxForce, -180.0, bikePos, audiPos)
		a.SetForces(force)
		return
	}
	// Use the average social capital to decide whether to pedal in the voted direciton or not
	probabilityOfConformity := a.Modules.SocialCapital.GetAverage(a.Modules.SocialCapital.SocialCapital)
	randomNumber := rand.Float64()
	agentPosition := a.GetLocation()
	lootboxID := direction
	if randomNumber > probabilityOfConformity {
		lootboxID = a.Modules.Environment.GetHighestGainLootbox()
	}
	lootboxPosition := a.Modules.Environment.GetLootboxPos(lootboxID)
	force := a.Modules.Utils.GetForcesToTarget(agentPosition, lootboxPosition)
	a.SetForces(force)
}

func (a *AgentTwo) UpdateGameState(gameState objects.IGameState) {
	a.BaseBiker.UpdateGameState(gameState)
	a.Modules.Environment.SetGameState(gameState)
}
