package Modify

import "log"

func Argcheckers(s string) error {
	for _, w := range s {
		if w > 127 {
			log.Fatal("Only Printable Characters")
		}
	}
	return nil
}
