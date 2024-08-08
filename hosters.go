package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type Hoster struct {
	Org         string
	Inn         int64
	HandleStart net.IP
	HandleEnd   net.IP
	HandleCIDR  net.IPNet
	Owner       string
	Address     string
	Location    string
}

func parseHostersData(data [][]string) ([]Hoster, error) {
	if len(data) < 1 {
		return nil, fmt.Errorf("no hosters data provided")
	}
	hosters := []Hoster{}
	// Обрабатываем каждую строку из CSV файла
	for i, record := range data {
		hoster, err := createHoster(record)
		if err != nil {
			log.Printf("Error parsing data for hoster at row %d: %v\n", i, err)
			continue
		}
		hosters = append(hosters, hoster)
	}

	return hosters, nil
}

func createHoster(record []string) (Hoster, error) {

	result := Hoster{}

	if len(record) < 3 {
		return result, errors.New(fmt.Sprintf("no much fields: %v\n", record))
	}

	// Парсим название хостера
	result.Org = strings.TrimSpace(record[0])

	// Парсим ИНН хостера
	inn, err := strconv.ParseInt(record[1], 10, 64)
	if err != nil {
		log.Printf("unable to parse hoster '%s' inn as INT64: %v", result.Org, err)
	} else {
		result.Inn = inn
	}

	// Парсим диапазон IP адресов хостера
	ipRange := strings.Split(record[2], "-")
	switch len(ipRange) {
	case 1:
		cidr, err := parseCIDR(ipRange[0])
		if err != nil {
			return result, errors.New(fmt.Sprintf("unable to parse hoster '%s' IP CIDR: '%s'", result.Org, record[2]))
		}
		result.HandleCIDR = cidr
	case 2:
		ips := parseIPs(ipRange)
		if len(ips) != 2 {
			return result, errors.New(fmt.Sprintf("unable to parse hoster '%s' IP range: '%s'", result.Org, record[2]))
		}
		result.HandleStart = ips[0]
		result.HandleEnd = ips[1]
	default:
		return result, errors.New(fmt.Sprintf("does not know what to do with hoster '%s' handle: '%s'", result.Org, record[2]))
	}

	if len(record) > 3 {
		result.Owner = strings.TrimSpace(record[3])
	}

	if len(record) > 4 {
		result.Address = strings.TrimSpace(record[4])
	}

	if len(record) > 5 {
		result.Location = strings.TrimSpace(record[5])
	}

	return result, nil
}
