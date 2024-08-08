package main

import (
	"fmt"
	"log"
	"os"
)

const (
	dataFileName = "data.csv"
)

func main() {
	input := os.Args[1:]
	if len(input) == 0 {
		noInputProvided()
		os.Exit(0)
	}

	IPs := parseIPs(input)
	if len(IPs) == 0 {
		noInputProvided()
		os.Exit(0)
	}
	log.Printf("Received %d IP addresses to check\n", len(IPs))

	data, err := readCSVFile(dataFileName, ';')
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to read data file: %v", err))
	}

	hosters, err := parseHostersData(data)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to parse data into hosters array: %v", err))
	}
	if len(hosters) == 0 {
		log.Fatal("Hosters data not loaded")
	}
	log.Printf("Loaded info about %d hosters\n", len(hosters))

	runChecker(IPs, hosters)
	log.Println("Done")
}

func noInputProvided() {
	log.Println("No IP address provided. Please provide at least one IP address.")
	log.Println("Ex.:")
	log.Printf("\t%s 193.106.92.161 193.106.92.177\n", os.Args[0])
}
