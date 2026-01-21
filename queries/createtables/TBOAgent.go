package createtables

import (
	"fmt"
	"tbo_backend/clients"
	"tbo_backend/objects"
	"tbo_backend/utils"
)

func CreateTBOAgentsTable() {

	defer utils.HandlePanic()

	// CQL query to create the table
	// 'mobile_number' is the Partition Key (Primary Key)
	query := `
    CREATE TABLE IF NOT EXISTS ` + objects.TBOAgentsCollection + ` (
        agent_id text PRIMARY KEY,
        name text,
        is_active boolean,
        is_deleted boolean,
		is_permanently_deleted boolean,
		business_type text,
		country text,
		city text,
		kyc_verified boolean,
		gst_verified boolean,
		past_blacklist_flag boolean,
		total_bookings bigint,
		successful_bookings_ratio double,
		avg_booking_value double,
		cancellation_rate double,
		refund_rate double,
		no_show_rate double,
		repeated_customer_rate double,
		booking_velocity double,
		credit_limit double,
		avg_credit_limit_utilized double,
		overdue_days_avg double,
		overdue_count bigint,
		max_overdue_days bigint,
		chargeback_count bigint,
		payment_delay_ratio double,
		unique_device_count int,
		ip_country_mismatch_ratio double,
		booking_time_entropy double,
		weekend_booking_ratio double,
		fraud_booking_ratio double,
		manual_review_fail_rate double,
		policy_violation_count int,
		created_at timestamp,
		updated_at timestamp
    )`

	fmt.Println("Creating table 'tbo users' if it does not exist...")

	// Execute the query
	if err := clients.ScyllaSession.Query(query).Exec(); err != nil {
		return
	}

	fmt.Println("Table 'tbo users' is ready.")
}
