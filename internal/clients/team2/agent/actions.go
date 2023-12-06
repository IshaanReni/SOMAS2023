package agent

import (
	"SOMAS2023/internal/clients/team2/modules"
	"SOMAS2023/internal/common/objects"
	"SOMAS2023/internal/common/utils"
	"SOMAS2023/internal/common/voting"
	"maps"
	"math/rand"

	"github.com/google/uuid"
)

func (a *AgentTwo) DecideWeights(action utils.Action) map[uuid.UUID]float64 {
	// TODO: All actions have equal weights. Weighting by AgentId based on social capital.
	return a.BaseBiker.DecideWeights(action)
}

func (a *AgentTwo) VoteLeader() voting.IdVoteMap {
	// TODO: We vote for ourselves and highest SC agent.
	// Equal weights for both.
	return a.BaseBiker.VoteLeader()
}

func (a *AgentTwo) DecideGovernance() utils.Governance {
	// TODO: All possibilities except dictatorship.
	return a.BaseBiker.DecideGovernance()
}

func (a *AgentTwo) DecideAllocation() voting.IdVoteMap {
	socialCapital := maps.Clone(a.Modules.SocialCapital.SocialCapital)
	// Iterate through agents in social capital
	for id := range socialCapital {
		// Iterate through fellow bikers
		for _, biker := range a.GetFellowBikers() {
			// If this agent is a fellow biker, move on
			if biker.GetID() == id {
				continue
			}
		}
		// This agent is not a fellow biker - remove it from SC
		delete(socialCapital, id)
	}
	// We give ourselves 1.0
	socialCapital[a.GetID()] = 1.0
	return socialCapital
}

func (a *AgentTwo) DecideDictatorAllocation() voting.IdVoteMap {
	socialCapital := a.DecideAllocation()

	// Calculate the total social capital
	totalSocialCapital := 0.0
	for _, sc := range socialCapital {
		totalSocialCapital += sc
	}

	// Distribute the allocation based on each agent's share of the total social capital
	result := make(voting.IdVoteMap)
	for agentID, sc := range socialCapital {
		result[agentID] = sc / totalSocialCapital
	}
	return result
}

func (a *AgentTwo) VoteForKickout() map[uuid.UUID]int {
	// TODO: Vote for the agents with a Social Capital lower than a threshold.
	return a.BaseBiker.VoteForKickout()
}

func (a *AgentTwo) DecideJoining(pendingAgents []uuid.UUID) map[uuid.UUID]bool {
	// TODO: Accept all agents we don't know about or are higher in social capital.
	// If we know about them and they have a lower social capital, reject them.
	return a.BaseBiker.DecideJoining(pendingAgents)
}

func (a *AgentTwo) ProposeDirection() uuid.UUID {
	// TODO: Propose direction of lootbox with highest gain of our color.
	return a.BaseBiker.ProposeDirection()
}

func (a *AgentTwo) FinalDirectionVote(proposals map[uuid.UUID]uuid.UUID) voting.LootboxVoteMap {
	// TODO: If Social Capital of agent who proposed a lootbox is higher than a threshold, vote for it. Weight based on SC.
	// Otherwise, set a weight of 0.
	return a.BaseBiker.FinalDirectionVote(proposals)
}

func (a *AgentTwo) ChangeBike() uuid.UUID {
	decisionInputs := modules.DecisionInputs{SocialCapital: a.Modules.SocialCapital, Enviornment: a.Modules.Environment, AgentID: a.GetID()}
	isChangeBike, bikeId := a.Modules.Decision.MakeBikeChangeDecision(decisionInputs)
	if isChangeBike {
		return bikeId
	} else {
		return a.Modules.Environment.BikeId
	}
}

func (a *AgentTwo) DecideAction() objects.BikerAction {
	avgSocialCapital := a.Modules.SocialCapital.GetAverage(a.Modules.SocialCapital.SocialCapital)

	if avgSocialCapital > ChangeBikeSocialCapitalThreshold {
		// Pedal if members of the bike have high social capital.
		return objects.Pedal
	} else {
		// Otherwise, change bikes.
		return objects.ChangeBike
	}
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

func (a *AgentTwo) SetBike(bikeId uuid.UUID) {
	a.Modules.Environment.BikeId = bikeId
	a.BaseBiker.SetBike(bikeId)
}
