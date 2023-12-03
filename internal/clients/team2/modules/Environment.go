package modules

import (
	objects "SOMAS2023/internal/common/objects"
	"SOMAS2023/internal/common/utils"
	"math"

	"github.com/google/uuid"
)

type EnvironmentModule struct {
	AgentId   uuid.UUID
	GameState objects.IGameState
	BikeId    uuid.UUID
}

///
/// GameState
///

func (e *EnvironmentModule) SetGameState(gameState objects.IGameState) {
	e.GameState = gameState
}

///
/// Lootboxes
///

func (e *EnvironmentModule) GetLootBoxes() map[uuid.UUID]objects.ILootBox {
	return e.GameState.GetLootBoxes()
}

func (e *EnvironmentModule) GetLootBoxById(lootboxId uuid.UUID) objects.ILootBox {
	return e.GetLootBoxes()[lootboxId]
}

func (e *EnvironmentModule) GetLootBoxesByColor(color utils.Colour) map[uuid.UUID]objects.ILootBox {
	lootboxes := e.GetLootBoxes()
	lootboxesFiltered := make(map[uuid.UUID]objects.ILootBox)
	for _, lootbox := range lootboxes {
		if lootbox.GetColour() == color {
			lootboxesFiltered[lootbox.GetID()] = lootbox
		}
	}
	return lootboxesFiltered
}

func (e *EnvironmentModule) GetNearestLootbox(agentId uuid.UUID) uuid.UUID {
	nearestLootbox := uuid.Nil
	minDist := math.MaxFloat64
	for _, lootbox := range e.GetLootBoxes() {
		dist := e.GetDistanceToLootbox(lootbox.GetID())
		if dist < minDist {
			minDist = dist
			nearestLootbox = lootbox.GetID()
		}
	}
	return nearestLootbox
}

func (e *EnvironmentModule) GetNearestLootboxByColor(agentId uuid.UUID, color utils.Colour) uuid.UUID {
	nearestLootbox := uuid.Nil
	minDist := math.MaxFloat64
	for _, lootbox := range e.GetLootBoxesByColor(color) {
		dist := e.GetDistanceToLootbox(lootbox.GetID())
		if dist < minDist {
			minDist = dist
			nearestLootbox = lootbox.GetID()
		}
	}
	return nearestLootbox
}

func (e *EnvironmentModule) GetDistanceToLootbox(lootboxId uuid.UUID) float64 {
	bikePos, agntPos := e.GetBikeById(e.BikeId).GetPosition(), e.GetLootBoxById(lootboxId).GetPosition()

	return e.GetDistance(bikePos, agntPos)
}

///
/// Bikes
///

func (e *EnvironmentModule) GetAudi() objects.IAudi {
	return e.GameState.GetAudi()
}

func (e *EnvironmentModule) GetBikes() map[uuid.UUID]objects.IMegaBike {
	return e.GameState.GetMegaBikes()
}

func (e *EnvironmentModule) GetBikeById(bikeId uuid.UUID) objects.IMegaBike {
	return e.GetBikes()[bikeId]
}

func (e *EnvironmentModule) GetDistanceToAudi() float64 {
	bikePos, audiPos := e.GetBikeById(e.BikeId).GetPosition(), e.GetAudi().GetPosition()

	return e.GetDistance(bikePos, audiPos)
}

///
/// Utils
///

func (e *EnvironmentModule) GetDistance(pos1, pos2 utils.Coordinates) float64 {
	return math.Sqrt(math.Pow(pos1.X-pos2.X, 2) + math.Pow(pos1.Y-pos2.Y, 2))
}
