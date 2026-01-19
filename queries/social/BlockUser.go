package social

import (
	"fmt"
	"tbo_backend/clients"
	"tbo_backend/objects"
	"tbo_backend/utils"
	"time"
)

func BlockUser(friendId *string, userId *string) bool {

	defer utils.HandlePanic()

	uniqueKey := *userId + "_" + *friendId
	uniqueKeyReverse := *friendId + "_" + *userId
	deleteQuery := `DELETE FROM ` + objects.BlockCollection + ` WHERE user_friend_id IN (?, ?)`

	if err := clients.ScyllaSession.Query(deleteQuery, uniqueKey, uniqueKeyReverse).Exec(); err != nil {
		fmt.Println(err)
		return false
	}

	// 1. SELECT: Check if user exists
	var count int
	// CORRECT SYNTAX:
	checkQuery := `SELECT count(*) FROM ` + objects.BlockCollection + ` WHERE user_friend_id = ? ALLOW FILTERING`
	// We use Query().Scan() to get the count
	if err := clients.ScyllaSession.Query(checkQuery, uniqueKey).Scan(&count); err != nil {
		fmt.Println(err)
		return false
	}

	if count == 0 {
		// --- CASE 1: INSERT (User does not exist) ---
		// We set BOTH created_at and updated_at
		insertQuery := `
            INSERT INTO ` + objects.BlockCollection + ` (
			    user_friend_id,
                user_id, 
                friend_id, 
                created_at
            ) VALUES (?, ?, ?, ?)`

		if err := clients.ScyllaSession.Query(insertQuery, uniqueKey, userId, friendId, time.Now()).Exec(); err != nil {
			fmt.Println(err)
			return false
		}
	}

	return true
}

func UnBlockUser(friendId *string, userId *string) bool {

	defer utils.HandlePanic()

	uniqueKey := *userId + "_" + *friendId
	deleteQuery := `DELETE FROM ` + objects.BlockCollection + ` WHERE user_friend_id = ?`

	if err := clients.ScyllaSession.Query(deleteQuery, uniqueKey).Exec(); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func FetchBlockedUserList(userId *string) []objects.BlockedUserDetails {

	defer utils.HandlePanic()

	// 1. FIRST QUERY: Get the list of Friend IDs (Mobile Numbers)
	var blockedUserIds []string
	var friendId string

	selectQuery := `SELECT friend_id FROM ` + objects.BlockCollection + ` WHERE user_id = ? ALLOW FILTERING`

	// FIX: Use Iter() to loop through multiple rows
	iter := clients.ScyllaSession.Query(selectQuery, userId).Iter()
	for iter.Scan(&friendId) {
		blockedUserIds = append(blockedUserIds, friendId)
	}
	if err := iter.Close(); err != nil {
		fmt.Println("Error fetching blocked list:", err)
		return nil
	}

	// If no users are blocked, return empty now to avoid a database error in the next query
	if len(blockedUserIds) == 0 {
		return []objects.BlockedUserDetails{}
	}

	// 2. SECOND QUERY: Fetch details using the array in an IN clause
	var userDetailsList []objects.BlockedUserDetails

	// Assuming 'objects.UserCollection' is your user table name
	// Note: 'IN' operator works best if mobile_number is a Partition Key
	detailsQuery := `
		SELECT name, user_name, profile_photo 
		FROM ` + objects.UsersCollection + ` 
		WHERE mobile_number IN ?`

	iter2 := clients.ScyllaSession.Query(detailsQuery, blockedUserIds).Iter()

	var name, userName, profilePhoto string
	for iter2.Scan(&name, &userName, &profilePhoto) {
		userDetailsList = append(userDetailsList, objects.BlockedUserDetails{
			Name:         name,
			UserName:     userName,
			ProfilePhoto: profilePhoto,
		})
	}

	if err := iter2.Close(); err != nil {
		fmt.Println("Error fetching user details:", err)
		return nil
	}

	return userDetailsList
}
