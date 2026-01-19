package profile

import (
	"tbo_backend/clients"
	"tbo_backend/objects"
)

func FetchProfile(userId *string) (*objects.UserProfileData, error) {

	var user objects.UserProfileData

	query := `SELECT name, user_name, profile_photo FROM ` + objects.UsersCollection + ` WHERE mobile_number = ?`

	if err := clients.ScyllaSession.Query(query, userId).Scan(&user.Name, &user.UserName, &user.ProfilePhoto); err != nil {
		return nil, err
	}

	return &user, nil
}
