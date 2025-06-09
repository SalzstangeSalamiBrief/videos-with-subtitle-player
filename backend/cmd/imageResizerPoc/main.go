package imageResizerPoc

import (
	"fmt"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type Command struct {
	command string
	arg     string
}

const imageSourceFolder = "./images/lossy"
const imageSource = "thumbnail.jpg"

func main() {
	commands := [][]Command{
		[]Command{{command: "-quality", arg: "10"}},
		[]Command{{command: "-quality", arg: "20"}},
		[]Command{{command: "-quality", arg: "50"}},
		[]Command{{command: "-quality", arg: "70"}},
		[]Command{{command: "-quality", arg: "80"}},
		[]Command{{command: "-quality", arg: "90"}},
		[]Command{{command: "-resize", arg: "640x"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "10"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "20"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "50"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "70"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "80"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "90"}},
	}

	for _, command := range commands {
		handleCommand(command)
	}

}

func handleCommand(commandArguments []Command) {
	fmt.Printf("Start for commands %v\n", commandArguments)
	extension := filepath.Ext(imageSource)
	filename := strings.TrimSuffix(imageSource, extension)
	inputImageRelativePath := path.Join(imageSourceFolder, imageSource)
	fmt.Println("inputImageRelativePath", inputImageRelativePath)
	resizedImageName := fmt.Sprintf("%v%v%v", filename, getFilenameFromCommands(commandArguments), extension)
	fmt.Println("resizedImageName", resizedImageName)
	resizedImagePath := path.Join(imageSourceFolder, resizedImageName)

	commandArgsArray := convertCommandsArrayToStringArray(commandArguments)
	command := exec.Command("magick", inputImageRelativePath)
	command.Args = append(command.Args, commandArgsArray...)
	command.Args = append(command.Args, resizedImagePath)
	err := command.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Finish for commands %v\n", commandArguments)
}

func getFilenameFromCommands(commands []Command) string {
	result := ""
	for _, command := range commands {
		result += fmt.Sprintf("%v%v", command.command, command.arg)
	}

	return result
}

func convertCommandsArrayToStringArray(commands []Command) []string {
	var result []string
	for _, command := range commands {
		result = append(result, command.command)
		result = append(result, command.arg)
	}

	return result
}
