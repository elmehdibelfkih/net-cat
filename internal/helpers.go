package internal

import "fmt"

func MessageValid(str string) error {
	esc := '\x1b'
	for _, v := range str {
		if v == esc {
			return fmt.Errorf("you cannot enter control characters")
		}
	}
	return nil
}
