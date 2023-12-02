package team2

import (
	obj "SOMAS2023/internal/common/objects"

	"github.com/MattSScott/basePlatformSOMAS/messaging"
	"github.com/google/uuid"
)

func (a *AgentTwo) CreateForcesMessage() obj.ForcesMessage {
	// Currently this returns a default message which sends to all bikers on the biker agent's bike
	// For team's agent, add your own logic to communicate with other agents

	return obj.ForcesMessage{
		BaseMessage: messaging.CreateMessage[obj.IBaseBiker](a, a.GetFellowBikers()),
		AgentId:     a.GetID(),
		AgentForces: a.forces,
	}
}

func (a *AgentTwo) CreateKickOffMessage(kickOff bool, agentId uuid.UUID) obj.KickOffAgentMessage {
	return obj.KickOffAgentMessage{
		BaseMessage: messaging.CreateMessage[obj.IBaseBiker](a, a.GetFellowAgents()),
		AgentId:     agentId,
		KickOff:     kickOff,
	}
}

func (a *AgentTwo) SendKickOffMessage(kickOff bool, agentId uuid.UUID) {
	msg := a.CreateKickOffMessage(kickOff, agentId)
	recipients := msg.GetRecipients()
	for _, recip := range recipients {
		if a.GetID() == recip.GetID() {
			continue
		}
		msg.InvokeMessageHandler(recip)
	}
}

func (a *AgentTwo) HandleKickOffMessage(msg obj.KickOffAgentMessage) {
	agentId := msg.AgentId
	a.updateInstitution(agentId, InstitutionEventWeight_KickedOut, InstitutionKickoffEventValue)
}

func (a *AgentTwo) HandleForcesMessage(msg obj.ForcesMessage) {

	agentId := msg.AgentId
	agentForces := msg.AgentForces
	optimalLootbox := a.votedDirection
	optimalForces := a.GetVotedLootboxForces(optimalLootbox)

	EventValue := a.RuleAdhereanceValue(agentId, optimalForces, agentForces)
	a.updateInstitution(agentId, InstitutionEventWeight_Adhereance, EventValue)

}

func (a *AgentTwo) HandleJoiningMessage(msg obj.JoiningAgentMessage) {

	// sender := msg.BaseMessage.GetSender()
	// agentId := msg.AgentId
	// bikeId := msg.BikeId
	agentId := msg.AgentId
	a.updateInstitution(agentId, InstitutionEventWeight_Accepted, InstitutionAcceptedEventValue)
}
