package main

import (
	"DDOS_ARMY/antena"
	"DDOS_ARMY/doss"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	isServer := flag.Bool("server", false, "run as server")
	isClient := flag.Bool("client", false, "run as client")
	port := flag.Int("port", 8080, "port number")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "OPTIONS:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *isServer {
		fmt.Printf("Running as server on port %d\n", *port)
		err := antena.StartServer(*port)
		if err != nil {
			return
		}
	} else if *isClient {
		fmt.Println("Running as client")
		//add test client
		client := doss.NewDefaultClient()
		response := client.Ping("http://127.0.0.1", 8080)
		log.Print(response)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
