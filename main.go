package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/montanaflynn/stats"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	data := []float64{}
	for s.Scan() {
		in := s.Text()
		v, err := strconv.ParseFloat(in, 64)
		if err != nil {
			io.WriteString(os.Stderr, "(ignoring): "+err.Error())
			continue
		}
		data = append(data, v)
	}

	mean, _ := stats.Mean(data)
	min, _ := stats.Min(data)
	max, _ := stats.Max(data)
	p50, _ := stats.Percentile(data, 50)
	p75, _ := stats.Percentile(data, 75)
	p90, _ := stats.Percentile(data, 90)
	p95, _ := stats.Percentile(data, 95)
	p99, _ := stats.Percentile(data, 99)

	fmt.Printf("total: %v\n", len(data))
	fmt.Printf("min: %v\n", min)
	fmt.Printf("mean: %v\n", mean)
	fmt.Printf("max: %v\n", max)
	fmt.Printf("p50: %v\n", p50)
	fmt.Printf("p75: %v\n", p75)
	fmt.Printf("p90: %v\n", p90)
	fmt.Printf("p95: %v\n", p95)
	fmt.Printf("p99: %v\n", p99)

}
