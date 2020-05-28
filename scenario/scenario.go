package scenario

import "github.com/Gaoey/stubby/stub"

func New(scenarioId string, stubs []stub.Stubby) *Scenario {
	scenarioes := initializeScenario(stubs)
	collection := NewScenarioCollection(scenarioes)
	scenario := collection.GetScenarioByID(scenarioId)
	return scenario
}

func initializeScenario(stubs []stub.Stubby) []*Scenario {
	s1 := NewScenario("scenario1", "Test Case 1", "For Test 1", stubs)
	s1.AddStep(stub.EXAMPLE, "cbpay1")

	return []*Scenario{s1}
}
