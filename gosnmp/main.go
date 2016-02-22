package main

import (
	"fmt"
	"github.com/alouca/gosnmp"
)

func main() {
	const cmdTimeout = 10
	const cmdOid = ".1.3.6.1.4.1.12356.101.4.1.1" // Device Version
	//const cmdOid = ".1.3.6.1.4.1.12356.101.4.1.4.0" // MemUsage

	s, err := gosnmp.NewGoSNMP("10.168.1.1", "public", gosnmp.Version2c, cmdTimeout)
	if err != nil {
		fmt.Printf("Error creating SNMP instance: %s\n", err.Error())
		return
	}
	//s.SetDebug(true)
	//s.SetVerbose(true)

	s.SetTimeout(cmdTimeout)
	oid := cmdOid
	for i := 0; i < 10; i++ {
		fmt.Printf("Getting %s\n", oid)
		resp, err := s.GetNext(oid)
		if err != nil {
			fmt.Printf("Error getting response: %s\n", err.Error())
		} else {
			for _, v := range resp.Variables {
				fmt.Printf("%s -> ", v.Name)
				switch v.Type {
				case gosnmp.OctetString:
					if s, ok := v.Value.(string); ok {
						fmt.Printf("%s\n", s)
					} else {
						fmt.Printf("Response is not a string\n")
					}
				default:
					fmt.Printf("Type: %s(%#x) - Value: %v\n", v.Type, int(v.Type), v.Value)
				}
				oid = v.Name
			}
		}
	}
}
