package MxRecords

import (
	"log"
	"net"
)

var HasMX bool

// CheckMxRecords checks for MX records
func CheckMxRecords(domain string) {
	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecord) > 0 {
		HasMX = true
	}
}
