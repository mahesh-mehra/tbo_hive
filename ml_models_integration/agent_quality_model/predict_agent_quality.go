package agent_quality_model

import "fmt"

func PredictAgentQuality(floatFeatures []float64) (float64, error) {
	// checking if model is not in memory
	if agentQualityModel == nil {
		return 0, fmt.Errorf("agent quality model not initialized")
	}

	// returning prediction
	return agentQualityModel.Predict(floatFeatures, []string{})
}
