package repositories

import (
	"fmt"
	"tbo_backend/objects"
	"time"

	"github.com/gocql/gocql"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	UpsertLoginOtp(mobile string, otp string) error
	ValidateOtp(mobile string, otp string) (string, error)
}

// ScyllaUserRepository implements UserRepository for ScyllaDB
type ScyllaUserRepository struct {
	session *gocql.Session
}

func NewScyllaUserRepository(session *gocql.Session) UserRepository {
	return &ScyllaUserRepository{
		session: session,
	}
}

func (r *ScyllaUserRepository) UpsertLoginOtp(mobile string, otp string) error {
	// 1. SELECT: Check if user exists
	var count int
	checkQuery := `SELECT count(*) FROM ` + objects.UsersCollection + ` WHERE mobile_number = ? ALLOW FILTERING`

	if err := r.session.Query(checkQuery, mobile).Scan(&count); err != nil {
		return err
	}

	now := time.Now()

	if count == 0 {
		// --- CASE 1: INSERT (User does not exist) ---
		insertQuery := `
            INSERT INTO ` + objects.UsersCollection + ` (
                mobile_number, 
                otp, 
                name, 
                is_active, 
                is_deleted, 
                created_at, 
                updated_at
            ) VALUES (?, ?, '', true, false, ?, ?)`

		fmt.Printf("User %s not found. Creating new record...\n", mobile)
		if err := r.session.Query(insertQuery, mobile, otp, now, now).Exec(); err != nil {
			return err
		}

	} else {
		// --- CASE 2: UPDATE (User exists) ---
		updateQuery := `
            UPDATE ` + objects.UsersCollection + ` 
            SET otp = ?, 
			updated_at = ?,  is_active = true, is_deleted = false
            WHERE mobile_number = ?`

		fmt.Printf("User %s found. Updating OTP...\n", mobile)
		if err := r.session.Query(updateQuery, otp, now, mobile).Exec(); err != nil {
			return err
		}
	}

	return nil
}

func (r *ScyllaUserRepository) ValidateOtp(mobile string, otp string) (string, error) {
	var name string

	// The Query
	query := `
        SELECT name 
        FROM ` + objects.UsersCollection + ` 
        WHERE mobile_number = ? AND otp = ? 
        ALLOW FILTERING`

	err := r.session.Query(query, mobile, otp).Scan(&name)

	if err != nil {
		if err == gocql.ErrNotFound {
			return "", nil // Not found is not an error in this context, just invalid OTP
		}
		// Log unexpected DB errors
		fmt.Println("ScyllaDB Error:", err)
		return "", err
	}

	return name, nil
}
