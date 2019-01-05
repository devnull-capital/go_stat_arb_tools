// +build unit

package statsArb

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func fetchFemaleBirthsData() ([]float64, error) {
	csvFile, err := os.Open("test_data/daily-total-female-births.csv")
	if err != nil {
		log.Printf("err opening births file\n%v", err)
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	// note: read the first line...
	_, err = reader.Read()
	if err != nil {
		log.Printf("err reading births file\n%v", err)
		return nil, err
	}

	var data []float64
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("err reading csv\n%v", err)
			return nil, err
		}

		f, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			log.Printf("err parsing %s to float\n%v", line[1], err)
			return nil, err
		}

		data = append(data, f)
	}

	return data, nil
}

func fetchTechStockData() ([2][]float64, error) {
	var data [2][]float64

	aaplFile, err := os.Open("test_data/aapl_20080601_20131231.csv")
	if err != nil {
		log.Printf("err opening aapl file\n%v", err)
		return data, err
	}
	aaplReader := csv.NewReader(bufio.NewReader(aaplFile))
	// note: read the first line...
	_, err = aaplReader.Read()
	if err != nil {
		log.Printf("err reading aapl file\n%v", err)
		return data, err
	}
	googFile, err := os.Open("test_data/goog_20080601_20131231.csv")
	if err != nil {
		log.Printf("err opening goog file\n%v", err)
		return data, err
	}
	googReader := csv.NewReader(bufio.NewReader(googFile))
	// note: read the first line...
	_, err = googReader.Read()
	if err != nil {
		log.Printf("err opening goog file\n%v", err)
		return data, err
	}

	for {
		line, err := aaplReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("err reading aapl csv\n%v", err)
			return data, err
		}

		f, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			log.Printf("err parsing %s to float\n%v", line[1], err)
			return data, err
		}

		data[0] = append(data[0], f)
	}
	for {
		line, err := googReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("err reading goog csv\n%v", err)
			return data, err
		}

		f, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			log.Printf("err parsing %s to float\n%v", line[1], err)
			return data, err
		}

		data[1] = append(data[1], f)
	}

	return data, nil
}
