package credit_risk_model

import "fmt"

func PredictCreditRisk(floatFeatures []float64, catFeatures []string) (float64, error) {
	// checking if model is not in memory
	if creditRiskModel == nil {
		return 0, fmt.Errorf("credit risk model not initialized")
	}

	// returning prediction
	return creditRiskModel.Predict(floatFeatures, catFeatures)
}
