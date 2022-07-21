package rlimiter

type Config struct {
	Sender   Sender
	Worker   uint8
	MaxLimit int64
}

func (c *Config) SetSender(s Sender) *Config {
	c.Sender = s

	return c
}

func (c *Config) SetWorker(w uint8) *Config {
	c.Worker = w

	return c
}

func (c *Config) SetMaxLimit(m int64) *Config {
	c.MaxLimit = m

	return c
}

func NewConfig() *Config {
	return &Config{}
}

func MergeConfigs(confs ...*Config) (*Config, error) {
	cn := NewConfig()

	for _, conf := range confs {
		if conf.Worker > 0 {
			cn.Worker = conf.Worker
		}

		if conf.MaxLimit > 0 {
			cn.MaxLimit = conf.MaxLimit
		}

		if conf.Sender != nil {
			cn.Sender = conf.Sender
		}
	}

	if cn.Worker == 0 {
		cn.Worker = 1
	}

	if cn.MaxLimit < 0 {
		cn.MaxLimit = 1
	}

	if cn.Sender == nil {
		return nil, ErrSenderRequired
	}

	return cn, nil
}
