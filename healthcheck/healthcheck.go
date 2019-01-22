package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var flgIgnoreTLS = flag.Bool("ignore-tls", false, "ignore the issuer of the https certficate (allows health check on self signed cert)")

func main() {
	flag.Parse()

	var client = &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: *flgIgnoreTLS,
			},
		},
	}

	for i, arg := range os.Args {
		if i == 0 || arg == "-ignore-tls" {
			continue
		}

		resp, err := client.Get(arg)
		if err != nil {
			fmt.Println("[ERROR]", err.Error())
			os.Exit(1) // Unhealthy
		}
		if resp.StatusCode == 200 {
			continue
		} else {
			fmt.Println("[FAIL] url=" + arg + " statuscode=" + strconv.Itoa(resp.StatusCode))
			os.Exit(1) // Unhealthy
		}
	}

	os.Exit(0) // Healthy
}
