package device_risk_model

import (
	"fmt"
	"github.com/korsajan/gocatboost"
	"sync"
	"tbo_backend/objects"
)

var (
	deviceRiskModel *gocatboost.Catboost
	loadOnce        sync.Once
	loadErr         error
)

func LoadDeviceRiskModel() error {
	loadOnce.Do(func() {
		// loading model in memory
		cb, err := gocatboost.FromFile(objects.ConfigObj.CatboostModelsPaths.DeviceRiskModel)

		// checking for error
		if err != nil {
			loadErr = err
			return
		}

		// updating booking behaviour model
		deviceRiskModel = cb

		fmt.Println("Device Risk Model Loaded!")
	})

	return loadErr
}
