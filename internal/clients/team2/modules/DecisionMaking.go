package modules

import (
	"github.com/google/uuid"
)

// DecisionInputs - Inputs for making decisions
type DecisionInputs struct {
	SocialCapital *SocialCapital
	AgentID       uuid.UUID
}

// DecisionOutputs - Struct for outputs of different decision types
type DecisionOutputs struct {
	KickAgentID       uuid.UUID
	ShouldAcceptAgent bool
	ShouldChangeBike  bool
	BikeID            uuid.UUID
	GovernanceID      int
}

// DecisionModule - Module for handling various decisions
type DecisionModule struct{}

// NewDecisionModule - Constructor for DecisionModule
func NewDecisionModule() *DecisionModule {
	return &DecisionModule{}
}

// Based on social capital, decide which agent to kick through minimum capital
func (dm *DecisionModule) MakeKickDecision(inputs DecisionInputs) uuid.UUID {
	agentID := inputs.SocialCapital.GetAgentWithMinimumSocialCapital()
	return agentID
}

// Accept based on larger than accept threshold
func (dm *DecisionModule) MakeAcceptAgentDecision(inputs DecisionInputs) bool {
	socialCapitalScore := inputs.SocialCapital.SocialCapital[inputs.AgentID]
	return socialCapitalScore > AcceptThreshold
}

func (dm *DecisionModule) MakeBikeChangeDecision(inputs DecisionInputs) (bool, uuid.UUID) {
	// Logic to decide on bike change
	shouldChangeBike = false
	if inputs.SocialCapital.GetAverage() < LeaveBikeThreshold {
		shouldChangeBike = true
		bikeID = inputs.SocialCapital.GetBikeWithMaximumSocialCapital()
	}
	return shouldChangeBike, bikeID
}

// Decide on governance
func (dm *DecisionModule) MakeGovernanceDecision(inputs DecisionInputs) int {
	return 1
}
