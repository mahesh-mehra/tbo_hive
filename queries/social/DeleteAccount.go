package social

import (
	"fmt"
	"tbo_backend/clients"
	"tbo_backend/objects"
	"tbo_backend/utils"
)

func DeleteAccount(userId *string) bool {

	defer utils.HandlePanic()

	insertQuery := `
		UPDATE ` + objects.UsersCollection + ` 
		SET is_deleted = true
		WHERE user_id = ?`

	if err := clients.ScyllaSession.Query(insertQuery, userId).Exec(); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func DeactivateAccount(userId *string) bool {

	defer utils.HandlePanic()

	insertQuery := `
		UPDATE ` + objects.UsersCollection + ` 
		SET is_active = false
		WHERE user_id = ?`

	if err := clients.ScyllaSession.Query(insertQuery, userId).Exec(); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
