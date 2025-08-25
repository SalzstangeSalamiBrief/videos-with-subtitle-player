package main

import (
	"backend/internal/config"
	"backend/pkg/services/imageConverter/webp"
	"github.com/go-co-op/gocron/v2"
	"log"
)

func main() {
	// TODO ADD CUSTOM CONFIGURATION => SET ALL VALUES USING CONFIGURATION
	initializedConfiguration := config.InitializeConfiguration()

	scheduler, schedulerError := gocron.NewScheduler()
	if schedulerError != nil {
		log.Fatal(schedulerError)
	}

	//defer func() { _ = scheduler.Shutdown() }()

	_, webpJobError := scheduler.NewJob(
		gocron.DailyJob(1,
			gocron.NewAtTimes(gocron.NewAtTime(16, 40, 0)),
		),
		gocron.NewTask(func() {
			// TODO DOES NOT HIT
			err := webp.ExecuteWebpConversion(webp.ExecuteWebpConversionConfiguration{RootPath: initializedConfiguration.RootPath, ShouldDeleteNonWebpImages: true})
			if err != nil {
				log.Fatal(err)
			}
		}),
	)

	if webpJobError != nil {
		log.Fatal(webpJobError)
	}

	scheduler.Start()
}
