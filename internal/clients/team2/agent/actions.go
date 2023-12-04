package agent

import (
	"SOMAS2023/internal/common/objects"
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
	bikes := a.EnvironmentModule.GetBikes()

	// Find bike to change to.
	chgBikeId := uuid.Nil
	for id, bike := range bikes {
		cnt := 0
		for _, agent := range bike.GetAgents() {
			sc, ok := a.SocialCapital[agent.GetID()]
			if ok && sc < 0.5 {
				cnt++
			}
		}
		if cnt > 2 {
			chgBikeId = id
		}
	}

	if chgBikeId != uuid.Nil {
		// If found, change to that bike.
		return chgBikeId
	} else {
		// Otherwise, change to a random bike.
		i, targetI := 0, rand.Intn(len(bikes))
		for id := range bikes {
			if i == targetI {
				return id
			}
			i++
		}
		panic("No bikes found to change to.")
	}

	// If none found, change to a random bike.

}

func (a *AgentTwo) DecideAction() objects.BikerAction {
	return objects.Pedal
	// TODO: Use SC in decide action.

	// fmt.Println("DecideAction entering")
	// // lootBoxlocation := Vector{X: 0.0, Y: 0.0} // need to change this later on (possibly need to alter the updateTrustworthiness function)
	// //update agent's trustworthiness every round pretty much at the start of each epoch
	// a.gameState = a.GetGameState()

	// // fmt.Println("DecideAction megabikes: ", a.gameState.GetMegaBikes())
	// for id := range a.EnvironmentModule.GetBikerAgents() {
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
	// TODO: Use SC in decide forces.
	// if a.EnvironmentModule.IsAudiNear() {
	// 	// Move in opposite direction to Audi.
	// 	bikePos, audiPos := a.EnvironmentModule.GetBike().GetPosition(), a.EnvironmentModule.GetAudi().GetPosition()

	// 	deltaX := -audiPos.X + bikePos.X
	// 	deltaY := -audiPos.Y + bikePos.Y
	// 	steerA := math.Atan2(deltaY, deltaX)/math.Pi - a.EnvironmentModule.GetBikeOrientation()

	// 	forces := utils.Forces{
	// 		Pedal:   utils.BikerMaxForce,
	// 		Brake:   0.0,
	// 		Turning: utils.TurningDecision{SteerBike: true, SteeringForce: steerA},
	// 	}
	// 	a.SetForces(forces)
	// } else {
	// 	// Move towards lootbox with highest gain.
	// 	// TODO: Use SC in decision.
	// 	lootbox := a.EnvironmentModule.GetHighestGainLootbox()

	// 	bikePos, lootboxPos := a.EnvironmentModule.GetBike().GetPosition(), a.EnvironmentModule.GetLootboxPos(lootbox)
	// 	deltaX := lootboxPos.X - bikePos.X
	// 	deltaY := lootboxPos.Y - bikePos.Y
	// 	steerA := math.Atan2(deltaY, deltaX)/math.Pi - a.EnvironmentModule.GetBikeOrientation()

	// 	forces := utils.Forces{
	// 		Pedal:   utils.BikerMaxForce,
	// 		Brake:   0.0,
	// 		Turning: utils.TurningDecision{SteerBike: true, SteeringForce: steerA},
	// 	}
	// 	a.SetForces(forces)
	// }
}

func (a *AgentTwo) UpdateGameState(gameState objects.IGameState) {
	a.gameState = gameState
	a.EnvironmentModule.SetGameState(gameState)
}
