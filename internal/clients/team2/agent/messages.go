package agent

import (
	obj "SOMAS2023/internal/common/objects"

	"github.com/MattSScott/basePlatformSOMAS/messaging"
	"github.com/google/uuid"
)

func (a *AgentTwo) CreateForcesMessage() obj.ForcesMessage {
	return obj.ForcesMessage{
		BaseMessage: messaging.CreateMessage[obj.IBaseBiker](a, a.GetFellowBikers()),
		AgentId:     a.GetID(),
		AgentForces: a.forces,
	}
}

func (a *AgentTwo) CreateKickOffMessage() obj.KickOffAgentMessage {
	agentId := a.SocialCapitalModule.GetMinimumSocialCapital()
	kickOff := false
	if agentId != a.GetID() {
		kickOff = true
	}

	return obj.KickOffAgentMessage{
		BaseMessage: messaging.CreateMessage[obj.IBaseBiker](a, a.GetFellowBikers()),
		AgentId:     agentId,
		KickOff:     kickOff,
	}
}

func (a *AgentTwo) HandleKickOffMessage(msg obj.KickOffAgentMessage) {
	agentId := msg.AgentId

	if agentId != uuid.Nil {
		a.SocialCapitalModule.UpdateSocialNetwork(agentId, SocialEventValue_AgentSentMsg, SocialEventWeight_AgentSentMsg)
		a.SocialCapitalModule.UpdateInstitution(agentId, InstitutionEventValue_Kickoff, InstitutionEventWeight_Kickoff)
	}
}

func (a *AgentTwo) HandleForcesMessage(msg obj.ForcesMessage) {
	agentId := msg.AgentId
	optimalLootbox := a.votedDirection
	optimalForces := a.GetVotedLootboxForces(optimalLootbox)
	eventValue := ProjectForce(optimalForces, msg.AgentForces)

	a.SocialCapitalModule.UpdateSocialNetwork(agentId, SocialEventValue_AgentSentMsg, SocialEventWeight_AgentSentMsg)
	a.SocialCapitalModule.UpdateInstitution(agentId, InstitutionEventWeight_Adhereance, eventValue)
}

func (a *AgentTwo) HandleJoiningMessage(msg obj.JoiningAgentMessage) {
	agentId := msg.AgentId

	a.SocialCapitalModule.UpdateSocialNetwork(agentId, SocialEventValue_AgentSentMsg, SocialEventWeight_AgentSentMsg)
	a.SocialCapitalModule.UpdateInstitution(agentId, InstitutionEventValue_Accepted, InstitutionEventWeight_Accepted)
}
