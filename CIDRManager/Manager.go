package CIDRManager

import (
	"encoding/binary"
	"errors"
	"math/rand"
	"net"
	"sync"
)

var (
	EOCIDR = errors.New("EOCIDR (End of CIDR)")
)

type CIDRManager struct {
	CIDR    string
	Ipv4Min uint32
	Ipv4Max uint32
	Size    int
	Filter  *BSet
	Mutex   sync.Mutex
}

func NewCIDR(CIDR string) *CIDRManager {
	_, ipNet, _ := net.ParseCIDR(CIDR)
	Size := CountIPsInCIDR(ipNet)

	IPv4Min := binary.BigEndian.Uint32(ipNet.IP.To4())
	IPv4Max := IPv4Min | ^binary.BigEndian.Uint32(net.IP(ipNet.Mask).To4())

	return &CIDRManager{
		CIDR:    CIDR,
		Ipv4Min: IPv4Min,
		Ipv4Max: IPv4Max,
		Size:    Size,
		Filter:  NewBSet(),
		Mutex:   sync.Mutex{},
	}
}

func (c *CIDRManager) IsUsed(ip uint32) bool {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	return c.Filter.Contains(c.Uint32ToIP(ip))
}

func (c *CIDRManager) SetUsed(ip uint32) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Filter.Add(c.Uint32ToIP(ip))
}

func (c *CIDRManager) GetRandomIP() (string, error) {
	if c.Filter.Count() == c.Size {
		return "", EOCIDR
	}

	for {
		// Generate a random number within the range of the CIDR size
		randomIndex := rand.Intn(c.Size)
		ip := c.Ipv4Min + uint32(randomIndex)

		ipParsed := c.Uint32ToIP(ip)
		if !c.IsUsed(ip) {
			c.SetUsed(ip)
			return ipParsed, nil
		}
	}
}

func (c *CIDRManager) Uint32ToIP(ip uint32) string {
	return net.IPv4(byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip)).String()
}

func (c *CIDRManager) IPToUInt32(ip string) uint32 {
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ipAddr.To4())
}

func CountIPsInCIDR(ipNet *net.IPNet) int {
	maskSize, _ := ipNet.Mask.Size()
	if maskSize == 0 {
		return 1 << 32 // 2^32 = 4,294,967,296
	}

	return 1 << (32 - maskSize)
}
