package record

import (
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Record = (*bankCardRecord)(nil)

type bankCardRecord struct {
	id             string
	lastUpdateDate time.Time

	cardCode string
	month    time.Month
	day      uint32
	code     uint
}

func bankCardRecordFromProto(id string, lastUpdateDate time.Time, p *proto.BankCardRecord) *bankCardRecord {
	return &bankCardRecord{
		id:             id,
		lastUpdateDate: lastUpdateDate,

		cardCode: p.CardCode,
		month:    time.Month(p.Month),
		day:      p.Day,
		code:     uint(p.Code),
	}
}

func (r *bankCardRecord) GetId() string {
	return r.id
}

func (r *bankCardRecord) GetLastUpdateDate() time.Time {
	return r.lastUpdateDate
}

func (r *bankCardRecord) ToProto() *proto.Record {
	return &proto.Record{
		Id:             r.id,
		LastUpdateDate: timestamppb.New(r.lastUpdateDate),

		Record: &proto.Record_BankCardRecord{
			BankCardRecord: &proto.BankCardRecord{
				CardCode: r.cardCode,
				Month:    uint32(r.month),
				Day:      r.day,
				Code:     uint32(r.code),
			},
		},
	}
}
