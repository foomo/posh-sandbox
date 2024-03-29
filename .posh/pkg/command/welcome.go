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
		l    log.Logger
		cfg  config.Welcome
		name string
	}
	WelcomeOption func(*Welcome) error
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func WelcomeWithConfig(v config.Welcome) WelcomeOption {
	return func(o *Welcome) error {
		o.cfg = v
		return nil
	}
}

func WelcomeWithConfigKey(v string) WelcomeOption {
	return func(o *Welcome) error {
		return viper.UnmarshalKey(v, &o.cfg)
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewWelcome(l log.Logger, opts ...WelcomeOption) (*Welcome, error) {
	inst := &Welcome{
		l:    l,
		name: "welcome",
	}
	for _, opt := range opts {
		if opt != nil {
			if err := opt(inst); err != nil {
				return nil, err
			}
		}
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

func (c *Welcome) Help(ctx context.Context, r *readline.Readline) string {
	return `Print a welcome message

Usage:
welcome
`
}
