package command

import (
	"context"

	"github.com/foomo/posh/pkg/command/tree"
	"github.com/foomo/posh/pkg/log"
	"github.com/foomo/posh/pkg/prompt/goprompt"
	"github.com/foomo/posh/pkg/readline"
)

type Print struct {
	l           log.Logger
	commandTree tree.Root
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewPrint(l log.Logger) *Print {
	inst := &Print{
		l: l.Named("print"),
	}
	inst.commandTree = tree.New(&tree.Node{
		Name:        "print",
		Description: "print a message",
		Args: tree.Args{
			{
				Name:   "message",
				Repeat: false,
				Suggest: func(ctx context.Context, t tree.Root, r *readline.Readline) []goprompt.Suggest {
					return []goprompt.Suggest{
						{Text: "hello world"},
					}
				},
			},
		},
		Execute: func(ctx context.Context, r *readline.Readline) error {
			l.Info(r.Args())
			return nil
		},
	})
	return inst
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (c *Print) Name() string {
	return c.commandTree.Node().Name
}

func (c *Print) Description() string {
	return c.commandTree.Node().Description
}

func (c *Print) Complete(ctx context.Context, r *readline.Readline, d goprompt.Document) []goprompt.Suggest {
	return c.commandTree.Complete(ctx, r)
}

func (c *Print) Execute(ctx context.Context, r *readline.Readline) error {
	return c.commandTree.Execute(ctx, r)
}

func (c *Print) Help(ctx context.Context, r *readline.Readline) string {
	return c.commandTree.Help(ctx, r)
}
