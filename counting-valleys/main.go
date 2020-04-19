package main

import (
	"fmt"
	"strings"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	totalHike := 0
	currentLevel := 0
	steps := strings.Split(s, "")
	for i := 0; i < len(steps); i++ {
		if steps[i] == "U" {
			currentLevel++

		}
		if steps[i] == "D" {
			currentLevel--

		}
		if steps[i] == "U" && currentLevel == 0 {
			totalHike++
		}
	}
	return int32(totalHike)
}

func main() {
	type args struct {
		number int32
		string string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "1",
			args: args{
				number: 8,
				string: "UDDDUDUU",
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				number: 12,
				string: "DDUUDDUDUUUD",
			},
			want: 2,
		},
	}

	for _, testCase := range tests {

		if got := countingValleys(testCase.args.number, testCase.args.string); got != testCase.want {
			fmt.Printf("Test Failed , Actual is %d (expected = %d)", got, testCase.want)
		} else {
			fmt.Printf("Run Test : %s (PASS)\n", testCase.name)
		}
	}

}
