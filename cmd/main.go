package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"omni-balance/internal/daemons"
	_ "omni-balance/internal/daemons/cross_chain"
	_ "omni-balance/internal/daemons/monitor"
	_ "omni-balance/internal/daemons/rebalance"
	_ "omni-balance/internal/daemons/token_price"
	"omni-balance/internal/db"
	"omni-balance/internal/models"
	"omni-balance/utils"
	"omni-balance/utils/configs"
	"omni-balance/utils/notice"
	"omni-balance/utils/provider"
	_ "omni-balance/utils/provider/bridge/darwinia"
	_ "omni-balance/utils/provider/bridge/helix"
	_ "omni-balance/utils/provider/bridge/okx"
	_ "omni-balance/utils/provider/bridge/routernitro"
	_ "omni-balance/utils/provider/cex/gate"
	_ "omni-balance/utils/provider/dex/uniswap"
	_ "omni-balance/utils/wallets/safe"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	config        = new(configs.Config)
	defaultsUsage = flag.Usage
	ctx, cancel   = context.WithCancel(context.TODO())
	// fill it when building
	version       string
	commitMessage string
	commitTime    string
)

func Usage(_ *cli.Context) error {
	defaultsUsage()
	fmt.Printf("Supported providers:\n")
	for providerType, providerFns := range provider.ListProviders() {
		fmt.Printf(" %s:\n", providerType)
		for _, fn := range providerFns {
			providerObj, err := fn(*config, true)
			if err != nil {
				logrus.Panicf("init provider error: %v", err)
			}
			fmt.Printf("  %s:\n", providerObj.Name())
			for _, v := range providerObj.Help() {
				fmt.Printf("   %s\n", v)
			}
		}
	}
	return nil
}

func Action(cli *cli.Context) error {
	if err := initConfig(ctx, cli.Bool("placeholder"), cli.String("conf"), cli.String("port")); err != nil {
		return errors.Wrap(err, "init config")
	}
	if config.Debug {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors:          true,
			ForceQuote:             true,
			DisableLevelTruncation: false,
			QuoteEmptyFields:       true,
		})
	}

	if !config.Debug {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	if err := notice.Init(notice.Type(config.Notice.Type), config.Notice.Config); err != nil {
		logrus.Warnf("init notice error: %v", err)
	}
	notice.SetMsgInterval(config.Notice.Interval)

	if err := db.InitDb(*config); err != nil {
		return errors.Wrap(err, "init db")
	}

	if err := db.DB().AutoMigrate(
		new(models.Order),
		new(models.OrderProcess),
		new(models.TokenPrice)); err != nil {
		return errors.Wrap(err, "auto migrate db")
	}

	if err := daemons.Run(ctx, *config); err != nil {
		return errors.Wrap(err, "run daemons")
	}
	utils.FinishInit()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
	time.Sleep(time.Second * 5)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "omni-balance"
	app.Action = Action
	app.Commands = []*cli.Command{
		{
			Name:    "version",
			Usage:   "show version",
			Aliases: []string{"v"},
			Action: func(c *cli.Context) error {
				fmt.Printf("Version: %s\n", version)
				fmt.Printf("Commit: %s\n", commitMessage)
				fmt.Printf("Build time: %s\n", commitTime)
				return nil
			},
		},
		{
			Name:   "list",
			Usage:  "list supported providers and docs",
			Action: Usage,
		},
		{
			Name:  "tasks",
			Usage: "list supported tasks",
			Action: func(_ *cli.Context) error {
				fmt.Println(daemons.Help())
				return nil
			},
		},
		{
			Name:  "example",
			Usage: "create a example config file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "output",
					Usage:   "output file path",
					Value:   "./config.yaml.example",
					Aliases: []string{"o"},
				},
			},
			Action: func(c *cli.Context) error {
				if err := CreateExampleConfig(c.String("output")); err != nil {
					return errors.Wrap(err, "create example config")
				}
				return nil
			},
		},
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "conf",
			Aliases: []string{"c"},
			Usage:   "config file path",
			Value:   "./config.yaml",
		},
		&cli.BoolFlag{
			Name: "placeholder",
			Usage: fmt.Sprintf("enable placeholder, you can use placeholder to replace private key, Example: Fill '{{privateKey}}' in config.yaml."+
				"Run with -p to enable placeholder, Example: SERVER_PORT=:8080 %s -c ./config.yaml -p"+
				"Waiting for 'waiting for placeholder...' log, send placeholder data according to the prompt.", os.Args[0]),
			Aliases: []string{"p"},
		},
		&cli.StringFlag{
			Name:  "port",
			Usage: "When the placeholder parameter is set to true, you can specify and set the listening address of the HTTP server that receives the placeholder.",
			Value: ":8080",
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Error(err)
	}
}
