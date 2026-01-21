package fraud_history_model

import "fmt"

func PredictFraudRisk(floatFeatures []float64, catFeatures []string) (float64, error) {
	// checking if model is not in memory
	if fraudHistoryModel == nil {
		return 0, fmt.Errorf("fraud history model not initialized")
	}

	// returning prediction
	return fraudHistoryModel.Predict(floatFeatures, catFeatures)
}
