package team2

import (
	"fmt"

	"github.com/google/uuid"
)

// TODO: function CalculateSocialCapital
func (a *AgentTwo) CalculateSocialCapital() {
	// Implement this method
	// Hardcode the weightings for now: Trust 1, Institution 0, Network 0
	// Calculate social capital of all agents
	// Calculate trustworthiness of all agents
	// Calculate social networks of all agents
	// Calculate institutions of all agents
	// Iterate over each agent

	// Normalize Social Capital Values.
	normNetworks := normalizeMapValues(a.Network)
	normtrustworthiness := normalizeMapValues(a.Trust)
	normInstitution := normalizeMapValues(a.Institution)

	// Caluculate.
	for agentID := range a.Trust {
		network := normNetworks[agentID] // Assuming these values are already calculated
		institution := normInstitution[agentID]
		trustworthiness := normtrustworthiness[agentID]

		a.SocialCapital[agentID] = TrustWeight*trustworthiness + InstitutionWeight*institution + NetworkWeight*network
	}
}

func (a *AgentTwo) updateTrustworthiness(agentID uuid.UUID, actualAction, expectedAction ForceVector) {
	// Calculates the cosine Similarity of actual and expected vectors. One issue is that it does not consider magnitude, only direction
	// TODO: Take magnitude into account
	similarity := cosineSimilarity(actualAction, expectedAction)

	// CosineSimilarity output ranges from -1 to 1. Need to scale it back to 0-1
	normalisedTrustworthiness := (similarity + 1) / 2

	// Bad action but with high trustworthiness in prev rounds, we feel remorse and we forgive them
	if a.Trust[agentID] > normalisedTrustworthiness && a.forgivenessCounter <= 3 { // If they were trustworthy in prev rounds, we feel remorse and we forgive them
		a.forgivenessCounter++
		a.Trust[agentID] = (a.Trust[agentID]*float64(a.GameIterations) + (normalisedTrustworthiness + forgivenessFactor*(normalisedTrustworthiness-a.Trust[agentID]))) / (float64(a.GameIterations) + 1)
	} else if a.forgivenessCounter > 3 {
		// More than 3 rounds of BETRAYAL, we don't forgive them anymore...
		a.Trust[agentID] = (a.Trust[agentID]*float64(a.GameIterations) + normalisedTrustworthiness) / (float64(a.GameIterations) + 1)
	} else {
		// Good action with high trustworthiness
		a.forgivenessCounter = 0
		a.Trust[agentID] = (a.Trust[agentID]*float64(a.GameIterations) + normalisedTrustworthiness) / (float64(a.GameIterations) + 1)
	}

	fmt.Println("Trust: ", a.Trust)

}

/// ///
/// Networks
/// ///

func (a *AgentTwo) UpdateEnergyLevel(energyLevel float64) {
	// Signal that a loot box has collected.
	// We treat this as a social event and update the Network parameter in Social Capital.
	if energyLevel > 0.0 {
		bikeId := a.GetBike()
		fellowBikers := a.gameState.GetMegaBikes()[bikeId].GetAgents()
		for _, biker := range fellowBikers {
			bikerId := biker.GetID()
			if _, ok := a.Network[bikerId]; !ok {
				a.Network[bikerId] = 0.0
			}
			a.Network[bikerId] += 1.0 * energyLevel
		}
	}

	// Update energy level.
	a.energyLevel += energyLevel
}

// // func (a *AgentTwo) updateInstitution(agentID uuid.UUID) float64 {

// // 	// return 0.5 // This is just a placeholder value
// // }

// // func (a *AgentTwo) calculateTrustworthiness(agentID uuid.UUID) float64 {

// // 	return 0.5 // This is just a placeholder value
// // }

// // func (a *AgentTwo) calculateInstitution(agentID uuid.UUID) float64 {

// // 	// return 0.5 // This is just a placeholder value
// // }
