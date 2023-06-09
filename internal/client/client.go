package client

import (
	"io/ioutil"
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
		RegisterUser(login, password string) error
		LoginUser(login, password string) error
		Sync() error
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
			{
				Name:        "register",
				Usage:       "mpass register <login>",
				Description: "register to the mpass server",
				Action: func(cCtx *cli.Context) error {
					login := cCtx.Args().First()
					if login == "" {
						return errors.New("login was not provided")
					}

					password, err := newParamReader(params.Printer, params.Scanner, "Password").
						String().
						StripWhitespaces(false).
						NotEmpty(true).
						Read()
					if err != nil {
						return err
					}

					err = params.ClientService.RegisterUser(login, password)
					if err != nil {
						return err
					}

					params.Printer.Printf("user %q was successfully registered\n", login)

					return nil
				},
			},
			{
				Name:        "login",
				Usage:       "mpass login <login>",
				Description: "login to the mpass server",
				Action: func(cCtx *cli.Context) error {
					login := cCtx.Args().First()
					if login == "" {
						return errors.New("login was not provided")
					}

					password, err := newParamReader(params.Printer, params.Scanner, "Password").
						String().
						StripWhitespaces(false).
						NotEmpty(true).
						Read()
					if err != nil {
						return err
					}

					err = params.ClientService.LoginUser(login, password)
					if err != nil {
						return err
					}

					params.Printer.Printf("user %q was successfully logged in\n", login)

					return nil
				},
			},
			{
				Name:        "sync",
				Usage:       "mpass sync",
				Description: "sync local and server database",
				Action: func(cCtx *cli.Context) error {
					return params.ClientService.Sync()
				},
			},
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
						Description: "add the text to the store with defined key",
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
					{
						Name:        "file",
						Usage:       "mpass set file <key> <file_path>",
						Description: "add the file to the store with defined key",
						Action: func(cCtx *cli.Context) error {
							key := cCtx.Args().First()
							if key == "" {
								return errors.New("key was not provided")
							}

							filePath := cCtx.Args().Get(1)
							if filePath == "" {
								return errors.New("file_path was not provided")
							}

							data, err := ioutil.ReadFile(filePath)
							if err != nil {
								return errors.Wrapf(err, "failed to read file %q", filePath)
							}

							rec := record.NewBinaryRecord(key, data)

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
