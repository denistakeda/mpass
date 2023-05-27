package record

import (
	"github.com/denistakeda/mpass/proto"
	"time"
)

var _ Record = (*textRecord)(nil)

type textRecord struct {
	id             string
	lastUpdateDate time.Time

	text string
}

func textRecordFromProto(id string, lastUpdateDate time.Time, p *proto.TextRecord) *textRecord {
	return &textRecord{
		id:             id,
		lastUpdateDate: lastUpdateDate,

		text: p.Text,
	}
}

func (r *textRecord) GetId() string {
	return r.id
}

func (r *textRecord) GetLastUpdateDate() time.Time {
	return r.lastUpdateDate
}
