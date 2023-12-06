package main

import (
	"fmt"
	"log"
	"os"
	usecases "videos-with-subtitle-player/useCases"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rootPath := os.Getenv("ROOT_PATH")

	content := usecases.GetFileTree(rootPath)
	fmt.Printf("content %v\n", content)

}
