package agent_risk_model

import (
	"fmt"
	"github.com/korsajan/gocatboost"
	"sync"
	"tbo_backend/objects"
)

var (
	agentRiskModel *gocatboost.Catboost
	loadOnce       sync.Once
	loadErr        error
)

func LoadAgentRiskModel() error {
	loadOnce.Do(func() {
		// loading model in memory
		cb, err := gocatboost.FromFile(objects.ConfigObj.CatboostModelsPaths.AgentRiskModel)

		// checking for error
		if err != nil {
			loadErr = err
			return
		}

		// updating agent risk model
		agentRiskModel = cb

		fmt.Println("Agent Risk Model loaded!")
	})

	return loadErr
}
