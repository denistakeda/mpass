package record_store

import (
	"context"
	"database/sql"
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
	if len(records) == 0 {
		return nil
	}

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start a transaction: %w", err)
	}

	for _, rec := range records {
		switch r := rec.(type) {
		case *record.LoginPasswordRecord:
			err = upsertLoginPasswordRecord(ctx, tx, r, login)
		case *record.BankCardRecord:
			err = upsertBankCardRecord(ctx, tx, r, login)
		case *record.BinaryRecord:
			err = upsertBinaryRecord(ctx, tx, r, login)
		case *record.TextRecord:
			err = upsertTextRecord(ctx, tx, r, login)
		default:
			return fmt.Errorf("unknown record type: %v", r)
		}
	}

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to add records: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit records: %w", err)
	}

	return nil
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
			fmt.Sprintf("select * from %s where user_login=$2", tableName),
			login,
		)
	}
}

func upsertLoginPasswordRecord(ctx context.Context, tx *sqlx.Tx, r *record.LoginPasswordRecord, userLogin string) error {
	var old record.LoginPasswordRecord
	err := tx.GetContext(ctx, &old, "select * from login_password_record where id=$1", r.ID)

	if err != nil {
		_, err = tx.ExecContext(ctx,
			"insert into login_password_record (id, last_update_date, login, password, user_login) values ($1, $2, $3, $4, $5)",
			r.ID, r.LastUpdateDate, r.Login, r.Password, userLogin,
		)
	} else if r.LastUpdateDate.After(old.LastUpdateDate) {
		_, err = tx.ExecContext(ctx,
			"update login_password_record set login=$1, password=$2",
			r.Login, r.Password,
		)
	}

	return err
}

func upsertBankCardRecord(ctx context.Context, tx *sqlx.Tx, r *record.BankCardRecord, userLogin string) error {
	var old record.BankCardRecord
	err := tx.GetContext(ctx, &old, "select * from bank_card_record where id=$1", r.ID)

	if err != nil {
		_, err = tx.ExecContext(ctx,
			"insert into bank_card_record (id, last_update_date, card_number, month, day, code, user_login) values ($1, $2, $3, $4, $5, $6, $7)",
			r.ID, r.LastUpdateDate, r.CardNumber, r.Month, r.Day, r.Code, userLogin,
		)
	} else if r.LastUpdateDate.After(old.LastUpdateDate) {
		_, err = tx.ExecContext(ctx,
			"update bank_card_record set card_number=$1, month=$2, day=$3, code=$4",
			r.CardNumber, r.Month, r.Day, r.Code,
		)
	}

	return err
}

func upsertBinaryRecord(ctx context.Context, tx *sqlx.Tx, r *record.BinaryRecord, userLogin string) error {
	var old record.BinaryRecord
	err := tx.GetContext(ctx, &old, "select * from binary_record where id=$1", r.ID)

	if err != nil {
		_, err = tx.ExecContext(ctx,
			"insert into binary_record (id, last_update_date, binary, user_login) values ($1, $2, $3, $4)",
			r.ID, r.LastUpdateDate, r.Binary, userLogin,
		)
	} else if r.LastUpdateDate.After(old.LastUpdateDate) {
		_, err = tx.ExecContext(ctx,
			"update binary_record set binary=$1",
			r.Binary,
		)
	}

	return err
}

func upsertTextRecord(ctx context.Context, tx *sqlx.Tx, r *record.TextRecord, userLogin string) error {
	var old record.TextRecord
	err := tx.GetContext(ctx, &old, "select * from text_record where id=$1", r.ID)

	if err != nil {
		_, err = tx.ExecContext(ctx,
			"insert into text_record (id, last_update_date, text, user_login) values ($1, $2, $3, $4)",
			r.ID, r.LastUpdateDate, r.Text, userLogin,
		)
	} else if r.LastUpdateDate.After(old.LastUpdateDate) {
		_, err = tx.ExecContext(ctx,
			"update text_record set text=$1",
			r.Text,
		)
	}

	return err
}
