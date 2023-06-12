package record

import (
	"encoding/gob"
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Record = (*bankCardRecord)(nil)

func init() {
	gob.Register(&bankCardRecord{})
}

type bankCardRecord struct {
	ID             string    `db:"id"`
	LastUpdateDate time.Time `db:"last_update_date"`

	CardNumber string     `db:"card_number"`
	Month      time.Month `db:"month"`
	Day        uint32     `db:"day"`
	Code       uint       `db:"code"`
}

func NewBankCardRecord(cardNumber string, month time.Month, day uint32, code uint) *bankCardRecord {
	return &bankCardRecord{
		ID:             cardNumber,
		LastUpdateDate: time.Now(),

		CardNumber: cardNumber,
		Month:      month,
		Day:        day,
		Code:       code,
	}
}

func bankCardRecordFromProto(id string, lastUpdateDate time.Time, p *proto.BankCardRecord) *bankCardRecord {
	return &bankCardRecord{
		ID:             id,
		LastUpdateDate: lastUpdateDate,

		CardNumber: p.CardCode,
		Month:      time.Month(p.Month),
		Day:        p.Day,
		Code:       uint(p.Code),
	}
}

func (r *bankCardRecord) GetId() string {
	return r.ID
}

func (r *bankCardRecord) GetLastUpdateDate() time.Time {
	return r.LastUpdateDate
}

func (r *bankCardRecord) ToProto() *proto.Record {
	return &proto.Record{
		Id:             r.ID,
		LastUpdateDate: timestamppb.New(r.LastUpdateDate),

		Record: &proto.Record_BankCardRecord{
			BankCardRecord: &proto.BankCardRecord{
				CardCode: r.CardNumber,
				Month:    uint32(r.Month),
				Day:      r.Day,
				Code:     uint32(r.Code),
			},
		},
	}
}

func (r *bankCardRecord) ProvideToClient(printer printer) error {
	printer.Printf("Card Number: %s\nDate: %d/%d   Code: %s", r.CardNumber, r.Month, r.Day, r.Code)
	return nil
}
