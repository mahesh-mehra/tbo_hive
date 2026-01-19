package social

import (
	"fmt"
	"tbo_backend/clients"
	"tbo_backend/objects"
	"tbo_backend/utils"
	"time"
)

func FollowRequest(friendId *string, userId *string) bool {

	defer utils.HandlePanic()

	uniqueKey := *userId + "_" + *friendId
	// 1. SELECT: Check if user exists
	var count int
	// CORRECT SYNTAX:
	checkQuery := `SELECT count(*) FROM ` + objects.FollowCollection + ` WHERE user_friend_id = ? ALLOW FILTERING`
	// We use Query().Scan() to get the count
	if err := clients.ScyllaSession.Query(checkQuery, uniqueKey).Scan(&count); err != nil {
		fmt.Println(err)
		return false
	}

	if count == 0 {
		// --- CASE 1: INSERT (User does not exist) ---
		// We set BOTH created_at and updated_at
		insertQuery := `
            INSERT INTO ` + objects.FollowCollection + ` (
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

func UnFollowRequest(friendId *string, userId *string) bool {

	defer utils.HandlePanic()

	// Safety check for nil pointers
	if userId == nil || friendId == nil {
		return false
	}

	uniqueKey := *userId + "_" + *friendId

	// Direct DELETE query
	// We do NOT need to check if it exists first.
	deleteQuery := `DELETE FROM ` + objects.FollowCollection + ` WHERE user_friend_id = ?`

	// Execute
	if err := clients.ScyllaSession.Query(deleteQuery, uniqueKey).Exec(); err != nil {
		fmt.Println("Unfollow Error:", err)
		return false
	}

	return true
}
