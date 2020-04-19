package main

import "fmt"

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	totalStock := 0
	for baseIndex := 0; baseIndex < len(ar); baseIndex++ {
		for matchIndex := baseIndex + 1; matchIndex < len(ar); matchIndex++ {
			if ar[baseIndex] == ar[matchIndex] && ar[baseIndex] != 0 {
				totalStock++
				ar[baseIndex] = 0
				ar[matchIndex] = 0
			}
		}
	}
	return int32(totalStock)
}

func main() {
	sockColours := 9
	sockList := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	stocks := sockMerchant(int32(sockColours), sockList)
	fmt.Printf("stock sock = %d", stocks)
}
