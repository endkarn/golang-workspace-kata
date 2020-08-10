package two_sum_test

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "submit#1",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "submit#2",
			args: args{
				nums:   []int{3, 2, 3},
				target: 6,
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func twoSum(nums []int, target int) []int {
	var indices []int
	rounds := 0
	for rounds < len(nums) {
		targetNumber := target
		nums = nums[rounds:]

		for index, num := range nums {
			if targetNumber-num > 0 {
				targetNumber = targetNumber - num
				indices = append(indices, rounds+index)

			} else if targetNumber-num == 0 {
				indices = append(indices, rounds+index)
				return indices
			}
		}
		rounds++
		indices = []int{}
	}
	return []int{}
}
