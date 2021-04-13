package balance

import "fmt"

type BalancerMgr struct {
	allBalancer map[string]Balancer
}

//default BalancerMgr
var mgr = BalancerMgr{
	allBalancer: make(map[string]Balancer),
}

func RegisterBalancer(name string, balancer Balancer) {
	mgr.allBalancer[name] = balancer
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	//select
	balancer, ok := mgr.allBalancer[name]

	if !ok {
		err = fmt.Errorf("Not found %s balancer", name)
		return
	}
	inst, err = balancer.DoBalance(insts)
	return
}
