package bench

import (
	"github.com/antonito/gfile/pkg/session/bench"
	"github.com/antonito/gfile/pkg/session/common"
	"gitlab.com/manacore-backend/lib/pkg/log"
	"gopkg.in/urfave/cli.v1"
)

func handler(c *cli.Context) error {
	isMaster := c.Bool("master")

	sess := bench.NewWith(bench.Config{
		Master: isMaster,
		Configuration: common.Configuration{
			OnCompletion: func() {
			},
		},
	})
	return sess.Start()
}

// New creates the command
func New() cli.Command {
	log.Traceln("Installing 'bench' command")
	return cli.Command{
		Name:    "bench",
		Aliases: []string{"sb"},
		Usage:   "Benchmark the connexion",
		Action:  handler,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "master, m",
				Usage: "Is creating the SDP offer?",
			},
		},
	}
}