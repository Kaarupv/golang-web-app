package main

import "log"

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "samson"

	log.Println(myMap["dog"])
}
