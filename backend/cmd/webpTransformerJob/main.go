package main

import (
	"backend/internal/configuration/webpTransformer"
	"backend/pkg/services/imageConverter/webp"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	initializedConfiguration, configurationError := webpTransformer.NewWebpTransformerConfiguration()
	if configurationError != nil {
		log.Fatal(configurationError)
	}

	osExitChannel := make(chan os.Signal, 1)
	signal.Notify(osExitChannel, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(time.Duration(initializedConfiguration.GetExecutionIntervalInMinutes()) * time.Minute)
	quitChannel := make(chan struct{}, 1)

	var conversionMutex sync.Mutex

	for {
		select {
		case <-ticker.C:
			if !conversionMutex.TryLock() {
				log.Printf("Conversion is already running. Skip duplicate execution.")
				continue
			}

			log.Printf("Start webp transformer at '%v'\n", time.Now())
			err := webp.ExecuteWebpConversion(webp.ExecuteWebpConversionConfiguration{RootPath: initializedConfiguration.GetRootPath(), ShouldDeleteNonWebpImages: initializedConfiguration.GetShouldDeleteNonWebpImages()})
			if err != nil {
				log.Println(err)
				conversionMutex.Unlock()
				quitChannel <- struct{}{}
			}

			log.Printf("Finish webp transformer at '%v'\n", time.Now())
			conversionMutex.Unlock()

		case <-osExitChannel:
		case <-quitChannel:
			signal.Stop(osExitChannel)
			close(osExitChannel)
			close(quitChannel)
			ticker.Stop()
			os.Exit(0)
		}
	}
}
