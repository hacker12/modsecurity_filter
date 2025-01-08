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
	ipPrefix := flag.String("ip_prefix", "", "Optional IP prefix to filter")
	showMsg := flag.Bool("show_msg", false, "Set to true to show the msg field")
	showUri := flag.Bool("show_uri", false, "Set to true to show the uri field")
	version := flag.Bool("version", false, "Show version")

	flag.Parse()

	if *version {
		fmt.Printf("modsecurity_filter version %s\n", Version)
		return
	}

	file, err := os.Open(*logfile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Adjusted regex patterns for log format
	var ipRegex *regexp.Regexp
	if *ipPrefix != "" {
		ipRegex = regexp.MustCompile(fmt.Sprintf(`\[client (%s[0-9.]+):[0-9]+\]`, *ipPrefix))
	} else {
		ipRegex = regexp.MustCompile(`\[client ([0-9.]+):[0-9]+\]`)
	}

	idRegex := regexp.MustCompile(`\[id "([0-9]+)"\]`)
	msgRegex := regexp.MustCompile(`\[msg "([^"]+)"\]`)
	uriRegex := regexp.MustCompile(`\[uri "([^"]+)"\]`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Match IP and ID
		ipMatches := ipRegex.FindStringSubmatch(line)
		idMatches := idRegex.FindStringSubmatch(line)

		// Match optional fields MSG and URI
		msgMatches := msgRegex.FindStringSubmatch(line)
		uriMatches := uriRegex.FindStringSubmatch(line)

		if len(ipMatches) > 1 && len(idMatches) > 1 {
			// Build the output based on optional flags
			output := fmt.Sprintf("IP: %s, ID: %s", ipMatches[1], idMatches[1])

			if *showMsg && len(msgMatches) > 1 {
				output += fmt.Sprintf(", MSG: %s", msgMatches[1])
			}
			if *showUri && len(uriMatches) > 1 {
				output += fmt.Sprintf(", URI: %s", uriMatches[1])
			}

			fmt.Println(output)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
