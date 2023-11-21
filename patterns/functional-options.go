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

//
// main code:
//
//s := newServer(OptFunc(Opts{"10", 100, false}))
//s := newServer()
//fmt.Println("New Default Server : ", s)
//s1 := patterns.NewServer(patterns.WithMaxConn(99))
//fmt.Println("New Server with Max Conn : ", s1)
//s2 := patterns.NewServer(patterns.WithTLS)
//fmt.Println("New Server with TLS true : ", s2)
//s3 := patterns.NewServer(patterns.WithID("id"))
//fmt.Println("New Server with ID : ", s3)
//s4 := patterns.NewServer(patterns.WithTLS, patterns.WithMaxConn(99), patterns.WithID("newID"))
//fmt.Println("New Server with TLS true, ID and Max Conn : ", s4)
//
