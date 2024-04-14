package Core

import "C"
import (
	"FTPDumper/CIDRManager"
	"FTPDumper/Utility"
	"bufio"
	"errors"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"sync"
)

type EscannerType string

var (
	ESTDIN EscannerType = "stdin"
	EFILE  EscannerType = "file"
	ECIDR  EscannerType = "cidr"
	EIP    EscannerType = "ip"
)

type IReader interface {
	Next() (string, error)
	Close()
}

type StdinReader struct {
	scanner *bufio.Scanner
}

type CIDRReader struct {
	sync.Mutex
	Cidrs    []*CIDRManager.CIDRManager
	CidrsLen int
}

func NewReader(scanner string, method EscannerType) IReader {
	switch method {
	case ESTDIN:
		return &StdinReader{scanner: bufio.NewScanner(os.Stdin)}
	case EFILE:
		file, _ := os.Open(scanner)
		return NewFileReader(bufio.NewReader(file))
	case ECIDR:
		reader := &CIDRReader{
			Cidrs: make([]*CIDRManager.CIDRManager, 0),
		}
		cidrRegex := regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}/\d{1,2}\b`)
		cidrs := cidrRegex.FindAllString(scanner, -1)
		for _, cidr := range cidrs {
			reader.Cidrs = append(reader.Cidrs, CIDRManager.NewCIDR(cidr))
		}
		reader.CidrsLen = len(reader.Cidrs)
		return reader
	case EIP:
		return &StdinReader{scanner: bufio.NewScanner(strings.NewReader(scanner + "\n"))}
	}

	return nil
}

func NewFileReader(reader *bufio.Reader) *CIDRReader {
	cidrs := make([]*CIDRManager.CIDRManager, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}

		line = strings.TrimSpace(line)

		if !Utility.IsCIDRv4(line) {
			continue
		}

		cidrs = append(cidrs, CIDRManager.NewCIDR(line))
	}

	return &CIDRReader{Cidrs: cidrs, CidrsLen: len(cidrs)}
}

func (r *StdinReader) Next() (string, error) {
	if r.scanner.Scan() {
		return r.scanner.Text(), nil
	}
	if err := r.scanner.Err(); err != nil {
		return "", err
	}
	return "", io.EOF
}

func (r *StdinReader) Close() {
	r.scanner = nil
}

func (c *CIDRReader) Next() (string, error) {
	c.Lock()
	defer c.Unlock()
	numb := rand.Intn(c.CidrsLen)
	cidr := c.Cidrs[numb]
	ip, err := cidr.GetRandomIP()
	if errors.Is(err, CIDRManager.EOCIDR) {
		c.Cidrs = append(c.Cidrs[:numb], c.Cidrs[numb+1:]...)
		c.CidrsLen--
		return c.Next()
	}
	return ip, err
}

func (c *CIDRReader) Close() {
	c.Cidrs = nil
	c.CidrsLen = 0
}
