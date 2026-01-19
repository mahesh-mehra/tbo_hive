package createtables

import (
	"fmt"
	"tbo_backend/clients"
	"tbo_backend/objects"
	"tbo_backend/utils"
)

func CreateUserTable() {

	defer utils.HandlePanic()

	// CQL query to create the table
	// 'mobile_number' is the Partition Key (Primary Key)
	query := `
    CREATE TABLE IF NOT EXISTS ` + objects.UsersCollection + ` (
        mobile_number text PRIMARY KEY,
        name text,
		user_name text,
		profile_photo text,
        otp text,
        is_active boolean,
        is_deleted boolean,
		is_permanently_deleted boolean,
		created_at timestamp,
		updated_at timestamp
    )`

	fmt.Println("Creating table 'users' if it does not exist...")

	// Execute the query
	if err := clients.ScyllaSession.Query(query).Exec(); err != nil {
		return
	}

	fmt.Println("Table 'users' is ready.")
}

func AlterUserTable() {

	defer utils.HandlePanic()

	// CQL query to create the table
	// 'mobile_number' is the Partition Key (Primary Key)
	query := `
    	ALTER TABLE ` + objects.UsersCollection + ` ADD profile_photo text
    `

	// Execute the query
	if err := clients.ScyllaSession.Query(query).Exec(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Table 'users' is altered.")
}
