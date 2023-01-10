package command

import (
	"context"

	"github.com/foomo/posh/pkg/command/tree"
	"github.com/foomo/posh/pkg/log"
	"github.com/foomo/posh/pkg/prompt"
	"github.com/foomo/posh/pkg/readline"
)

type Print struct {
	l           log.Logger
	commandTree *tree.Root
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewPrint(l log.Logger) *Print {
	inst := &Print{
		l: l.Named("print"),
	}
	inst.commandTree = &tree.Root{
		Name:        "print",
		Description: "print a message",
		Node: &tree.Node{
			Args: tree.Args{
				{
					Name:   "message",
					Repeat: false,
					Suggest: func(ctx context.Context, t *tree.Root, r *readline.Readline) []prompt.Suggest {
						return []prompt.Suggest{
							{Text: "hello world"},
						}
					},
				},
			},
			Execute: func(ctx context.Context, r *readline.Readline) error {
				l.Info(r.Args())
				return nil
			},
		},
	}
	return inst
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (c *Print) Name() string {
	return c.commandTree.Name
}

func (c *Print) Description() string {
	return c.commandTree.Description
}

func (c *Print) Complete(ctx context.Context, r *readline.Readline, d prompt.Document) []prompt.Suggest {
	return c.commandTree.RunCompletion(ctx, r)
}

func (c *Print) Execute(ctx context.Context, r *readline.Readline) error {
	return c.commandTree.RunExecution(ctx, r)
}

func (c *Print) Help() string {
	return `Print a message

Usage:
  welcome [message]
`
}
