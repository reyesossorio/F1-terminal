package main

import (
	"fmt"

	"github.com/reyesossorio/f1-terminal/internal/service"
	"github.com/reyesossorio/f1-terminal/internal/storage"
)

func main() {
	drivers := storage.NewDriverStorage()
	sessions := storage.NewSessionStorage()

	service := service.NewRaceService(drivers, sessions)

	err := service.SaveLatestSession()
	if err != nil {
		fmt.Println(err)
	}

	err = service.LazyDriversRaceResults(10, false)
	if err != nil {
		fmt.Println(err)
	}
	driversInSession := service.GetDriversInSession()
	for _, driver := range driversInSession{
		fmt.Println(driver)
	}
}
