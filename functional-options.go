package main

type OptFunc func(*Opts)

type Opts struct {
	id      string
	maxConn int
	tls     bool
}

type Server struct {
	Opts
}

func defaultOpts() Opts {
	return Opts{
		id:      "1",
		maxConn: 10,
		tls:     false,
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func withID(s string) OptFunc {
	return func(opts *Opts) {
		opts.id = s
	}
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}
