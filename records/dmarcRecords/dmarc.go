package dmarcRecords

import (
	"log"
	"net"
	"strings"
)

var DMARCRecord string
var HasDMARC bool

// CheckDmarcRecords checks for the DMARC Records
func CheckDmarcRecords(domain string) {
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARCi") {
			HasDMARC = true
			DMARCRecord = record
			break
		}
	}
}
