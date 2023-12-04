package agent

import (
	"SOMAS2023/internal/clients/team2/modules"
	"SOMAS2023/internal/common/utils"

	"github.com/google/uuid"
)

type Action struct {
	AgentID         uuid.UUID
	Action          string
	Force           utils.Forces
	GameLoop        int32
	lootBoxlocation modules.ForceVector //utils.Coordinates
}

func ProjectForce(actual, expected utils.Forces) float64 {
	actualVec := modules.GetForceVector(actual)
	expectVec := modules.GetForceVector(expected)
	return actualVec.CosineSimilarity(*expectVec) * actual.Pedal
}
