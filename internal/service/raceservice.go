package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/reyesossorio/f1-terminal/internal/domain"
	"github.com/reyesossorio/f1-terminal/internal/storage"
)

type RaceService struct {
	drivers  *storage.DriverStorage
	sessions *storage.SessionStorage
}

func NewRaceService(storage *storage.DriverStorage, sessions *storage.SessionStorage) *RaceService {
	return &RaceService{
		drivers:  storage,
		sessions: sessions,
	}
}

func (r *RaceService) SaveDriversResults(driversResults []domain.DriverResult) error {
	for _, driverResult := range driversResults {
		err := r.drivers.SaveDriverResult(driverResult)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RaceService) SaveDriversInfo(driversInfo []domain.Driver) error {
	for _, driverInfo := range driversInfo {
		err := r.drivers.SaveDriverInfo(driverInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RaceService) LazyDriversInfo(driversNumbers []int) error {
	query := fmt.Sprintf("https://api.openf1.org/v1/drivers?session_key=%s", "latest")
	for _, number := range driversNumbers {
		query += fmt.Sprintf("&driver_number=%d", number)
	}
	resp, err := http.Get(query)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("LazyDriversInfo error code %d", resp.StatusCode)
	}
	var drivers []domain.Driver
	if err := json.NewDecoder(resp.Body).Decode(&drivers); err != nil {
		fmt.Println(err)
	}
	return r.SaveDriversInfo(drivers)
}

func (r *RaceService) SaveLatestSession() error {
	resp, err := http.Get("https://api.openf1.org/v1/sessions?session_key=latest")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("SaveLatestSession error code %d", resp.StatusCode)
	}

	var session []domain.Session
	if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
		return err
	}
	err = r.sessions.SaveSession(session[0])
	if err != nil {
		return err
	}
	r.sessions.SetCurSession(session[0].SessionKey)
	return nil
}

func (r *RaceService) LazyDriversRaceResults(position int, greater bool) error {
	query := fmt.Sprintf("https://api.openf1.org/v1/session_result?session_key=%d", r.sessions.GetCurSession())
	if greater {
		query += fmt.Sprintf("&position%%3E%%3D%d", position) // >=
	} else {
		query += fmt.Sprintf("&position%%3C%%3D%d", position) // <=
	}

	fmt.Print(query)

	resp, err := http.Get(query)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("LazyDriversRaceResults error code %d", resp.StatusCode)
	}

	var driversResults []domain.DriverResult
	if err := json.NewDecoder(resp.Body).Decode(&driversResults); err != nil {
		fmt.Println(err)
	}

	err = r.SaveDriversResults(driversResults)
	if err != nil {
		return err
	}
	return nil
}

func (r *RaceService) GetLastLapTime(driverNumber int) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.openf1.org/v1/laps?session_key=%d&driver_number=%d", r.sessions.GetCurSession(), driverNumber))
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != http.StatusOK {
		return -1, fmt.Errorf("GetLastLapTime error code %d", resp.StatusCode)
	}
	var laps []domain.Laps
	if err := json.NewDecoder(resp.Body).Decode(&laps); err != nil {
		return -1, err
	}
	return laps[len(laps)-1].LapDuration, nil
}

func (r *RaceService) GetDriversInSession() []*domain.DriverInfo {
	return r.drivers.GetDrivers()
}
