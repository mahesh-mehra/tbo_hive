package fraud_history_model

import (
	"fmt"
	"github.com/korsajan/gocatboost"
	"sync"
	"tbo_backend/objects"
)

var (
	fraudHistoryModel *gocatboost.Catboost
	loadOnce          sync.Once
	loadErr           error
)

func LoadFraudHistoryModel() error {
	loadOnce.Do(func() {
		// loading model in memory
		cb, err := gocatboost.FromFile(objects.ConfigObj.CatboostModelsPaths.FraudHistoryModel)

		// checking for error
		if err != nil {
			loadErr = err
			return
		}

		// updating fraud history model
		fraudHistoryModel = cb

		fmt.Println("Fraud History Model Loaded!")
	})

	return loadErr
}
