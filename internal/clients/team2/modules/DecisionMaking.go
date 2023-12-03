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

type KickDecisionModule struct {
	IDecisionModule[DecisionInputs, KickDecisionOutput]
}

func NewKickDecisionMoudle() *KickDecisionModule {
	return &KickDecisionModule{}
}

func (kd *KickDecisionModule) MakeDecision(inputs DecisionInputs) (KickDecisionOutput, error) {

	socialCapitalScore := inputs.SocialCapital.GetSocialCapital(inputs.AgentID)

	shouldKick := false
	if socialCapitalScore < KickThreshold {
		shouldKick = true
	}

	return KickDecisionOutput{
		ShouldKick: shouldKick,
		AgentID:    inputs.AgentID,
	}, nil
}

type AcceptAgentDecisionOutput struct {
	ShouldAccept bool
	AgentID      uuid.UUID
}

type AcceptAgentDecisionModule struct {
	IDecisionModule[DecisionInputs, AcceptAgentDecisionModule]
}

func NewAcceptAgentDecision() *AcceptAgentDecisionModule {
	return &AcceptAgentDecisionModule{}
}

func (ad *AcceptAgentDecisionModule) MakeDecision(inputs DecisionInputs) (AcceptAgentDecisionOutput, error) {

	socialCapitalScore := inputs.SocialCapital.GetSocialCapital(inputs.AgentID)

	shouldAccept := false
	if socialCapitalScore > AcceptThreshold {
		shouldAccept = true
	}

	return AcceptAgentDecisionOutput{
		ShouldAccept: shouldAccept,
		AgentID:      inputs.AgentID,
	}, nil
}

type BikeChangeDecisionOutput struct {
	ShouldChange bool
	BikeID       uuid.UUID
}

type BikeChangeDecisionModule struct {
	IDecisionModule[DecisionInputs, BikeChangeDecisionOutput]
}

func NewBikeChangeDecision() *AcceptAgentDecisionModule {
	return &AcceptAgentDecisionModule{}
}

func (bc *BikeChangeDecisionModule) MakeDecision(inputs DecisionInputs) (BikeChangeDecisionOutput, error) {

	bikeID := uuid.Nil
	return BikeChangeDecisionOutput{
		ShouldChange: true,
		BikeID:       bikeID,
	}, nil
}

type GovernanceDecisionOutput struct {
	GovernanceID int
}

type GovernanceDecisionModule struct {
	IDecisionModule[DecisionInputs, GovernanceDecisionOutput]
}

func NewGovernanceDecision() *AcceptAgentDecisionModule {
	return &AcceptAgentDecisionModule{}
}

func (gd *GovernanceDecisionModule) MakeDecision(inputs DecisionInputs) (GovernanceDecisionOutput, error) {

	//defualt vote for leader
	governanceID := 1

	return GovernanceDecisionOutput{
		GovernanceID: governanceID,
	}, nil
}
