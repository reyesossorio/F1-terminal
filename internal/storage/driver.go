package storage

import (
	"fmt"

	"github.com/reyesossorio/f1-terminal/internal/domain"
)

type DriverStorage struct {
	Drivers map[int]*domain.DriverInfo
}

func NewDriverStorage() *DriverStorage {
	return &DriverStorage{
		Drivers: make(map[int]*domain.DriverInfo),
	}
}

func (s *DriverStorage) SaveDriverResult(driverResult domain.DriverResult) error {
	val, ok := s.Drivers[driverResult.Position]
	if ok {
		return fmt.Errorf("position %d already saved for driver %d", driverResult.Position, val.DriverNumber)
	}
	s.Drivers[driverResult.Position] = &domain.DriverInfo{
		Dnf:          driverResult.Dnf,
		Dns:          driverResult.Dns,
		Dnq:          driverResult.Dnq,
		DriverNumber: driverResult.DriverNumber,
		GapToLeader:  driverResult.GapToLeader,
	}
	return nil
}

func (s *DriverStorage) SaveDriverInfo(driverInfo domain.Driver) error {
	var found bool
	for _, driver := range s.Drivers {
		if driver.DriverNumber == driverInfo.Number {
			driver.TeamName = driverInfo.TeamName
			driver.Name = driverInfo.Name
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("driver with number %d not found in storage", driverInfo.Number)
	}
	return nil
}

func (s *DriverStorage) GetDrivers() []*domain.DriverInfo {
	driversList := make([]*domain.DriverInfo, 0, len(s.Drivers))

	for _, driver := range s.Drivers {
		driversList = append(driversList, driver)
	}
	return driversList
}
