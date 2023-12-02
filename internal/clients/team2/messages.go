package team2

import (
	obj "SOMAS2023/internal/common/objects"

	"github.com/MattSScott/basePlatformSOMAS/messaging"
	"github.com/google/uuid"
)

//////
/// Kickoff Messages
//////

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
	// TODO: Update Instituition Social Capital Values.
}
