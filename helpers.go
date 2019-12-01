package helpers

func Swap(left *int, right *int) {
	tmp := *left
	*left = *right
	*right = tmp
}

