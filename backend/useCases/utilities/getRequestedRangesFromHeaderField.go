package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

func GetRequestedRangesFromHeaderField(rangeHeaderWithPrefix string, chunkSize int64) (start int64, end int64) {
	if rangeHeaderWithPrefix == "" {
		return 0, chunkSize
	}

	rangeHeaderWithoutPrefix := strings.TrimPrefix(rangeHeaderWithPrefix, "bytes=")
	rangeParts := strings.Split(rangeHeaderWithoutPrefix, "-")
	if len(rangeParts) != 2 {
		return 0, 0
	}

	start, err := strconv.ParseInt(rangeParts[0], 10, 64)
	if err != nil {
		fmt.Println("start", err.Error())
		return 0, 0
	}

	end = 0

	if rangeParts[1] == "" {
		return start, start + chunkSize
	}

	end, err = strconv.ParseInt(rangeParts[1], 10, 64)
	if err != nil {
		fmt.Println("end", err.Error())
		return 0, 0
	}

	if end == 0 {
		end = start + chunkSize
	}

	return start, end
}
