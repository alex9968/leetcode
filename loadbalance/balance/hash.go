package balance

import (
	"errors"
	"fmt"
	"hash/crc32"
	"math/rand"
)

func init() {
	RegisterBalancer("hash", &HashBalance{})
}

type HashBalance struct {
	Name string
	Age  int
}

func (h *HashBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)

	if lens == 0 {
		err = errors.New("No backend instance")
		return
	}

	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]
	return
}
