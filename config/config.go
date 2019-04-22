package config

import "flag"

var (
	APort  int
	BPort  int
	IsMock bool
)

func init() {
	a := flag.Int("a", 8080, "admin port")
	b := flag.Int("b", 9090, "business port")
	m := flag.Bool("m", false, "is mock ?")
	flag.Parse()
	APort = *a
	BPort = *b
	IsMock = *m
}
