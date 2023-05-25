package record

import (
	"time"

	"github.com/denistakeda/mpass/internal/ports"
	"github.com/denistakeda/mpass/proto"
)

var _ ports.Record = (*bankCardRecord)(nil)

type bankCardRecord struct {
	id string

	cardCode string
	month    time.Month
	day      uint32
	code     uint
}

func bankCardRecordFromProto(id string, p *proto.BankCardRecord) *bankCardRecord {
	return &bankCardRecord{
		id:       id,
		cardCode: p.CardCode,
		month:    time.Month(p.Month),
		day:      p.Day,
		code:     uint(p.Code),
	}
}
