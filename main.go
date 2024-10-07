package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	linearstats "linearstats/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("The command to run the programm will be: go  run . main.go data.txt")
		return
	}
	if os.Args[1] != "data.txt" {
		fmt.Println("Error: the file containing data is data.txt not", os.Args[1])
		return
	}
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []float64
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			continue
		}
		data = append(data, number)
	}
	size := float64(len(data))
	Sx, Sy, Sxy, P_Sx, P_Sy := linearstats.CalcSums(data)

	slope := linearstats.CalcSlope(data, size, Sx, Sy, Sxy, P_Sx)
	b := (Sy - (slope * Sx)) / size
	PCC := linearstats.PearsonCorrelationCoefficient(size, Sx, Sy, Sxy, P_Sx, P_Sy)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", PCC)
}
