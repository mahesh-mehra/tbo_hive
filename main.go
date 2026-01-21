package main

import (
	"fmt"
	"tbo_backend/applications"
	"tbo_backend/ml_models_integration"
	"tbo_backend/queries/createtables"
	"tbo_backend/utils"
)

func main() {

	// reading the local.json file
	fmt.Println("Loading config file...")
	utils.LoadConfig()

	// connect scylla
	applications.ConnectScylla()

	// connect kafka
	applications.ConnectKafkaP()

	// connect redis
	applications.ConnectRedis()

	// create scyladb tables
	createtables.CreateUserTable()
	createtables.CreateFollowTable()
	createtables.BlockUserTable()
	createtables.CreateTBOAgentsTable()

	// loading model in memory
	ml_models_integration.LoadCatboostModels()

	// connect to http host fiber framework
	applications.ConnectHttp()
}
