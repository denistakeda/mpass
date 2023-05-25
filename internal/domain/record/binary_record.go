package record

import (
	"github.com/denistakeda/mpass/internal/ports"
	"github.com/denistakeda/mpass/proto"
)

var _ ports.Record = (*binaryRecord)(nil)

type binaryRecord struct {
	id string

	binary []byte
}

func binaryRecordFromProto(id string, p *proto.BinaryRecord) *binaryRecord {
	return &binaryRecord{
		id:     id,
		binary: p.Binary,
	}
}