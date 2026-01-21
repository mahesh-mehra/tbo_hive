package agent_risk_model

import "fmt"

func PredictAgentRisk(floatFeatures []float64, catFeatures []string) (float64, error) {
	// checking if model is not in memory
	if agentRiskModel == nil {
		return 0, fmt.Errorf("agent risk model not initialized")
	}

	// returning prediction
	return agentRiskModel.Predict(floatFeatures, catFeatures)
}
