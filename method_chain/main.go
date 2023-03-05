package main

import (
	"errors"
	"fmt"
	"log"
)

type (
	ClientA struct {
	}
	ClientB struct {
	}
	ClientC struct {
	}
)

type Registry struct {
	ca  *ClientA
	cb  *ClientB
	cc  *ClientC
	err error
}

func (r *Registry) withClientA() *Registry {
	if r.err != nil {
		return r
	}
	fmt.Println("client A initialed")
	r.ca = &ClientA{}
	return r
}

func (r *Registry) withClientB() *Registry {
	if r.err != nil {
		return r
	}
	r.err = errors.New("error at initial client B")
	return r
}

func (r *Registry) withClientC() *Registry {
	if r.err != nil {
		return r
	}
	fmt.Println("client C initialed")
	r.cc = &ClientC{}
	return r
}

func main() {
	c := Registry{}
	d := c.withClientA().withClientB().withClientC()
	if d.err != nil {
		log.Fatalf("can not initial Clients due to %v", d.err)
	}
}
