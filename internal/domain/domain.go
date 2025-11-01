package domain

import (
	"encoding/json"
	"fmt"
)

type DriverSessionResults struct {
	Number      int
	Name        string
	Team        string
	GapToLeader float64 // floatnumber, -x for number of laps
}

type SessionResult struct {
	SessionName string
	SessionType string
}

type Session struct {
	SessionKey       int    `json:"session_key"`
	Location         string `json:"location"`
	DateStart        string `json:"date_start"`
	DateEnd          string `json:"date_end"`
	SessionType      string `json:"session_type"`
	SessionName      string `json:"session_name"`
}

type Driver struct {
	TeamName string `json:"team_name"`
	Name     string `json:"name_acronym"`
	Number   int    `json:"driver_number"`
}
type DriverResult struct {
	Dnf          bool `json:"dnf"`
	Dns          bool `json:"dns"`
	Dnq          bool `json:"dnq"`
	DriverNumber int  `json:"driver_number"`
	GapToLeader  Gap  `json:"gap_to_leader"`
	Position     int  `json:"position"`
}

type Laps struct {
	LapDuration float64 `json:"lap_duration"`
	LapNumber   int     `json:"lap_number"`
}

type DriverInfo struct {
	Name         string
	TeamName     string
	DriverNumber int
	GapToLeader  Gap
	Dnf          bool
	Dns          bool
	Dnq          bool
}

type Gap struct {
	Value interface{}
}

func (g *Gap) UnmarshalJSON(data []byte) error {
	// Intentar decodificar como string
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		g.Value = s
		return nil
	}

	// Intentar decodificar como float64
	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		g.Value = f
		return nil
	}

	// Intentar decodificar como array (slice genérico)
	var arr []float64
	if err := json.Unmarshal(data, &arr); err == nil {
		g.Value = arr
		return nil
	}

	// Si nada funciona, guardar el valor bruto
	g.Value = string(data)
	return nil
}

func (g Gap) String() string {
	return fmt.Sprintf("%v", g.Value)
}
