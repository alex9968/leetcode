package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalancer("random", &RodomBalance{})
}

type RodomBalance struct {
}

func (r *RodomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}
