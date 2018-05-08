package main

import "strings"

func GetDirLevel(name string) int {
	tmp := strings.Split(name, "/")
	return len(tmp)
}
