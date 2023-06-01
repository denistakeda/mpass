package client

import (
	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

type (
	printer interface {
		Printf(format string, a ...any)
	}
	scanner interface {
		Readln() (string, error)
	}
	clientService interface {
		SetRecord(record.Record) error
		GetRecord(string) (record.Record, error)
	}
)

type NewClientParams struct {
	Printer       printer
	Scanner       scanner
	ClientService clientService
}

func New(params NewClientParams) *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			// {
			// 	Name:  "sync",
			// 	Usage: "sync local database with server",
			// 	Action: func(cCtx *cli.Context) error {
			// 		panic("implement")
			// 	},
			// },
			// {
			// 	Name:  "login",
			// 	Usage: "login to the server",
			// 	Action: func(cCtx *cli.Context) error {
			// 		panic("implement")
			// 	},
			// },
			// {
			// 	Name:  "register",
			// 	Usage: "register a new user on the server",
			// 	Action: func(cCtx *cli.Context) error {
			// 		panic("implement")
			// 	},
			// },
			// {
			// 	Name:  "get",
			// 	Usage: "get key from database",
			// 	Action: func(cCtx *cli.Context) error {
			// 		panic("implement")
			// 	},
			// },
			{
				Name: "set",
				Subcommands: []*cli.Command{
					{
						Name:        "password",
						Usage:       "mpass set password <login>",
						Description: "add the login/password item to the store",
						Action: func(cCtx *cli.Context) error {
							login := cCtx.Args().First()
							if login == "" {
								return errors.New("login was not provided")
							}

							params.Printer.Printf("Enter your password for login %q: ", login)

							password, err := params.Scanner.Readln()
							if err != nil {
								return errors.New("failed to read password")
							}
							if password == "" {
								return errors.New("password should not be empty")
							}

							rec := record.NewLoginPasswordRecord(login, password)

							return params.ClientService.SetRecord(rec)
						},
					},
				},
			},
			{
				Name:        "get",
				Usage:       "mpass get <key>",
				Description: "gets the value by key from the local database",
				Action: func(cCtx *cli.Context) error {
					key := cCtx.Args().First()
					if key == "" {
						return errors.New("key is not provided")
					}

					rec, err := params.ClientService.GetRecord(key)
					if err != nil {
						return err
					}

					rec.ProvideToClient(params.Printer)

					return nil
				},
			},
		},
	}
}
