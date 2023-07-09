package main

import (
	"bufio"
	"fmt"
	"github.com/nitishfy/Email-Verification-Tool/domain"
	"log"
	"os"
)

func main() {
	// create a scanner object with standard input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	// each time the input is provided from terminal, call the checkDomain function
	for scanner.Scan() {
		// pass the text you've input to the checkDomain function
		domain.CheckDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read fromthe input: %v", err)
	}

}
