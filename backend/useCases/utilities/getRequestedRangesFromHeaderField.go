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
	if stringifiedStart == "" && stringifiedEnd == "" {
		return 0, 0
	}

	start, err := getStart(stringifiedStart)
	if err != nil {
		fmt.Println("start", err.Error())
		return 0, 0
	}

	end, err = getEnd(stringifiedEnd, start, fileSize, chunkSize)
	if err != nil {
		fmt.Println("end", err.Error())
		return 0, 0
	}

	return start, end
}

func getStringifiedRange(rangeHeaderWithPrefix string) (stringifiedStart string, stringifiedEnd string) {
	rangeHeaderWithoutPrefix := strings.TrimPrefix(rangeHeaderWithPrefix, "bytes=")
	rangeParts := strings.Split(rangeHeaderWithoutPrefix, "-")
	if len(rangeParts) != 2 {
		return "", ""
	}

	return rangeParts[0], rangeParts[1]
}

func getStart(stringifiedStart string) (int64, error) {
	start, err := strconv.ParseInt(stringifiedStart, 10, 64)
	if err != nil {
		return 0, err
	}

	return start, nil
}

func getEnd(stringifiedEnd string, start int64, fileSize int64, chunkSize int64) (int64, error) {
	if stringifiedEnd == "" {
		return start + chunkSize, nil
	}

	end, err := strconv.ParseInt(stringifiedEnd, 10, 64)
	if err != nil {
		return 0, err
	}

	if end == 0 {
		end = start + chunkSize
	}

	if end > fileSize {
		end = fileSize
	}

	return end, nil
}
