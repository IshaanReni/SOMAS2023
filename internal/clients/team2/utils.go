package team2

import (
	"SOMAS2023/internal/common/utils"
	"fmt"
	"math"
	"sort"

	"github.com/google/uuid"
)

func (a *AgentTwo) ChooseLootBox(condition string) Lootboxes {
	// Implement this method
	// Calculate utility of all lootboxes (box colour and distance to it), to decide the action (e.g. pedal, brake, turn) to take
	var topLootboxes []struct {
		ID          uuid.UUID
		accordingTo float64
		Color       utils.Colour
	}
	// var top3Lootboxes []struct {
	// 	ID     uuid.UUID
	// 	Reward float64
	// 	Color  utils.Colour
	// }

	var condVal float64
	for _, lootbox := range a.gameState.GetLootBoxes() {
		if condition == "reward" {
			condVal = lootbox.GetTotalResources()
		} else if condition == "gain" {
			condVal = a.CalcExpectedGainForLootbox(lootbox.GetID())
		}
		topLootboxes = append(topLootboxes, struct {
			ID          uuid.UUID
			accordingTo float64
			Color       utils.Colour
		}{ID: lootbox.GetID(), accordingTo: condVal, Color: lootbox.GetColour()})
	}

	// Sort the lootboxes by gain from the mapping
	sort.Slice(topLootboxes, func(i, j int) bool {
		return topLootboxes[i].accordingTo > topLootboxes[j].accordingTo
	})

	// top3Lootboxes = topLootboxes
	if len(topLootboxes) > 3 {
		topLootboxes = topLootboxes[:3]
	}

	return topLootboxes
}

func (a *AgentTwo) ChooseOptimalBike() uuid.UUID {
	// Implement this method
	// Calculate utility of all bikes for our own survival (remember previous actions (has space, got lootbox, direction) of all bikes so you can choose a bike to move to to max our survival chances) -> check our reputation (trustworthiness, social networks, institutions)

	// Change to bike closest to optimal lootbox
	// determine optimal by looking at highest reward and good colo

	topLootboxes := a.ChooseLootBox("reward")
	var chosenLootbox uuid.UUID
	for _, top3 := range topLootboxes {
		if top3.Color == a.GetColour() {
			chosenLootbox = top3.ID
		} else {
			chosenLootbox = topLootboxes[0].ID
		}
	}

	// find the closest bike to the chosen lootbox
	lootboxLoc := a.gameState.GetLootBoxes()[chosenLootbox].GetPosition()
	shortestDist := math.MaxFloat64
	for _, bikeID := range a.gameState.GetMegaBikes() {
		bikeLoc := bikeID.GetPosition()
		currDist := math.Sqrt(math.Pow(lootboxLoc.X-bikeLoc.X, 2) + math.Pow(lootboxLoc.Y-bikeLoc.Y, 2))
		if currDist < shortestDist {
			// optimalBike := bikeID.GetID()
			shortestDist = currDist
		}
	}
	// change bike to optimalBike

	// - We change the bike if an agent sees more than N agents below a social capital threshold.
	var N int32 = 3
	SocialCapitalThreshold := 0.5
	// - N and the Social Capital Threshold could be varied.

	currentBikeID := a.GetBike()
	a.gameState = a.GetGameState()

	for bikeID, bike := range a.gameState.GetMegaBikes() {
		for _, agent := range bike.GetAgents() {
			if a.SocialCapital[agent.GetID()] > SocialCapitalThreshold {
				a.bikeCounter[bikeID]++
			}
		}
	}

	if a.bikeCounter[currentBikeID] > N {
		// Stay on bike
		return currentBikeID
	} else {
		// find max bike counter in a.bikeCounter map
		// change bike to that bike
		var maxValue int32 = 0
		maxBikeID := uuid.UUID{}
		for bikeID, counter := range a.bikeCounter {
			if maxValue < counter {
				maxValue = counter
				maxBikeID = bikeID
			}
		}
		return maxBikeID
	}
}

// find the number of agents on the bike
func (a *AgentTwo) GetAgentNum(bikeID uuid.UUID) float64 {
	var count = 0.0
	for _, agent := range a.gameState.GetMegaBikes()[bikeID].GetAgents() {
		// agentID := agent.GetID()
		_ = agent.GetID()
		count++
	}
	return count
}

// TODO: Once the MVP is complete, we can start thinking about this and then feed it into DecideForce
func (a *AgentTwo) CalcExpectedGainForLootbox(lootboxID uuid.UUID) float64 {
	// Implement this method
	// Calculate gain of going for a given lootbox(box colour and distance to it), to decide the action (e.g. pedal, brake, turn) to take

	currLocation := a.GetLocation()
	targetPos := a.gameState.GetLootBoxes()[lootboxID].GetPosition()

	deltaX := targetPos.X - currLocation.X
	deltaY := targetPos.Y - currLocation.Y
	angle := math.Atan2(deltaX, deltaY)
	angleInDegrees := angle * math.Pi / 180

	// Default BaseBiker will always
	turningDecision := utils.TurningDecision{
		SteerBike:     true,
		SteeringForce: angleInDegrees,
	}

	lootboxForces := utils.Forces{
		Pedal:   utils.BikerMaxForce,
		Brake:   0.0,
		Turning: turningDecision,
	}

	// assumes loot is equally divided among all agents on the bike
	energyFromLootbox := a.gameState.GetLootBoxes()[lootboxID].GetTotalResources() / a.GetAgentNum(a.GetBike())
	energyToPedal := lootboxForces.Pedal * utils.MovingDepletion //Moving depleteion is a constant 1 atm. TODO: Change this to a variable and check how Pedal relates to energy depletion

	expectedGain := energyFromLootbox - energyToPedal

	return expectedGain
}

func (a *AgentTwo) GetPreviousAction() {
	// -> get previous action of all bikes and bikers from last 5 gamestates

	// nearestLoot := a.nearestLoot()
	// currentLootBoxes := a.gameState.GetLootBoxes()
	// lootBoxlocation := currentLootBoxes[nearestLoot].GetPosition()
	lootBoxlocation_vector := ForceVector{X: 0.0, Y: 0.0} // need to change this later on (possibly need to alter the updateTrustworthiness function)
	//update agent's trustworthiness every round pretty much at the start of each epoch
	for _, bike := range a.gameState.GetMegaBikes() {
		for _, agent := range bike.GetAgents() {
			// Record the action
			action := Action{
				AgentID:         agent.GetID(),
				Action:          "DecideForce",
				Force:           agent.GetForces(),
				GameLoop:        a.GameIterations, // record the game loop number
				lootBoxlocation: lootBoxlocation_vector,
			}
			// If we have more than 5 actions, remove the oldest one
			if len(a.actions) >= 5 {
				a.actions = a.actions[1:]
			}

			// Append the new action
			a.actions = append(a.actions, action)
		}
	}
}

func (a *AgentTwo) AvoidOwdi(lootboxID uuid.UUID) bool {
	audiPos := a.GetGameState().GetAudi().GetPosition()
	lootboxPos := a.GetGameState().GetLootBoxes()[lootboxID].GetPosition()
	AudiRange := utils.Coordinates{
		X: 5.0,
		Y: 5.0,
	}
	fmt.Println("AudiRange: ", AudiRange)
	fmt.Println("audiPosX-: ", audiPos.X-AudiRange.X)
	fmt.Println("audiPosX+: ", audiPos.X+AudiRange.X)
	fmt.Println("audiPosY-: ", audiPos.Y-AudiRange.Y)
	fmt.Println("audiPosY+: ", audiPos.Y+AudiRange.Y)

	if audiPos.X-AudiRange.X < 0 || audiPos.Y-AudiRange.Y < 0 || audiPos.X+AudiRange.X > utils.GridWidth || audiPos.Y+AudiRange.Y > utils.GridHeight {

		AudiRange = utils.Coordinates{
			X: 0.0,
			Y: 0.0,
		}
	}
	// if we are near an owdi we want to avoid it
	if a.GetLocation().X > audiPos.X-AudiRange.X && a.GetLocation().X < audiPos.X+AudiRange.X && a.GetLocation().Y > audiPos.Y-AudiRange.Y && a.GetLocation().Y < audiPos.Y+AudiRange.Y {
		fmt.Println("avoiding owdi")
		return true
	}

	// if the lootbox is in the owdi's range we avoid it

	if lootboxPos.X > audiPos.X-AudiRange.X && lootboxPos.X < audiPos.X+AudiRange.X && lootboxPos.Y > audiPos.Y-AudiRange.Y && lootboxPos.Y < audiPos.Y+AudiRange.Y {
		fmt.Println("avoiding owdi")

		return true
	}

	return false
}

// returns the nearest lootbox with respect to the agent's bike current position
// in the MVP this is used to determine the pedalling forces as all agent will be
// aiming to get to the closest lootbox by default
func (a *AgentTwo) nearestLoot() uuid.UUID {
	currLocation := a.GetLocation()
	shortestDist := math.MaxFloat64
	var nearestBox uuid.UUID
	var currDist float64
	for _, loot := range a.gameState.GetLootBoxes() {
		x, y := loot.GetPosition().X, loot.GetPosition().Y
		currDist = math.Sqrt(math.Pow(currLocation.X-x, 2) + math.Pow(currLocation.Y-y, 2))
		if currDist < shortestDist {
			nearestBox = loot.GetID()
			shortestDist = currDist
		}
	}
	return nearestBox
}

func (a *AgentTwo) GetOptimalLootbox() uuid.UUID {

	topLootboxes := a.ChooseLootBox("gain")
	if len(topLootboxes) > 3 {
		topLootboxes = topLootboxes[:3]
	}
	fmt.Println("top3: ", topLootboxes)
	for _, top3 := range topLootboxes {
		if top3.Color == a.GetColour() {
			return top3.ID
		}
	}

	return topLootboxes[0].ID

}
