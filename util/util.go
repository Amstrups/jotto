package util

import "fmt"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func InBinary(n int32) string {
	return fmt.Sprintf("%032b", n)
}

func ToString(alpahbet string, n []int) string {
	b := []byte{}

	for _, v := range n {
		s := alphabet[v]
		b = append(b)
	}
	return string([]byte{n})
}
