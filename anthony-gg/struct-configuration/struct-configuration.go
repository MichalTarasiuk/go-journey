package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

func withTls(opts *Opts) {
	opts.tls = true
}

func withMaxConn(maxConn int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = maxConn
	}
}

type Server struct {
	opts Opts
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}

	return &Server{
		opts: o,
	}
}

func main() {
	s := newServer(withTls, withMaxConn(100))
	fmt.Println(s)
}
