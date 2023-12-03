package modules

import (
	"errors"

	"github.com/google/uuid"
)

type ISocialCapital interface {
	GetSocialCapital(agentID uuid.UUID) float64
}

type ISocialAttribute interface {
	UpdateValue(agentID uuid.UUID, newValue float64) error
}

type SocialCapital struct {
	Reputation    Reputation
	Institution   Institution
	SocialNetwork SocialNetwork
}

type Reputation struct {
	values map[uuid.UUID]float64
}

func (r *Reputation) UpdateValue(agentID uuid.UUID, EventValue float64) error {

	current, exists := r.values[agentID]
	if !exists {
		return errors.New("agent not found")
	}
	newValue := current + EventValue
	r.values[agentID] = newValue
	return nil
}

type Institution struct {
	values map[uuid.UUID]float64
}

func (i *Institution) UpdateValue(agentID uuid.UUID, EventValue float64) error {

	current, exists := i.values[agentID]
	if !exists {
		return errors.New("agent not found")
	}
	newValue := current + EventValue
	i.values[agentID] = newValue
	return nil
}

type SocialNetwork struct {
	values map[uuid.UUID]float64
}

func (sn *SocialNetwork) UpdateValue(agentID uuid.UUID, EventValue float64) error {

	current, exists := sn.values[agentID]
	if !exists {
		return errors.New("agent not found")
	}
	newValue := current + EventValue
	sn.values[agentID] = newValue
	return nil
}

func (sc *SocialCapital) GetSocialCapital(agentID uuid.UUID) float64 {
	return InstitutionWeight*sc.Institution.values[agentID] +
		NetworkWeight*sc.SocialNetwork.values[agentID] +
		ReputationWeight*sc.Reputation.values[agentID]
}

func NewSocialCapital() *SocialCapital {
	return &SocialCapital{
		Reputation:    Reputation{values: make(map[uuid.UUID]float64)},
		Institution:   Institution{values: make(map[uuid.UUID]float64)},
		SocialNetwork: SocialNetwork{values: make(map[uuid.UUID]float64)},
	}
}
