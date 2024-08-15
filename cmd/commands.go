package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"omni-balance/internal/actions"
	"omni-balance/internal/daemons"

	log "omni-balance/utils/logging"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
)

var (
	Commands = []*cli.Command{
		{
			Name: "balance",
			// 获取配置文件中所有token以及地址的余额
			Usage: "Get the balance of all tokens and addresses in the configuration file.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "c",
					Usage: "config file path",
				},
			},
			Action: func(c *cli.Context) error {
				result, err := daemons.FindAllChainBalance(context.TODO(), c.String("c"), true)
				if err != nil {
					return errors.Wrap(err, "find all chain balance")
				}
				w := table.NewWriter()
				w.AppendHeader(table.Row{"Chain", "Token", "Address", "Balance"})
				for chainName, chainItems := range result {
					for tokenName, TokenItems := range chainItems {
						for address, balance := range TokenItems {
							if balance.LessThan(decimal.RequireFromString("0")) {
								w.AppendRow(table.Row{chainName, tokenName, address, "not found in config"})
								continue
							}
							w.AppendRow(table.Row{chainName, tokenName, address, balance})
						}
					}
				}
				fmt.Println("\n", w.Render())
				return nil
			},
		},
		{
			Name:        "gate_liquidity",
			Usage:       "Create an order for the liquidity of Gate.",
			Description: "Create an order for the liquidity of Gate.",
			Action:      actions.DoGateLiquidity,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "server",
					Usage: "omni-balance server port",
					Value: "http://127.0.0.1:8080",
				},
				&cli.StringFlag{
					Name:     "apiKey",
					Required: true,
					Usage:    "omni-balance API key in config",
				},
				&cli.StringFlag{
					Name:     "tokenName",
					Required: true,
					Usage:    "a token name must be in config",
				},
				&cli.StringFlag{
					Name:     "fromChain",
					Required: true,
					Usage:    "deposit to the Gate exchange from that chain.",
				},
				&cli.StringFlag{
					Name:     "toChain",
					Required: true,
					Usage:    "Withdraw to that chain from the Gate exchange.",
				},
				&cli.StringFlag{
					Name:     "amount",
					Required: true,
					Usage:    "The number of tokens that need to be rebalanced.",
				},
				&cli.StringFlag{
					Name:     "address",
					Required: true,
					Usage:    "Use that address to rebalance, and note that this address must exist in the configuration file.",
				},
			},
		},

		{
			Name:  "del_order",
			Usage: "delete order by id",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "id",
					Usage: "order id",
				},
				&cli.StringFlag{
					Name:  "server",
					Usage: "server host, example: http://127.0.0.1:8080",
					Value: "http://127.0.0.1:8080",
				},
			},
			Action: func(c *cli.Context) error {
				u, err := url.Parse(c.String("server"))
				if err != nil {
					return errors.Wrap(err, "parse server url")
				}
				u.RawPath = "/remove_order"
				u.Path = u.RawPath
				var body = bytes.NewBuffer(nil)
				err = json.NewEncoder(body).Encode(map[string]interface{}{
					"id": c.Int("id"),
				})
				if err != nil {
					return errors.Wrap(err, "encode body")
				}
				resp, err := http.Post(u.String(), "application/json", body)
				if err != nil {
					return errors.Wrap(err, "post")
				}
				defer resp.Body.Close()
				data, _ := io.ReadAll(resp.Body)
				if resp.StatusCode != http.StatusOK {
					return errors.Errorf("http status code: %d, body is: %s", resp.StatusCode, data)
				}
				log.Infof("delete order #%d success", c.Int64("id"))
				return nil
			},
		},
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
)
