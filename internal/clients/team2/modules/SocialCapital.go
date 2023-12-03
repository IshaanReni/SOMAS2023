package modules

import (
	"math"

	"github.com/google/uuid"
)

type SocialCapital struct {
	forgivenessCounter int32
	SocialCapital      map[uuid.UUID]float64
	Reputation         map[uuid.UUID]float64
	Institution        map[uuid.UUID]float64
	SocialNetwork      map[uuid.UUID]float64
}

func (sc *SocialCapital) GetAverage(scComponent map[uuid.UUID]float64) float64 {
	var sum = 0.0
	for _, value := range scComponent {
		sum += value
	}
	return sum / float64(len(scComponent))
}

func (sc *SocialCapital) GetAgentWithMinimumSocialCapital() uuid.UUID {
	min := math.MaxFloat64
	minAgentId := uuid.Nil
	for agentId, value := range sc.Reputation {
		if sc.SocialCapital[agentId] < min {
			min = value
			minAgentId = agentId
		}
	}
	return minAgentId
}

func (sc *SocialCapital) ClipValues(input float64) float64 {
	value := input
	if value < 0 {
		value = 0
	}
	if value > 1 {
		value = 1
	}
	return value
}

func (sc *SocialCapital) UpdateValue(agentId uuid.UUID, eventValue float64, eventWeight float64, scComponent map[uuid.UUID]float64) {
	_, exists := scComponent[agentId]
	if !exists {
		scComponent[agentId] = sc.GetAverage(scComponent)
	}

	scComponent[agentId] = sc.ClipValues(scComponent[agentId] + eventValue*eventWeight)
}

func (sc *SocialCapital) UpdateReputation(agentId uuid.UUID, eventValue float64, eventWeight float64) {
	sc.UpdateValue(agentId, eventValue, eventWeight, sc.Reputation)
}

func (sc *SocialCapital) UpdateInstitution(agentId uuid.UUID, eventValue float64, eventWeight float64) {
	sc.UpdateValue(agentId, eventValue, eventWeight, sc.Institution)
}

func (sc *SocialCapital) UpdateSocialNetwork(agentId uuid.UUID, eventValue float64, eventWeight float64) {
	sc.UpdateValue(agentId, eventValue, eventWeight, sc.SocialNetwork)
}

// Must be called once every round.
func (sc *SocialCapital) UpdateSocialCapital(agentID uuid.UUID) float64 {
	reputation := sc.NormalizeValues(sc.Reputation)
	institution := sc.NormalizeValues(sc.Institution)
	socialNetwork := sc.NormalizeValues(sc.SocialNetwork)

	// Update Forgiveness Counter.
	newSocialCapital := ReputationWeight*reputation[agentID] + InstitutionWeight*institution[agentID] + NetworkWeight*socialNetwork[agentID]

	if sc.SocialCapital[agentID] < newSocialCapital {
		sc.forgivenessCounter = 0
	}

	if sc.SocialCapital[agentID] > newSocialCapital && sc.forgivenessCounter <= 3 {
		// Forgive if forgiveness counter is less than 3 and new social capital is less.
		sc.forgivenessCounter++
		sc.SocialCapital[agentID] = newSocialCapital + forgivenessFactor*(sc.SocialCapital[agentID]-newSocialCapital)
	} else {
		sc.SocialCapital[agentID] = newSocialCapital
	}
	return sc.SocialCapital[agentID]
}

func (sc *SocialCapital) NormalizeValues(component map[uuid.UUID]float64) map[uuid.UUID]float64 {
	if len(component) == 0 {
		return component
	}

	min := 0.0
	max := 0.0
	for _, value := range component {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	for key, value := range component {
		component[key] = (value - min) / (max - min)
	}

	return component
}

func NewSocialCapital() *SocialCapital {
	return &SocialCapital{
		forgivenessCounter: 0,
		SocialCapital:      make(map[uuid.UUID]float64),
		Reputation:         make(map[uuid.UUID]float64),
		Institution:        make(map[uuid.UUID]float64),
		SocialNetwork:      make(map[uuid.UUID]float64),
	}
}
