package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

func GetRequestedRangesFromHeaderField(rangeHeaderWithPrefix string, chunkSize int64, fileSize int64) (start int64, end int64) {
	if rangeHeaderWithPrefix == "" {
		return 0, chunkSize
	}

	stringifiedStart, stringifiedEnd := getStringifiedRange(rangeHeaderWithPrefix)
	start, err := getStart(stringifiedStart)
	if err != nil {
		fmt.Println("start", err.Error())
		return 0, 0
	}

	end, err = getEnd(getEndInput{stringifiedEnd, start, fileSize, chunkSize})
	if err != nil {
		fmt.Println("end", err.Error())
		return 0, 0
	}

	return start, end
}

func getStringifiedRange(rangeHeaderWithPrefix string) (stringifiedStart string, stringifiedEnd string) {
	stringifiedStart = ""
	stringifiedEnd = ""
	rangeHeaderWithoutPrefix := strings.TrimPrefix(rangeHeaderWithPrefix, "bytes=")
	rangeParts := strings.Split(rangeHeaderWithoutPrefix, "-")

	if len(rangeParts) == 1 {
		stringifiedStart = rangeParts[0]
	}

	if len(rangeParts) == 2 {
		stringifiedStart = rangeParts[0]
		stringifiedEnd = rangeParts[1]
	}

	return stringifiedStart, stringifiedEnd
}

func getStart(stringifiedStart string) (int64, error) {
	if stringifiedStart == "" {
		return 0, nil
	}

	start, err := strconv.ParseInt(stringifiedStart, 10, 64)
	if err != nil {
		return 0, err
	}

	return start, nil
}

type getEndInput struct {
	stringifiedEnd string
	start          int64
	fileSize       int64
	chunkSize      int64
}

func getEnd(input getEndInput) (int64, error) {
	if input.stringifiedEnd == "" {
		return input.start + input.chunkSize, nil
	}

	end, err := strconv.ParseInt(input.stringifiedEnd, 10, 64)
	if err != nil {
		return 0, err
	}

	if end == 0 {
		end = input.start + input.chunkSize
	}

	if end > input.fileSize {
		end = input.fileSize
	}

	return end, nil
}
