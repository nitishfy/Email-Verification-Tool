package domain

import (
	"fmt"
	"github.com/nitishfy/Email-Verification-Tool/records/dmarcRecords"
	MxRecords "github.com/nitishfy/Email-Verification-Tool/records/mxRecords"
	"github.com/nitishfy/Email-Verification-Tool/records/spfRecords"
)

func CheckDomain(domain string) {
	MxRecords.CheckMxRecords(domain)
	spfRecords.CheckSpfRecords(domain)
	dmarcRecords.CheckDmarcRecords(domain)
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, MxRecords.HasMX, spfRecords.HasSPF, spfRecords.SPFRecord, dmarcRecords.HasDMARC, dmarcRecords.DMARCRecord)
}
