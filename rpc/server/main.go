package main

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		panic(e)
	}

	go http.Serve(l, nil)

	select {}
}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith struct{}

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
