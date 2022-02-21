package file

import (
	"encoding/csv"
	"fare-estimation/internal/ride"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func ReadCsv(filePath string) map[int][]ride.Ride {
	rides := make(map[int][]ride.Ride)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(f)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	for _, line := range records {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		timeValue, err := strconv.ParseInt(line[3], 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		tm := time.Unix(timeValue, 0)
		rides[id] = append(rides[id], ride.Ride{
			ID:   id,
			Lat:  line[1],
			Long: line[2],
			Time: tm,
		})
	}

	return rides
}
