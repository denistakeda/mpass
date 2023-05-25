package record

import (
	"github.com/denistakeda/mpass/internal/ports"
	"github.com/denistakeda/mpass/proto"
)

var _ ports.Record = (*textRecord)(nil)

type textRecord struct {
	id string

	text string
}

func textRecordFromProto(id string, p *proto.TextRecord) *textRecord {
	return &textRecord{
		id:   id,
		text: p.Text,
	}
}