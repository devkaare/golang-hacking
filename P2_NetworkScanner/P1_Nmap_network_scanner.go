package main

import (
	"context"
	"fmt"
	"github.com/Ullaakut/nmap"
	"log"
	"time"
)

func main() {
	targetIP := "192.168.0.1/24"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	defer cancel()
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(targetIP),
		nmap.WithPorts("80, 443, 999"),
		nmap.WithContext(ctx),
	)

	if err != nil {
		log.Fatal("error :", err)
	}

	results, warning, err := scanner.Run()
	if err != nil {
		log.Fatal("error :", err)
	}

	if warning != nil {
		log.Fatalf("error : %s\n", warning)
	}

	for _, host := range results.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("IP: %q", host.Addresses[0])
		if len(host.Addresses) > 1 {
			fmt.Printf("MAC: %s\n", host.Addresses[1])
		}

		for _, port := range host.Ports {
			fmt.Printf("\t Port %d %s %s %s \n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}
}
