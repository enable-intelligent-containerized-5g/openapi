package models

import (
	"encoding/json"
	"math"
	"os"
	"strings"
	"time"
)

// Divide a value in a PrometheusResult other float64
func DivideValues(results *[]PrometheusResult, divisor float64) {
	if math.IsNaN(divisor) || divisor == 0 {
		divisor = 1
	}
	for i := range *results {
		(*results)[i].Value /= divisor
	}
}

// Save a struct in a JSON file
func SaveToJson(filename string, data interface{}) error {
	// Crear el archivo
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert the data
	indentedData, err := json.MarshalIndent(data, "", "    ") 
	if err != nil {
		return err
	}

	// Write the data
	_, err = file.Write(indentedData)
	return err
}


// Load CSV files
func LoadCsvFiles(dirPath string) (files []string, err error) {
	filesDir, err := os.ReadDir(dirPath)
	if err != nil {
		return files, err
	}

	for _, file := range filesDir {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".csv") {
			files = append(files, file.Name())
		}
	}

	return files, nil
}

// Get percentil
func GetPercentil(value float64, limit float64) float64 {
	load := value / limit
	if math.IsNaN(value) || limit == 0 {
		load = 0
	}
	return load
}

// SubtractSeconds subtracts seconds from a given date
func SubtractSeconds(date time.Time, seconds int64) time.Time {
	return date.Add(-time.Duration(seconds) * time.Second)
}

func ParseTimeToSeconds(startTime *time.Time, endTime *time.Time) int64 {
	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	return endTimeUnix - startTimeUnix
}

// Struct for CSV data
type CsvData struct {
	Pod        string
	Container  string
	Timestamp1 int64
	CpuUsage1  float64
	Timestamp2 int64
	CpuUsage2  float64
	Timestamp3 int64
	CpuUsage3  float64
	Timestamp4 int64
	CpuUsage4  float64
}