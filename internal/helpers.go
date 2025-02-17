package internal

import "fmt"

func MessageValid(str []byte) error {
	for _, v := range str {
		if v <= 31 && v != '\n' && v != '\r' {
			return fmt.Errorf("you cannot enter control characters")
		}
	}
	return nil
}
