package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type PhoneNumber struct {
	ID      int    `csv:"id"`
	Number  string `csv:"number"`
	Country string `csv:"country"`
}

func readCSV(filename string) ([]PhoneNumber, error) {
	var phoneNumbers []PhoneNumber
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := gocsv.Unmarshal(file, &phoneNumbers); err != nil {
		return nil, err
	}
	return phoneNumbers, nil
}

func main() {

	fileName := flag.String("csv", "phonenumbers.csv", "Filepath of the csv file containing phone numbers")
	flag.Parse()

	pnums, err := readCSV(*fileName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pnums)

}
