package main

import (
	"log"
	"net"
	"net/netip"
)

func runChecker(IPs []net.IP, hosters []Hoster) {
	for _, ip := range IPs {
		if found, hoster := checkIP(ip, hosters); found {
			log.Printf("Found IP %s at hoster '%v'\n", ip, hoster)
		}
	}
}
func checkIP(ip net.IP, hosters []Hoster) (bool, Hoster) {
	for _, hoster := range hosters {
		if isHosterHasIp(hoster, ip) {
			return true, hoster
		}
	}
	return false, Hoster{}
}

func isHosterHasIp(h Hoster, ip net.IP) bool {

	if len(h.HandleStart) > 0 && len(h.HandleEnd) > 0 {
		addressCheck, err := netip.ParseAddr(ip.String())
		if err != nil {
			return false
		}
		addressStart, err := netip.ParseAddr(h.HandleStart.String())
		if err != nil {
			return false
		}
		addressEnd, err := netip.ParseAddr(h.HandleEnd.String())
		if err != nil {
			return false
		}
		return addressCheck.Compare(addressStart)+addressEnd.Compare(addressCheck) > 0
	}
	return h.HandleCIDR.Contains(ip)
}
