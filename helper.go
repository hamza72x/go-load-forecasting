package main

import (
	"log"
	"regexp"
)

func getNumbersOnly(str string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(str, "")
}
