package same_sqrt_array

import "math"

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil || len(array1) != len(array2) {
		return false
	}
	for i := 0; i < len(array2); i++ {
		if i == 0 {
			if array1[0] != array2[0] {
				return false
			}
		} else {
			if math.Sqrt(float64(array2[i])) != float64(array1[i-1]) {
				return false
			}
		}
	}
	return true
}
