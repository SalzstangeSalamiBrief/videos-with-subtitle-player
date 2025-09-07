package main

import (
	"backend/internal/config/webpTransformer"
	"backend/pkg/services/imageConverter/webp"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// TODO ADD CUSTOM CONFIGURATION => SET ALL VALUES USING CONFIGURATION
	initializedConfiguration := webpTransformer.NewWebpTransformerConfiguration()

	osExitChannel := make(chan os.Signal, 1)
	signal.Notify(osExitChannel, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(time.Duration(initializedConfiguration.GetExecutionIntervalInMinutes()) * time.Minute)
	quitChannel := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			log.Printf("Start webp transformer at '%v'\n", time.Now())
			err := webp.ExecuteWebpConversion(webp.ExecuteWebpConversionConfiguration{RootPath: initializedConfiguration.GetRootPath(), ShouldDeleteNonWebpImages: initializedConfiguration.GetShouldDeleteNonWebpImages()})
			if err != nil {
				log.Println(err)
				quitChannel <- struct{}{}
			}

			log.Printf("Finish webp transformer at '%v'\n", time.Now())
		case <-osExitChannel:
		case <-quitChannel:
			ticker.Stop()
			os.Exit(0)
		}
	}
}
