package record_store

import (
	"context"
	"fmt"

	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/errgroup"
)

type tableNameT string

const (
	loginPasswordTableName tableNameT = "login_password_record"
	binaryTableName                   = "binary_record"
	textTableName                     = "text_record"
	bankCardTableName                 = "bank_card_record"
)

type dbStore struct {
	db *sqlx.DB
}

func NewWithDb(db *sqlx.DB) *dbStore {
	return &dbStore{db: db}
}

func (s *dbStore) AddRecords(ctx context.Context, login string, records []record.Record) error {
	panic("implement!")
}

func (s *dbStore) AllRecords(ctx context.Context, login string) ([]record.Record, error) {
	var (
		res                  []record.Record
		loginPasswordRecords []record.Record
		binaryRecords        []record.Record
		textRecords          []record.Record
		bankCardRecords      []record.Record
	)

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(s.getRecords(gCtx, login, loginPasswordTableName, &loginPasswordRecords))
	g.Go(s.getRecords(gCtx, login, binaryTableName, &binaryRecords))
	g.Go(s.getRecords(gCtx, login, textTableName, &textRecords))
	g.Go(s.getRecords(gCtx, login, bankCardTableName, &bankCardRecords))

	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("failed to fetch all the records: %w", err)
	}

	res = append(res, loginPasswordRecords...)
	res = append(res, binaryRecords...)
	res = append(res, textRecords...)
	res = append(res, bankCardRecords...)

	return res, nil
}

func (s *dbStore) getRecords(ctx context.Context, login string, tableName tableNameT, recs *[]record.Record) func() error {
	return func() error {
		return s.db.SelectContext(
			ctx, recs,
			"select * from $1 where user_login=$2",
			tableName,
			login,
		)
	}
}
