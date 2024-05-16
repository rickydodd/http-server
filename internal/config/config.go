package config

type configBuilder struct {
	addr *string
	port *uint16
}

type config struct {
	Addr string
	Port uint16
}

func Config() *configBuilder {
	return &configBuilder{}
}

func (b *configBuilder) Build() (config, error) {
	conf := config{}

	if b.addr == nil {
		conf.Addr = "0.0.0.0"
	} else {
		conf.Addr = *b.addr
	}

	if b.port == nil {
		conf.Port = 80
	} else {
		conf.Port = *b.port
	}

	return conf, nil
}

func (b *configBuilder) Addr(addr string) *configBuilder {
	b.addr = &addr
	return b
}

func (b *configBuilder) Port(port uint16) *configBuilder {
	b.port = &port
	return b
}
