package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	version := "0.0.1"
	upstream := "https://www.github.com/nat-418/knock"
	supported_networks := []string{
		"tcp", "tcp4", "tcp6",
		"udp", "upd4", "udp6",
		"ip", "ip4", "ip6",
		"unix", "unixgram", "unixpacket",
	}

	var dest string
	var port string
	var network string
	var timeout int

	flag.StringVar(&dest, "dest", "", "Network destination address or name")
	flag.StringVar(&network, "net", "tcp", "Network type to use, see NETWORKS")
	flag.StringVar(&port, "port", "80", "Port number to knock on destination")
	flag.IntVar(&timeout, "time", 15, "Time to wait in seconds before giving up")
	flag.Usage = func() {
		fmt.Println(
			"knock v" + version + "\n" +
				"A simple network reachability tester\n\n" +
				"USAGE:\n" +
				"  knock [OPTIONS] destination port\n\n" +
				"OPTIONS:",
		)

		flag.PrintDefaults()

		fmt.Println(
			"\nNETWORKS:\n  " +
				"Supported network types are tcp, tcp4 (IPv4-only), " +
				"tcp6 (IPv6-only),\n  udp, udp4 (IPv4-only), " +
				"udp6 (IPv6-only), ip, ip4 (IPv4-only), " +
				"ip6 (IPv6-only),\n  unix, unixgram, and unixpacket.\n" +
				"\nTo report bugs or view source code, see " +
				upstream +
				".",
		)

	}

	flag.Parse()

	if !slices.Contains(supported_networks, network) {
		fmt.Fprintln(os.Stderr, "Error: unsupported network.")
		os.Exit(1)
	}

	rest := flag.Args()

	if len(rest) == 1 {
		if strings.Contains(rest[0], ":") {
			parts := strings.Split(rest[0], ":")
			dest = parts[0]
			port = parts[1]
		} else {
			dest = rest[0]
		}
	}

	if len(rest) == 2 {
		dest = rest[0]
		port = rest[1]
	}

	if len(rest) > 2 {
		fmt.Fprintln(os.Stderr, "Error: too many arguments.")
		os.Exit(1)
	}

	if strings.Contains(dest, ":") {
		fmt.Fprintln(os.Stderr, "Error: invalid destination.")
		os.Exit(1)
	}

	if dest == "" {
		fmt.Fprintln(os.Stderr, "Error: no destination specified.")
		os.Exit(1)
	}

	target := dest + ":" + port

	fmt.Fprintln(os.Stdout, "Trying to knock on "+target+"…")

	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		fmt.Fprintln(os.Stdout, "Failed: connection timed out after", timeout, "seconds.")
		os.Exit(1)
	}()

	_, err := net.Dial(network, target)

	if err != nil {
		fmt.Fprintln(os.Stdout, "Failed with error:\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintln(os.Stdout, "Succeeded.")
		os.Exit(0)
	}

}
