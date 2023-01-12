package pkg

import (
	"context"
	"fmt"
	"strings"

	"github.com/foomo/posh-providers/dreadl0ck/zeus"
	"github.com/foomo/posh-providers/foomo/gotsrpc"
	"github.com/foomo/posh-providers/onepassword"
	pkgcommand "github.com/foomo/posh-sandbox/posh/pkg/command"
	"github.com/foomo/posh/integration/ownbrew"
	"github.com/foomo/posh/pkg/cache"
	"github.com/foomo/posh/pkg/command"
	"github.com/foomo/posh/pkg/config"
	"github.com/foomo/posh/pkg/log"
	"github.com/foomo/posh/pkg/plugin"
	"github.com/foomo/posh/pkg/prompt"
	"github.com/foomo/posh/pkg/prompt/check"
	"github.com/foomo/posh/pkg/prompt/history"
	"github.com/foomo/posh/pkg/readline"
	"github.com/foomo/posh/pkg/require"
	"github.com/spf13/viper"
)

type Plugin struct {
	l        log.Logger
	c        cache.Cache
	commands command.Commands
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func New(l log.Logger) (plugin.Plugin, error) {
	inst := &Plugin{
		l:        l,
		c:        cache.MemoryCache{},
		commands: command.Commands{},
	}

	// add commands
	inst.commands.Add(
		gotsrpc.NewCommand(l, inst.c),
		zeus.NewCommand(l, inst.c),
		pkgcommand.NewGo(l, inst.c),
		pkgcommand.NewPrint(l),
		command.NewCache(l, inst.c),
		command.NewExit(l),
		command.NewHelp(l, inst.commands),
	)

	// Welcome
	if cmd, err := pkgcommand.NewWelcome(l,
		pkgcommand.WelcomeWithConfigKey("welcome"),
	); err != nil {
		return nil, err
	} else {
		inst.commands.Add(cmd)
	}

	// 1Password
	if onePassword, err := onepassword.New(l, inst.c, onepassword.WithTokenFilename(viper.GetString("onePassword.tokenFilename"))); err != nil {
		return nil, err
	} else if cmd, err := onepassword.NewCommand(l, onePassword); err != nil {
		return nil, err
	} else {
		inst.commands.Add(cmd)
	}
	return inst, nil
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

// Brew - $ posh brew => load plugin => call Brew()
func (p *Plugin) Brew(ctx context.Context, cfg config.Ownbrew) error {
	brew, err := ownbrew.New(p.l,
		ownbrew.WithDry(cfg.Dry),
		ownbrew.WithBinDir(cfg.BinDir),
		ownbrew.WithTapDir(cfg.TapDir),
		ownbrew.WithTempDir(cfg.TempDir),
		ownbrew.WithCellarDir(cfg.CellarDir),
		ownbrew.WithPackages(cfg.Packages...),
	)
	if err != nil {
		return err
	}
	return brew.Install(ctx)
}

func (p *Plugin) Require(ctx context.Context, cfg config.Require) error {
	return require.First(p.l,
		require.Envs(p.l, cfg.Envs),
		require.Packages(ctx, p.l, cfg.Packages),
		require.Scripts(ctx, p.l, cfg.Scripts),
		require.GitUser(ctx, p.l, require.GitUserName, require.GitUserEmail(`(.*)@(bestbytes\.com)`)),
	)
}

func (p *Plugin) Execute(ctx context.Context, args []string) error {
	r, err := readline.New(p.l)
	if err != nil {
		return err
	}

	if err := r.Parse(strings.Join(args, " ")); err != nil {
		return err
	}

	if cmd := p.commands.Get(r.Cmd()); cmd == nil {
		return fmt.Errorf("invalid [cmd] argument: %s", r.Cmd())
	} else {
		if value, ok := cmd.(command.Validator); ok {
			if err := value.Validate(ctx, r); err != nil {
				return err
			}
		}
		if err := cmd.Execute(ctx, r); err != nil {
			return err
		}
	}

	return nil
}

func (p *Plugin) Prompt(ctx context.Context, cfg config.Prompt) error {
	sh, err := prompt.New(p.l,
		prompt.WithTitle(cfg.Title),
		prompt.WithPrefix(cfg.Prefix),
		prompt.WithContext(ctx),
		prompt.WithCommands(p.commands),
		prompt.WithAliases(cfg.Aliases),
		prompt.WithCheckers(
			func(ctx context.Context, l log.Logger) check.Info {
				return check.Info{
					Name:   "one",
					Note:   "all good",
					Status: check.StatusSuccess,
				}
			},
			func(ctx context.Context, l log.Logger) check.Info {
				return check.Info{
					Name:   "two",
					Note:   "please take some action",
					Status: check.StatusFailure,
				}
			},
		),
		prompt.WithFileHistory(
			history.FileWithLimit(cfg.History.Limit),
			history.FileWithFilename(cfg.History.Filename),
			history.FileWithLockFilename(cfg.History.LockFilename),
		),
	)
	if err != nil {
		return err
	}
	return sh.Run()
}
