package firewall

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IpAddress struct {
	sour string
	dest string
}
type traffic struct {
	sys int64
	pas int64
}

// take the file again, go through each line and get the size, and source destination ip
func Inspect(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(file)
	traffic := traffic{}
	for scanner.Scan() {
		hex := scanner.Text()
		curTraffic, err := strconv.ParseInt(hex[4:8], 16, 64)
		if err != nil {
			return "", err
		}
		ip, err := NewIp(hex[24:])
		if err != nil {
			return "", err
		}
		const sys = "192.168"
		const pas = "10.0"

		if strings.HasPrefix(ip.sour, sys) || strings.HasPrefix(ip.dest, sys) {
			traffic.sys += curTraffic
			continue
		}
		traffic.pas += curTraffic

	}

	return fmt.Sprintf("%d/%d", traffic.sys, traffic.pas), nil
}

// convert a hexadecimal ip addresses and convert them to base 64
func NewIp(hex string) (*IpAddress, error) {
	if len(hex) > 16 || len(hex) == 0 {
		return nil, errors.New("firewall.NewIp: invalid input provided")
	}
	const half = 8
	out := IpAddress{}
	for idx, hexIP := range []string{hex[:half], hex[half:]} {
		var ip strings.Builder
		for i := 0; i < half; i = i + 2 {
			num, err := strconv.ParseInt(hexIP[i:i+2], 16, 64)
			if err != nil {
				return nil, err
			}

			ip.WriteString(strconv.Itoa(int(num)))
			if i < 6 {
				ip.WriteByte('.')
			}
		}
		if idx == 0 {
			out.sour = ip.String()
			continue
		}
		out.dest = ip.String()
	}
	return &out, nil
}
