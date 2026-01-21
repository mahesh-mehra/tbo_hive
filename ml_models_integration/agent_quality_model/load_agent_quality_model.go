package agent_quality_model

import (
	"fmt"
	"github.com/korsajan/gocatboost"
	"sync"
	"tbo_backend/objects"
)

var (
	agentQualityModel *gocatboost.Catboost
	loadOnce          sync.Once
	loadErr           error
)

func LoadAgentQualityModel() error {
	loadOnce.Do(func() {
		// loading model in memory
		cb, err := gocatboost.FromFile(objects.ConfigObj.CatboostModelsPaths.FinalAgentQualityModel)

		// checking for error
		if err != nil {
			loadErr = err
			return
		}

		// updating agent quality model
		agentQualityModel = cb

		fmt.Println("Agent Quality Model Loaded!")
	})

	return loadErr
}
