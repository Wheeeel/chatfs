package main

import "strings"

func GetDirLevel(name string) int {
	tmp := strings.Split(name, "/")
	return len(tmp)
}

func GetLastLevelName(name string) string {
	tmp := strings.Split(name, "/")
	return tmp[len(tmp)-1]
}

func GetServerName(name string) string {
	tmp := strings.Split(name, "/")
	return tmp[0]
}
