package record

import (
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (r *textRecord) ToProto() *proto.Record {
	return &proto.Record{
		Id:             r.id,
		LastUpdateDate: timestamppb.New(r.lastUpdateDate),

		Record: &proto.Record_TextRecord{
			TextRecord: &proto.TextRecord{Text: r.text},
		},
	}
}
