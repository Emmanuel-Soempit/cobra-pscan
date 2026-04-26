package scan

import (
	"fmt"
	"net"
	"time"
)

type PortState struct {
	Port int
	Open state
}

type state bool

func (s state) String() string {
	if s {
		return "open"
	}
	return "closed"
}

func scanPort(host string, port int) PortState {
	p := PortState{
		Port: port,
	}

	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	scanConn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		p.Open = false
	} else {
		p.Open = true
		scanConn.Close()
	}

	if err != nil {
		return p
	}

	scanConn.Close()
	p.Open = true

	return p
}

type Results struct {
	Host       string
	NotFound   bool
	PortStates []PortState
}

func Run(hl *HostsList, ports []int) []Results {
	res := make([]Results, 0, len(hl.Hosts))
	for _, host := range hl.Hosts {
		result := Results{
			Host: host,
		}
		if _, err := net.LookupHost(host); err != nil {
			result.NotFound = true
			res = append(res, result)
			continue
		}

		for _, port := range ports {
			result.PortStates = append(result.PortStates, scanPort(host, port))
		}
		res = append(res, result)
	}
	return res
}
