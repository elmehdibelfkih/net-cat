package internal

import "fmt"

func MessageValid(str string) error {
	for _, v := range str {
		if v <= '\x7f' && v >= '\x00' {
			return fmt.Errorf("you cannot enter control characters")
		}
	}
	return nil
}
