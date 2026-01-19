package profile

import (
	"fmt"
	"tbo_backend/clients"
	"tbo_backend/objects"
	"tbo_backend/utils"
	"time"
)

func UpdateProfile(name *string, userName *string, userId *string) bool {

	defer utils.HandlePanic()

	updateQuery := `
		UPDATE ` + objects.UsersCollection + ` 
		SET name = ?, user_name = ?,
		updated_at = ? WHERE mobile_number = ?`

	if err := clients.ScyllaSession.Query(updateQuery, name, userName, time.Now(), userId).Exec(); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
