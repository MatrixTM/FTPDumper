package Utility

import (
	"net"
	"os"
	"strings"
)

func IsInPipeline() bool {
	fi, _ := os.Stdin.Stat()
	return (fi.Mode() & os.ModeCharDevice) == 0
}

func IsCIDRv4(s string) bool {
	c, _, err := net.ParseCIDR(s)
	return err == nil && c.To4() != nil
}

func IsIPv4(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil && ip.To4() != nil
}

func SuffixAny(s string, suffixes []string) bool {
	for _, e := range suffixes {
		if strings.HasSuffix(s, e) {
			return true
		}
	}
	return false
}
