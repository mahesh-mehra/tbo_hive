package profile

import (
	"fmt"
	"tbo_backend/clients"
	"tbo_backend/objects"
	"tbo_backend/utils"
	"time"
)

func UpdateProfilePhoto(userId *string, imageName *string) bool {

	defer utils.HandlePanic()

	updateQuery := `
		UPDATE ` + objects.UsersCollection + ` 
		SET profile_photo = ?,
		updated_at = ? WHERE mobile_number = ?`

	if err := clients.ScyllaSession.Query(updateQuery, imageName, time.Now(), userId).Exec(); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
