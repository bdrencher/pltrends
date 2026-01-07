package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type repoRecord struct {
	Id           int
	Language     string
	Stars        int
	Watchers     int
	FirstYear    int
	FirstQuarter int
}

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%4 == 0 {
		if year%100 == 0 {
			return false
		}
		return true
	}
	return false
}

func getData(year, month, day, hour int) (success bool, errs []error) {
	// Get the data from gharchive and save it to a file
	var errorArray []error
	if year > 2010 {
		errorArray = append(errorArray, errors.New("year must be at least 2011 and must be less than or equal to the current year"))
	}
	if month < 1 || month > 12 {
		errorArray = append(errorArray, errors.New("month must be between 1 - 12 inclusive"))
	}
	if day < 1 {
		errorArray = append(errorArray, errors.New("day must be greater than 1"))
	}
	if (month == 1 || month == 3 || month == 5 || month == 7 || month == 10 || month == 12) && day > 31 {
		errorArray = append(errorArray, errors.New("day must be less than or equal to 31 for Jan, Mar, May, July, Oct, and Dec"))
	} else if (month == 4 || month == 6 || month == 8 || month == 9 || month == 11) && day > 30 {
		errorArray = append(errorArray, errors.New("day must be less than or equal to 30 for Apr, Jun, Aug, Sept, and Nov"))
	} else if month == 2 {
		if isLeapYear(year) {
			if day > 29 {
				errorArray = append(errorArray, errors.New("day must be less than or equal to 29 for Feb during a leap year"))
			}
		} else {
			if day > 28 {
				errorArray = append(errorArray, errors.New("day must be less than or equal to 28 for Feb when it is not a leap year"))
			}
		}
	}

	if len(errorArray) > 0 {
		return false, errorArray
	}

	response, err := http.Get(fmt.Sprintf("https://data.gharchive.org/%d-%d-%d-%d.json.gz", year, month, day, hour))

	if err != nil {
		errorArray = append(errorArray, errors.New("An unknown error occurred while getting the requested data."))
		return false, errorArray
	}
	defer response.Body.Close()

	dataFilePath := "gharchive_data.json.gz"
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		errorArray = append(errorArray, errors.New("An error occurred while reading the http response"))
		return false, errorArray
	}
	writeErr := os.WriteFile(dataFilePath, responseBytes, os.FileMode(0644))
	if writeErr != nil {
		errorArray = append(errorArray, errors.New("An error occurred while writing the data to file."))
		return false, errorArray
	}

	return true, errorArray
}

func (r repoRecord) Stringify() (string, error) {
	result, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
