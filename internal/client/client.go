package client

import (
	"time"

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

							password, err := newParamReader(params.Printer, params.Scanner, "Password").
								String().
								StripWhitespaces(true).
								NotEmpty(true).
								Read()
							if err != nil {
								return err
							}

							rec := record.NewLoginPasswordRecord(login, password)

							return params.ClientService.SetRecord(rec)
						},
					},
					{
						Name:        "card",
						Usage:       "mpass set card",
						Description: "add the bank card to the store",
						Action: func(cCtx *cli.Context) error {
							cardNumber, err := newParamReader(params.Printer, params.Scanner, "Card Number").
								String().
								StripWhitespaces(true).
								NotEmpty(true).
								Read()
							if err != nil {
								return err
							}

							month, err := newParamReader(params.Printer, params.Scanner, "Month").Month().Read()
							if err != nil {
								return err
							}

							day, err := newParamReader(params.Printer, params.Scanner, "Day").Day().Read()
							if err != nil {
								return err
							}

							cardCode, err := newParamReader(params.Printer, params.Scanner, "Card Code").NumRange(1, 999).Read()
							if err != nil {
								return err
							}

							rec := record.NewBankCardRecord(cardNumber, time.Month(month), uint32(day), uint(cardCode))

							return params.ClientService.SetRecord(rec)
						},
					},
					{
						Name:        "text",
						Usage:       "mpass set text <key>",
						Description: "add the text to the store",
						Action: func(cCtx *cli.Context) error {
							key := cCtx.Args().First()
							if key == "" {
								return errors.New("key was not provided")
							}

							text, err := newParamReader(params.Printer, params.Scanner, "Text").
								String().
								StripWhitespaces(false).
								NotEmpty(true).
								Read()
							if err != nil {
								return err
							}

							rec := record.NewTextRecord(key, text)

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

					err = rec.ProvideToClient(params.Printer)
					if err != nil {
						return err
					}

					return nil
				},
			},
		},
	}
}
