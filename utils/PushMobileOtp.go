package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"tbo_backend/objects"
	"time"
)

// PushMobileOtp sends OTP using Twilio-style API. Returns true on success.
func PushMobileOtp(contact string, otp string) bool {
	defer HandlePanic()

	// Basic config checks
	twURL := strings.TrimSpace(objects.ConfigObj.Twillio.Url)
	accountSid := strings.TrimSpace(objects.ConfigObj.Twillio.Sid)
	authToken := strings.TrimSpace(objects.ConfigObj.Twillio.Key)
	from := strings.TrimSpace(objects.ConfigObj.Twillio.Contact)

	if twURL == "" {
		fmt.Println("Twilio URL not configured (Twillio.Url is empty)")
		return false
	}
	// Ensure scheme is present (add https if user forgot)
	if !strings.HasPrefix(twURL, "http://") && !strings.HasPrefix(twURL, "https://") {
		twURL = "https://" + twURL
	}

	if accountSid == "" || authToken == "" {
		fmt.Println("Twilio credentials missing (Sid or Key)")
		return false
	}
	if from == "" {
		fmt.Println("Twilio 'From' number missing in config (Twillio.Contact)")
		return false
	}
	if contact == "" {
		fmt.Println("Destination contact is empty")
		return false
	}

	// Build form data
	formData := url.Values{
		"To":   {"+91" + contact},
		"From": {from},
		"Body": {fmt.Sprintf("Your verification code is %s, TBO Hive", otp)},
	}
	bodyStr := formData.Encode()
	req, err := http.NewRequest("POST", twURL, bytes.NewBufferString(bodyStr))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(accountSid, authToken)

	// Use context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// Read response body for debugging
	respBody, _ := io.ReadAll(resp.Body)
	respText := strings.TrimSpace(string(respBody))

	// Twilio produces 201 Created for successful message creation; accept 200 too if relevant
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		fmt.Printf("Error in response: status=%s code=%d body=%s\n", resp.Status, resp.StatusCode, respText)
		return false
	}

	// Optionally, parse JSON response to confirm "sid" or "status" exists.
	fmt.Println("OTP sent successfully. Response:", respText)
	return true
}
