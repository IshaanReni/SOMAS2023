package modules

import "github.com/google/uuid"

type IDecisionModule[I any, O any] interface {
	MakeDecision(input I) O
}

type DecisionInputs struct {
	SocialCapital *SocialCapital
	AgentID       uuid.UUID
}

type KickDecisionOutput struct {
	ShouldKick bool
	AgentID    uuid.UUID
}

type KickDecision struct {
	IDecisionModule[DecisionInputs, KickDecisionOutput]
	inputs *DecisionInputs
}

func (kd *KickDecision) MakeDecision(inputs DecisionInputs) (KickDecisionOutput, error) {
	
	socialCapitalScore := inputs.SocialCapital.GetSocialCapital(inputs.AgentID)

	shouldKick := false                     // Example logic based on the social capital score
	if socialCapitalScore <  { // Define someThreshold based on your criteria
		shouldKick = true
	}

	return KickDecisionOutput{
		ShouldKick: shouldKick,
		AgentID:    inputs.AgentID,
	}, nil // Return an error if something goes wrong in the decision process
}
