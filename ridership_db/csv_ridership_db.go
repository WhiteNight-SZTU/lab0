package ridershipDB

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type CsvRidershipDB struct {
	idIdxMap      map[string]int
	csvFile       *os.File
	csvReader     *csv.Reader
	num_intervals int
}

// Close implements RidershipDB.
func (c *CsvRidershipDB) Close() error {
	//panic("unimplemented")
	err := c.csvFile.Close()
	if err != nil {
		return err
	}
	return nil
}

// GetRidership implements RidershipDB.
func (c *CsvRidershipDB) GetRidership(lineId string) ([]int64, error) {
	//panic("unimplemented")
	boardings := make([]int64, c.num_intervals)

	for {
		records, err := c.csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if records[0] == lineId {
			sum, err := strconv.ParseInt(records[4], 10, 64)
			if err != nil {
				return nil, err
			}
			boardings[c.idIdxMap[records[2]]] += sum
		}

	}

	return boardings, nil
}

func (c *CsvRidershipDB) Open(filePath string) error {
	c.num_intervals = 9

	// Create a map that maps MBTA's time period ids to indexes in the slice
	c.idIdxMap = make(map[string]int)
	for i := 1; i <= c.num_intervals; i++ {
		timePeriodID := fmt.Sprintf("time_period_%02d", i)
		c.idIdxMap[timePeriodID] = i - 1
	}

	// create csv reader
	csvFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	c.csvFile = csvFile
	c.csvReader = csv.NewReader(c.csvFile)

	return nil
}

// TODO: some code goes here
// Implement the remaining RidershipDB methods
