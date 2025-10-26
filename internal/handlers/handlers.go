package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/reyesossorio/f1-terminal/internal/domain"
)

func FetchLatestResults() {
	resp, err := http.Get("https://api.openf1.org/v1/sessions?session_key=latest")
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
}
