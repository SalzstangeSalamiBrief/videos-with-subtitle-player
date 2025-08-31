package models

type ImageCLICommand struct {
	Command string
	Arg     string
}

func ConvertImageMagickCommandsArrayToArgumentsArray(commands []ImageCLICommand) []string {
	result := []string{}
	for _, command := range commands {
		if command.Command != "" {
			result = append(result, command.Command)
		}

		if command.Arg != "" {
			result = append(result, command.Arg)
		}
	}

	return result
}
