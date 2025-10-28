package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/reyesossorio/f1-terminal/internal/domain"
)

func FetchLatestResults() {
	// get latest session
	resp, err := http.Get(fmt.Sprintf("https://api.openf1.org/v1/sessions?session_key=%s", "latest"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(err)
	}

	var session []domain.Session
	if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
		fmt.Println(err)
	}
	fmt.Println(session)

	// get drivers results for session
	resp, err = http.Get(fmt.Sprintf("https://api.openf1.org/v1/session_result?session_key=%d", session[0].SessionKey))
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(err)
	}
	var driversResults []domain.DriverResult
	if err := json.NewDecoder(resp.Body).Decode(&driversResults); err != nil {
		fmt.Println(err)
	}
	fmt.Println(driversResults)

	// get drivers info
	resp, err = http.Get(fmt.Sprintf("https://api.openf1.org/v1/drivers?session_key=%d", session[0].SessionKey))
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(err)
	}
	var drivers []domain.Driver
	if err := json.NewDecoder(resp.Body).Decode(&drivers); err != nil {
		fmt.Println(err)
	}
	fmt.Println(drivers)

	//get lap times
	resp, err = http.Get(fmt.Sprintf("https://api.openf1.org/v1/laps?session_key=%d&driver_number=%d", session[0].SessionKey, driversResults[0].DriverNumber))
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(err)
	}
	var laps []domain.Laps
	if err := json.NewDecoder(resp.Body).Decode(&laps); err != nil {
		fmt.Println(err)
	}
	fmt.Println(laps)
}
