package main

import (
	"errors"
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}

	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		panic(err)
	}

	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arith: %d*%d=%d, %d/%d=%d remainder %d\n", args.A, args.B, reply, args.A, args.B, quo.Quo, quo.Rem)
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
