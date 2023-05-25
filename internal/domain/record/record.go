package record

import (
	"github.com/denistakeda/mpass/internal/ports"
	"github.com/denistakeda/mpass/proto"
)

func FromProto(rec *proto.Record) ports.Record {
	switch i := rec.Record.(type) {
	case *proto.Record_LoginPasswordRecord:
		return loginPasswordRecordFromProto(rec.Id, i.LoginPasswordRecord)
	case *proto.Record_TextRecord:
		return textRecordFromProto(rec.Id, i.TextRecord)
	case *proto.Record_BinaryRecord:
		return binaryRecordFromProto(rec.Id, i.BinaryRecord)
	case *proto.Record_BankCardRecord:
		return bankCardRecordFromProto(rec.Id, i.BankCardRecord)
	default:
		// Should never happen
		return nil
	}
}
