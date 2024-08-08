package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

func parseIPs(inputIPs []string) []net.IP {
	result := []net.IP{}
	for _, ip := range inputIPs {
		parsedIP := net.ParseIP(strings.TrimSpace(ip))
		if parsedIP == nil {
			log.Printf("failed to parse IP %s", ip)
			continue
		}
		result = append(result, parsedIP)
	}
	return result
}

func parseCIDR(inputCIDR string) (net.IPNet, error) {
	result := net.IPNet{}

	_, network, err := net.ParseCIDR(strings.TrimSpace(inputCIDR))
	if err != nil {
		return result, errors.New(fmt.Sprintf("failed to parse CIDR %s: %v", network, err))
	}
	result = *network

	return result, nil
}
