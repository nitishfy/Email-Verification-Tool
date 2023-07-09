package spfRecords

import (
	"log"
	"net"
	"strings"
)

var HasSPF bool
var SPFRecord string

// CheckSpfRecords checks for the SPF Records
func CheckSpfRecords(domain string) {
	spfRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range spfRecords {
		if strings.HasPrefix(record, "v=spfi") {
			HasSPF = true
			SPFRecord = record
			break
		}
	}
}
