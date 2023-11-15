package service

import (
	"fmt"
)

func CheckGrade(i int) string {
	fmt.Println("")
	
	switch {
	case i >= 80:
		return "A"
	case i >= 70:
		return "B"
	case i >= 60:
		return "C"
	case i >= 50:
		return "D"
	default:
		return "F"
	}

}
