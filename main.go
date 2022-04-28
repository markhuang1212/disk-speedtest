package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

func parseSizeArg(input string) int {
	val := []byte{}
	unit := []byte{}

	i := 0
	for ; i < len(input); i++ {
		if unicode.IsDigit(rune(input[i])) {
			val = append(val, input[i])
		} else {
			break
		}
	}

	for ; i < len(input); i++ {
		unit = append(unit, input[i])
	}

	if len(unit) == 0 {
		ret, _ := strconv.Atoi(string(val))
		return ret
	} else if len(unit) == 1 {
		if unit[0] == 'M' {
			ret, _ := strconv.Atoi(string(val))
			return ret
		} else if unit[0] == 'G' {
			ret, _ := strconv.Atoi(string(val))
			return ret * 1024
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func main() {
	sizeFlag := flag.String("size", "1G", "how many data to write in total")
	chunkFlag := flag.String("chunk", "64M", "how many data to write in one operation")
	flag.Parse()

	sizeToWrite := parseSizeArg(*sizeFlag)
	chunkSize := parseSizeArg(*chunkFlag)

	if sizeToWrite <= 0 || chunkSize <= 0 {
		panic("Invalid Size")
	}

	numChunk := sizeToWrite / chunkSize

	b := make([]byte, chunkSize*1024*1024)

	fd, err := os.Create("tmpfile")
	if err != nil {
		panic("Cannot Create Temp File")
	}

	start := time.Now()
	for i := 0; i < numChunk; i++ {
		fd.Write(b)
	}
	elapsed := time.Since(start).Nanoseconds()
	fd.Close()
	os.Remove("tmpfile")

	fmt.Printf("Elapsed Time: %dns \n", elapsed)
	fmt.Printf("Write Speed: %.2fMB/s", float32(sizeToWrite)/(float32(elapsed)/1e9))
}
