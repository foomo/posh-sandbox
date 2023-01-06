package command

import (
	"context"

	"github.com/foomo/posh-sandbox/posh/pkg/config"
	"github.com/foomo/posh/pkg/log"
	"github.com/foomo/posh/pkg/readline"
	"github.com/spf13/viper"
)

type (
	Welcome struct {
		l         log.Logger
		cfg       config.Welcome
		name      string
		configKey string
	}
	WelcomeOption func(*Welcome) error
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func WelcomeWithConfigKey(v string) WelcomeOption {
	return func(o *Welcome) error {
		o.configKey = v
		return nil
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewWelcome(l log.Logger, opts ...WelcomeOption) (*Welcome, error) {
	inst := &Welcome{
		l:         l,
		name:      "welcome",
		configKey: "welcome",
	}
	for _, opt := range opts {
		if opt != nil {
			if err := opt(inst); err != nil {
				return nil, err
			}
		}
	}
	if err := viper.UnmarshalKey(inst.configKey, &inst.cfg); err != nil {
		return nil, err
	}
	return inst, nil
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (c *Welcome) Name() string {
	return c.name
}

func (c *Welcome) Description() string {
	return "print a welcome message"
}

func (c *Welcome) Execute(ctx context.Context, r *readline.Readline) error {
	c.l.Success(c.cfg.Message)
	return nil
}

func (c *Welcome) Help() string {
	return `Print a welcome message

Usage:
  welcome
`
}