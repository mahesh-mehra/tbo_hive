package createtables

import (
	"fmt"
	"tbo_backend/clients"
	"tbo_backend/objects"
	"tbo_backend/utils"
)

func CreateFollowTable() {

	defer utils.HandlePanic()

	// CQL query to create the table
	// 'user_friend_id' is the Partition Key (Primary Key)
	query := `
    CREATE TABLE IF NOT EXISTS ` + objects.FollowCollection + ` (
		user_friend_id text PRIMARY KEY,
        user_id text,
        friend_id text,
		created_at timestamp
    )`

	fmt.Println("Creating table 'follow' if it does not exist...")

	// Execute the query
	if err := clients.ScyllaSession.Query(query).Exec(); err != nil {
		return
	}

	fmt.Println("Table 'follow' is ready.")
}
