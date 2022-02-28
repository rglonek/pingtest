package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/go-ping/ping"
)

var printf = log.Printf

func main() {
	dst := flag.String("host", "", "hostname to ping")
	cnt := flag.Int("number", 1, "number of pings to run")
	interval := flag.Duration("interval", time.Second, "how long to wait between ping attempts; also defines each ping and total timeout")
	priv := flag.Bool("privileged", false, "run in privileged mode")
	quiet := flag.Bool("quiet", false, "do not output any logging except for fatal errors and final result")
	rtts := flag.Bool("rtts", false, "set to print each individual packet RTT, not just stats; setting quiet disables this too")
	threshold := flag.Int("threshold", 100, "packet loss percentage threshold, when test is considered to be failed; will return with errno 127")
	flag.Parse()
	if *dst == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *quiet {
		printf = func(format string, v ...interface{}) {
		}
	}
	printf("Creating pinger")
	pinger, err := ping.NewPinger(*dst)
	if err != nil {
		log.Fatalf("Failed to create pinger: %s", err)
	}

	printf("Configuring")
	pinger.SetPrivileged(*priv)
	pinger.Count = *cnt
	pinger.Timeout = *interval * time.Duration(*cnt)
	pinger.Interval = *interval

	printf("Ping Start")
	err = pinger.Run()
	if err != nil {
		log.Fatalf("Failed to run pinger: %s", err)
	}
	printf("Ping End")

	stats := pinger.Statistics()
	printf("Addr:%s IP:%s Received:%d Sent:%d Duplicates:%d RTT(min/max/avg):%s/%s/%s Loss:%0.2f", stats.Addr, stats.IPAddr.String(), stats.PacketsRecv, stats.PacketsSent, stats.PacketsRecvDuplicates, stats.MinRtt, stats.MaxRtt, stats.AvgRtt, stats.PacketLoss)
	if *rtts {
		printf("RTTs: %v", stats.Rtts)
	}
	ploss := math.Ceil(stats.PacketLoss)
	if *quiet {
		fmt.Printf("%0.0f", ploss)
	}
	if int(ploss) >= *threshold {
		os.Exit(127)
	}
}
