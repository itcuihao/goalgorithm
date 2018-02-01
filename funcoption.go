package main

import (
	"fmt"
)

//  函数式选项模式
// https://lingchao.xin/post/functional-options-pattern-in-go.html

var defaultStuffClientOptions = StuffClientOptions{
	Retries: 3,
	Timeout: 2,
}

type StuffClientOptions struct {
	Retries int
	Timeout int
}

type StuffClientOption func(*StuffClientOptions)

func WithRetries(r int) StuffClientOption {
	return func(o *StuffClientOptions) {
		o.Retries = r
	}
}

func WithTimeout(t int) StuffClientOption {
	return func(o *StuffClientOptions) {
		o.Timeout = t
	}
}

type StuffClient interface {
	DoStuff() error
}

type stuffClient struct {
	conn    Connection
	timeout int
	retries int
}

type Connection struct{}

func NewStuffClient(conn Connection, opts ...StuffClientOption) StuffClient {
	options := defaultStuffClientOptions
	for _, o := range opts {
		o(&options)
	}
	return &stuffClient{
		conn:    conn,
		retries: options.Retries,
		timeout: options.Timeout,
	}
}

func (c stuffClient) DoStuff() error {
	return nil
}

func main() {
	x := NewStuffClient(Connection{})
	fmt.Println(x)

	x = NewStuffClient(Connection{}, WithRetries(1), WithRetries(2))
	fmt.Println(x)
}
