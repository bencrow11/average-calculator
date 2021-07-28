package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := askForArray()
	min, max := getMinMax(numbers)
	median := getMedian(numbers, 2, 1)
	mean := getMean(numbers)
	SD := getSD(numbers)
	iqRange := getIQRange(numbers)
	totalRange := getRange(numbers)

	postResults("Min:", min)
	postResults("Max:", max)
	postResults("Median:", median)
	postResults("Mean", mean)
	postResults("Standard Deviation:", SD)
	postResults("Interquartile Range:", iqRange)
	postResults("Range:", totalRange)
}

func askForArray() []float64 {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your array with each number separated by a space")
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	textSlice := strings.Split(text, " ")
	intSlice := make([]float64, len(textSlice))

	for i, s := range textSlice {
		a, err := strconv.ParseFloat(s, 32)
		intSlice[i] = a
		if err != nil {
			panic(err)
		}
	}
	sort.Float64s(intSlice)
	return intSlice
}

func getMinMax(data []float64) (float64, float64) {
	min := data[0]
	max := data[len(data)-1]
	return min, max
}

func getMedian(data []float64, position float64, multiplyer float64) float64 {
	middleNum := float64((len(data) + 1)) / position * multiplyer
	if middleNum == float64(int64(middleNum)) {
		return data[int(middleNum)-1]
	} else {
		return (data[int(middleNum-1.5)] + data[int(middleNum-0.5)]) / 2
	}
}

func getMean(data []float64) float64 {
	var total float64
	for _, d := range data {
		total += d
	}
	return total / float64(len(data))
}

func getSD(data []float64) float64 {
	mean := getMean(data)

	squaredData := make([]float64, len(data))
	for i, d := range data {
		result := (d - mean) * (d - mean)
		squaredData[i] = result
	}
	squaredMean := getMean(squaredData)
	standardDeviation := math.Sqrt(squaredMean)
	return standardDeviation
}

func getIQRange(data []float64) float64 {
	lowQ := getMedian(data, 4, 1)
	highQ := getMedian(data, 4, 3)
	return highQ - lowQ
}

func getRange(data []float64) float64 {
	min, max := getMinMax(data)
	return max - min
}

func postResults(s string, d float64) {
	fmt.Println()
	fmt.Println(s, d)
}
