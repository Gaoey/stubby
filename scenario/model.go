package scenario

import (
	"github.com/Gaoey/stubby/stub"
)

type ScenarioCollection struct {
	Scenarios map[string]*Scenario
}

func NewScenarioCollection(sce []*Scenario) ScenarioCollection {
	var sc ScenarioCollection
	for _, s := range sce {
		sc.Scenarios[s.ID] = s
	}
	return sc
}

func (s ScenarioCollection) GetScenarioByID(id string) *Scenario {
	return s.Scenarios[id]
}

type Scenario struct {
	ID         string
	Name       string
	Desciption string
	Stubs      []stub.Stubby
	Step       map[string]stub.Stubby
}

func (s *Scenario) AddStep(stubName string, stubId string) *Scenario {
	for _, e := range s.Stubs {
		if e.ID == stubId {
			s.Step[stubName] = e
		}
	}
	return s
}

func (s *Scenario) GetStubByName(stubName string) stub.Stubby {
	return s.Step[stubName]
}

func NewScenario(id, name, description string, stubs []stub.Stubby) *Scenario {
	return &Scenario{
		ID:         id,
		Name:       name,
		Desciption: description,
		Stubs:      stubs,
	}
}
