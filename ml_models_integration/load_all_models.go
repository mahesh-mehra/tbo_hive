package ml_models_integration

import (
	"fmt"
	"tbo_backend/ml_models_integration/agent_quality_model"
	"tbo_backend/ml_models_integration/agent_risk_model"
	"tbo_backend/ml_models_integration/booking_behaviour_model"
	"tbo_backend/ml_models_integration/credit_risk_model"
	"tbo_backend/ml_models_integration/device_risk_model"
	"tbo_backend/ml_models_integration/fraud_history_model"
)

func LoadCatboostModels() bool {
	// loading agent risk model
	err := agent_risk_model.LoadAgentRiskModel()

	// checking for error
	if err != nil {
		fmt.Printf("Failed to load agent risk model: %v\n", err)
		return false
	}

	// loading booking risk model
	err = booking_behaviour_model.LoadBookingBehaviourModel()

	// checking for error
	if err != nil {
		fmt.Printf("Failed to load booking behaviour model: %v\n", err)
		return false
	}

	// loading credit risk model
	err = credit_risk_model.LoadCreditRiskModel()

	// checking for error
	if err != nil {
		fmt.Printf("Failed to load credit risk model: %v\n", err)
		return false
	}

	// loading device risk model
	err = device_risk_model.LoadDeviceRiskModel()

	// checking for error
	if err != nil {
		fmt.Printf("Failed to load device risk model: %v\n", err)
		return false
	}

	// loading fraud history model
	err = fraud_history_model.LoadFraudHistoryModel()

	// checking for error
	if err != nil {
		fmt.Printf("Failed to load fraud history model: %v\n", err)
		return false
	}

	// loading agent quality model
	err = agent_quality_model.LoadAgentQualityModel()

	// checking for error
	if err != nil {
		fmt.Printf("Failed to load agent quality model: %v\n", err)
		return false
	}

	// returning final response
	return true
}
