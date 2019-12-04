package helpers

import (
	"math/rand"
	"time"
)

func Swap(left *int, right *int) {
	tmp := *left
	*left = *right
	*right = tmp
}

func Shuffle(arr []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(arr) > 0 {
		n := len(arr)
		randIndex := r.Intn(n)
		arr[n-1], arr[randIndex] = arr[randIndex], arr[n-1]
		arr = arr[:n-1]
	}
}

func SeedIntSlice(length int, max int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	arr := make([]int, length)
	for i := 0; i < len(arr); i++ {
		arr[i] = r.Intn(max)
	}
	return arr
}
