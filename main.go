package main

import (
	"DDOS_ARMY/client"
	"DDOS_ARMY/server"
	"flag"
	"fmt"
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
		for {
			fmt.Println(c)
			//ping server
			res, _ := c.Ping()
			fmt.Println(res)
			//get camp info
			res, _ = c.GetCampInfo()
			fmt.Println(res)

			//join camp
			res, err := c.JoinCamp()
			if err != nil {
				panic(err)
			}
			fmt.Println(res)
			res, _ = c.GetCampInfo()
			fmt.Println(res)

			fmt.Println(c.ReceiveOrder())
			var input string
			fmt.Scanln(&input)
		}
	}
}
