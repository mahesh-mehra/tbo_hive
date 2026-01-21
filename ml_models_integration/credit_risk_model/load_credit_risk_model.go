package credit_risk_model

import (
	"fmt"
	"github.com/korsajan/gocatboost"
	"sync"
	"tbo_backend/objects"
)

var (
	creditRiskModel *gocatboost.Catboost
	loadOnce        sync.Once
	loadErr         error
)

func LoadCreditRiskModel() error {
	loadOnce.Do(func() {
		// loading model in memory
		cb, err := gocatboost.FromFile(objects.ConfigObj.CatboostModelsPaths.CreditRiskModel)

		// checking for error
		if err != nil {
			loadErr = err
			return
		}

		// updating booking behaviour model
		creditRiskModel = cb

		fmt.Println("Credit Risk Model Loaded!")
	})

	return loadErr
}
