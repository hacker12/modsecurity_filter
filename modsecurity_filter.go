package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

// Version will be set during compilation
var Version string

func main() {
	logfile := flag.String("logfile", "confluence_error.log", "Log file to parse")
	ipPrefix := flag.String("ip_prefix", "", "IP prefix to filter")
	showMsg := flag.Bool("show_msg", false, "Set to true to show the msg field")
	showUri := flag.Bool("show_uri", false, "Set to true to show the uri field")
	version := flag.Bool("version", false, "Show version")

	flag.Parse()

	if *version {
		fmt.Printf("modsecurity_filter version %s\n", Version)
		return
	}

	if *ipPrefix == "" {
		fmt.Println("Error: ip_prefix is required")
		flag.Usage()
		return
	}

	file, err := os.Open(*logfile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	ipRegex := regexp.MustCompile(fmt.Sprintf(`\[client (%s[0-9.]+)\]`, *ipPrefix))
	idRegex := regexp.MustCompile(`\[id "([0-9]+)"\]`)
	msgRegex := regexp.MustCompile(`\[msg "([^"]+)"\]`)
	uriRegex := regexp.MustCompile(`\[uri "([^"]+)"\]`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ipMatches := ipRegex.FindStringSubmatch(line)
		idMatches := idRegex.FindStringSubmatch(line)
		msgMatches := msgRegex.FindStringSubmatch(line)
		uriMatches := uriRegex.FindStringSubmatch(line)

		if len(ipMatches) > 1 && len(idMatches) > 1 {
			if *showMsg && *showUri && len(msgMatches) > 1 && len(uriMatches) > 1 {
				fmt.Printf("IP: %s, ID: %s, MSG: %s, URI: %s\n", ipMatches[1], idMatches[1], msgMatches[1], uriMatches[1])
			} else if *showMsg && len(msgMatches) > 1 {
				fmt.Printf("IP: %s, ID: %s, MSG: %s\n", ipMatches[1], idMatches[1], msgMatches[1])
			} else if *showUri && len(uriMatches) > 1 {
				fmt.Printf("IP: %s, ID: %s, URI: %s\n", ipMatches[1], idMatches[1], uriMatches[1])
			} else {
				fmt.Printf("IP: %s, ID: %s\n", ipMatches[1], idMatches[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
