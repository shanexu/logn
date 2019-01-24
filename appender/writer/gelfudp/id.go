package gelfudp

import (
	"errors"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type IdGenerator interface {
	NextId() uint64
}

type DefaultIdGenerator struct {
	ip uint64
}

func NewDefaultIdGenerator(ip uint64) *DefaultIdGenerator {
	return &DefaultIdGenerator{ip}
}

func (g *DefaultIdGenerator) NextId() uint64 {
	timestamp := uint64(time.Now().Nanosecond())
	timestamp = timestamp << 32
	timestamp = timestamp >> 32
	return (g.ip << 32) | timestamp
}

func GuessIP() (uint64, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return 0, err
	}
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return 0, err
	}
	for _, ip := range ips {
		if ip.To4() != nil {
			parts := strings.Split(ip.String(), ".")
			a, _ := strconv.ParseUint(parts[0], 10, 64)
			b, _ := strconv.ParseUint(parts[1], 10, 64)
			c, _ := strconv.ParseUint(parts[2], 10, 64)
			d, _ := strconv.ParseUint(parts[3], 10, 64)
			return (a << 24) | (b << 16) | (c << 8) | d, nil
		}
	}
	return 0, errors.New("cannot resolve ip by hostname")
}
