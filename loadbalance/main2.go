package main

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"strconv"
	"time"
)

type Instance struct {
	port int
	host string
}

func (i *Instance) GetHost() string {
	return i.host
}
func (i *Instance) GetPort() int {
	return i.port
}

func (i *Instance) String() string {
	return i.host + ":" + strconv.Itoa(i.port)
}

type Balancer struct {
	Index   int
	Clients []*Instance
}

func NewBalancer() *Balancer {
	bl := Balancer{}
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.169.%d.%d", rand.Intn(255), rand.Intn(255))
		bl.Clients = append(bl.Clients, &Instance{host: host, port: 9090})
	}
	return &bl
}

func (b *Balancer) GetByHash() *Instance {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % len(b.Clients)
	return b.Clients[index]
}
func (b *Balancer) GetByRand() *Instance {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(b.Clients))
	return b.Clients[index]
}

func main() {

	bl := NewBalancer()

	for {
		fmt.Printf("%#v\n", bl.GetByHash())
		time.Sleep(time.Second)
	}
}
