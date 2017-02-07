package hamming

import "errors"

const testVersion = 5

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("error")
	}

	var count int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
