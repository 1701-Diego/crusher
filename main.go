package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

var portCheck = flag.Int(
	"port-check",
	0,
	"port to check",
)

var urlCheck = flag.String(
	"url-check",
	"",
	"url to check",
)

var timeout = flag.Duration(
	"timeout",
	1*time.Second,
	"dial timeout",
)

func main() {
	flag.Parse()

	if *portCheck <= 0 && *urlCheck == "" {
		fmt.Println("This mission's impossible: crusher needs either a port to check or a url to check.")
		os.Exit(1)
	}

	if *portCheck > 0 {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf(":%d", *portCheck), *timeout)
		if err != nil {
			fmt.Printf("portcheck failed: %s\n", err.Error())
			os.Exit(1)
		}
		conn.Close()
		fmt.Println("portcheck passed")
		os.Exit(0)
	}

	if *urlCheck != "" {
		client := &http.Client{
			Timeout: *timeout,
		}
		resp, err := client.Get(*urlCheck)
		if err != nil {
			fmt.Printf("urlcheck failed: %s\n", err.Error())
			os.Exit(1)
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("urlcheck failed: need 200, got %d\n", resp.StatusCode)
			os.Exit(1)
		}

		fmt.Println("urlcheck passed")
		os.Exit(0)
	}

	fmt.Println("What am I doing here?")
	os.Exit(1)
}
