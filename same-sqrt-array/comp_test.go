package same_sqrt_array

import "testing"

func TestComp(t *testing.T) {
	type args struct {
		array1 []int
		array2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1. success",
			args: args{
				array1: []int{121, 144, 19, 161, 19, 144, 19, 11},
				array2: []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			want: true,
		},
		{
			name: "2. fail",
			args: args{
				array1: []int{121, 144, 19, 161, 19, 144, 19, 11},
				array2: []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			want: false,
		},
		{
			name: "3. fail",
			args: args{
				array1: nil,
				array2: []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Comp(tt.args.array1, tt.args.array2); got != tt.want {
				t.Errorf("Comp() = %v, want %v", got, tt.want)
			}
		})
	}
}
