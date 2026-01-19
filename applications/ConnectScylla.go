package applications

import (
	"log"
	"time"

	"tbo_backend/clients"
	"tbo_backend/objects"
	"tbo_backend/utils"

	"github.com/gocql/gocql"
)

// ScyllaSession moved to clients package

func ConnectScylla() {

	defer utils.HandlePanic()

	// ScyllaDB cluster nodes
	cluster := gocql.NewCluster(
		objects.ConfigObj.Scylla,
	)

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: objects.ConfigObj.ScyllaUsername,
		Password: objects.ConfigObj.ScyllaPassword,
	}

	// Keyspace (must already exist)
	cluster.Keyspace = objects.ConfigObj.ScyllaNamespace

	// Consistency & performance tuning
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second

	// Recommended settings for ScyllaDB
	cluster.NumConns = 5
	cluster.ReconnectInterval = time.Second * 5
	cluster.DisableInitialHostLookup = false

	// Create session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Unable to connect to ScyllaDB: %v", err)
	}

	clients.ScyllaSession = session

}
