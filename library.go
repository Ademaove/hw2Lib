package hw2Library

import (
	_ "fmt"
	"github.com/jakehl/goid"
	"math"
)

func AlterString(s string) string{
	var res string
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			res += string(s[i] - 32)
		}else if s[i] >= 'A' && s[i] <= 'Z' {
			res += string(s[i] + 32)
		}else{
			res += string(s[i])
		}
	}
	return res
}

func GetRoots(a, b, c float64) (float64, float64) {
	dis := math.Sqrt(b*b - 4*a*c)
	denom := 2 * a
	x1 := (-b + dis) / denom
	x2 := (-b - dis) / denom
	return x1, x2
}

func GetUUID() *goid.UUID {
	return goid.NewV4UUID()
}



