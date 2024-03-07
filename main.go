package main

import (
	"flag"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"net"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	version := "0.0.2"
	upstream := "https://www.github.com/nat-418/knock"
	ok_nets := []string{
		"tcp", "tcp4", "tcp6",
		"udp", "upd4", "udp6",
		"ip", "ip4", "ip6",
		"unix", "unixgram", "unixpacket",
	}

	target, dest, port, timeout, network := parseCli(version, upstream, ok_nets)

	sp := startSpinner(dest, port)

	startTimer(timeout, sp)

	knock(target, network, timeout, sp)
}

func knock(target string, network string, timeout int, sp *spinner.Spinner) {
	_, err := net.Dial(network, target)

	if err != nil {
		sp.Stop()
		if strings.Contains(err.Error(), "connection refused") {
			failMsg("connection refused")
		} else {
			failMsg(err.Error())
		}
		os.Exit(1)
	} else {
		sp.Stop()
		color.New(color.Bold, color.FgGreen).Fprintln(os.Stdout, "Succeeded.")
		os.Exit(0)
	}
}

func parseCli(version string, upstream string, ok_nets []string) (string, string, string, int, string) {
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
			"\nNETWORKS:\n" +
				"  Supported network types are tcp, tcp4, tcp6, udp, udp4, udp6,\n" +
				"  ip, ip4, ip6, unix, unixgram, and unixpacket. Network types\n" +
				"  with 4 in the name are IPv4-only, and those with 6 are likewise\n" +
				"  IPv6-only.\n\n" +
				"To report bugs or view source code, see " +
				upstream +
				".",
		)

	}

	flag.Parse()

	if !slices.Contains(ok_nets, network) {
		abortMsg("unsupported network")
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
		abortMsg("too many arguments")
	}

	if strings.Contains(dest, ":") {
		abortMsg("invalid destination")
	}

	if dest == "" {
		abortMsg("no destination specified")
	}

	target := dest + ":" + port

	return target, dest, port, timeout, network
}

func startSpinner(dest string, port string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)

	s.Suffix =
		" Knocking on " +
			color.New(color.Bold, color.FgBlue).Sprint(dest) +
			":" +
			color.New(color.Bold, color.FgMagenta).Sprint(port) +
			"…"
	s.Color("green", "bold")
	s.Start()

	return s
}

func startTimer(timeout int, sp *spinner.Spinner) {
	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		sp.Stop()
		t := strconv.Itoa(timeout)
		failMsg("connection timed out after " + t + " seconds")
		os.Exit(1)
	}()
}

func abortMsg(s string) {
	color.New(color.Bold, color.FgRed).Fprint(os.Stderr, "Error:")
	fmt.Fprintln(os.Stderr, " "+s+".")
	os.Exit(1)
}

func failMsg(s string) {
	color.New(color.Bold, color.FgRed).Fprint(os.Stdout, "Failed:")
	fmt.Fprintln(os.Stdout, " "+s+".")
}
