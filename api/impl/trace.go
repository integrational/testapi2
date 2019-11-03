package impl

import "log"

func trace(s string) string {
	log.Println("=>", s)
	return s
}

func un(s string) {
	log.Println("<=", s)
}
