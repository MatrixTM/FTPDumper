package CIDRManager

import (
	"bytes"
	"encoding/binary"
	"net"
	"sync"
)

type BSet struct {
	set   map[[4]byte]struct{}
	mutex sync.RWMutex
}

func NewBSet() *BSet {
	return &BSet{
		set:   make(map[[4]byte]struct{}),
		mutex: sync.RWMutex{},
	}
}

func (h *BSet) Add(item string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	compressed := compressIPv4(item)

	h.set[[4]byte(compressed)] = struct{}{}
}

func (h *BSet) Contains(item string) bool {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	compressed := compressIPv4(item)

	_, exists := h.set[[4]byte(compressed)]
	return exists
}

func (h *BSet) Count() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return len(h.set)
}

func compressIPv4(ip string) []byte {
	var compressed bytes.Buffer

	binary.Write(&compressed, binary.BigEndian, net.ParseIP(ip).To4())

	return compressed.Bytes()
}
