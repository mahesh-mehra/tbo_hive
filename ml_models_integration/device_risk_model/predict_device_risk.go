package device_risk_model

import "fmt"

func PredictCreditRisk(floatFeatures []float64, catFeatures []string) (float64, error) {
	// checking if model is not in memory
	if deviceRiskModel == nil {
		return 0, fmt.Errorf("device risk model not initialized")
	}

	// returning prediction
	return deviceRiskModel.Predict(floatFeatures, catFeatures)
}
