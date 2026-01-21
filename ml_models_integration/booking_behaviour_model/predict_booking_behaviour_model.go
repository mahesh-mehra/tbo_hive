package booking_behaviour_model

import "fmt"

func PredictBookingRisk(floatFeatures []float64, catFeatures []string) (float64, error) {
	// checking if model is not in memory
	if bookingBehaviourModel == nil {
		return 0, fmt.Errorf("booking behaviour model not initialized")
	}

	// returning prediction
	return bookingBehaviourModel.Predict(floatFeatures, catFeatures)
}
