package record

import (
	"time"

	"github.com/denistakeda/mpass/proto"
)

type Record interface {
	GetId() string
	GetLastUpdateDate() time.Time
	ToProto() *proto.Record

	ProvideToClient(printer printer) error
}

type printer interface {
	Printf(format string, a ...any)
}

func FromProto(rec *proto.Record) Record {
	lastUpdateDate := rec.LastUpdateDate.AsTime()
	switch i := rec.Record.(type) {
	case *proto.Record_LoginPasswordRecord:
		return loginPasswordRecordFromProto(rec.Id, lastUpdateDate, i.LoginPasswordRecord)
	case *proto.Record_TextRecord:
		return textRecordFromProto(rec.Id, lastUpdateDate, i.TextRecord)
	case *proto.Record_BinaryRecord:
		return binaryRecordFromProto(rec.Id, lastUpdateDate, i.BinaryRecord)
	case *proto.Record_BankCardRecord:
		return bankCardRecordFromProto(rec.Id, lastUpdateDate, i.BankCardRecord)
	default:
		// Should never happen
		return nil
	}
}
