package main

import (
	"fmt"
	"io/ioutil"
	"os"
)


func loadData() []byte {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	
	return data
}


func main() {
	fmt.Println(string(loadData()))
}
