package main

import (
	"DDOS_ARMY/client"
	"DDOS_ARMY/server"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// -server
	sf := flag.Bool("server", false, "run as server")
	// -client
	cf := flag.Bool("client", false, "run as client")

	flag.Parse()

	if *sf {
		server.StartServer("127.0.0.1", "8080")
	}
	if *cf {
		c := client.GetClient()

		//handle ctrl+c and control+d signal
		exitChan := make(chan bool)

		// Handle signals in a separate goroutine
		go func(cl *client.Client) {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			<-c // Wait for a signal to be received

			// Go out from camp
			cl.LeaveCamp()
			exitChan <- true // Signal that the program should exit
		}(c)

		for {
			select {
			case <-exitChan:
				return
			default:
			}

			c.JoinCamp()

			res, _ := c.LeaveCamp()
			fmt.Print(res)

			var input string
			fmt.Scanln(&input)
		}
	}
}
