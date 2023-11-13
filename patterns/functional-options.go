package patterns

type OptFunc func(*Opts)

type Opts struct {
	id      string
	maxConn int
	tls     bool
}

type Server struct {
	Opts
}

func DefaultOpts() Opts {
	return Opts{
		id:      "1",
		maxConn: 10,
		tls:     false,
	}
}

func WithTLS(opts *Opts) {
	opts.tls = true
}

func WithMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func WithID(s string) OptFunc {
	return func(opts *Opts) {
		opts.id = s
	}
}

func NewServer(opts ...OptFunc) *Server {
	o := DefaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}
