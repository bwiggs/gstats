package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"text/tabwriter"

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

	w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', tabwriter.TabIndent)
	// fmt.Fprintln(w, "-----\t-----")
	fmt.Fprintf(w, "total\t%d\n", len(data))
	fmt.Fprintf(w, "min\t%.0f\n", min)
	fmt.Fprintf(w, "mean\t%.0f\n", mean)
	fmt.Fprintf(w, "max\t%.0f\n", max)
	fmt.Fprintf(w, "p50\t%.0f\n", p50)
	fmt.Fprintf(w, "p75\t%.0f\n", p75)
	fmt.Fprintf(w, "p90\t%.0f\n", p90)
	fmt.Fprintf(w, "p95\t%.0f\n", p95)
	fmt.Fprintf(w, "p99\t%.0f\n", p99)
	w.Flush()

}
