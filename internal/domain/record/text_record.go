package record

import (
	"encoding/gob"
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Record = (*textRecord)(nil)

func init() {
	gob.Register(&textRecord{})
}

type textRecord struct {
	ID             string    `db:"id"`
	LastUpdateDate time.Time `db:"last_update_date"`

	Text string `db:"text"`
}

func NewTextRecord(key string, text string) *textRecord {
	return &textRecord{
		ID:             key,
		LastUpdateDate: time.Now(),

		Text: text,
	}
}

func textRecordFromProto(id string, lastUpdateDate time.Time, p *proto.TextRecord) *textRecord {
	return &textRecord{
		ID:             id,
		LastUpdateDate: lastUpdateDate,

		Text: p.Text,
	}
}

func (r *textRecord) GetId() string {
	return r.ID
}

func (r *textRecord) GetLastUpdateDate() time.Time {
	return r.LastUpdateDate
}

func (r *textRecord) ToProto() *proto.Record {
	return &proto.Record{
		Id:             r.ID,
		LastUpdateDate: timestamppb.New(r.LastUpdateDate),

		Record: &proto.Record_TextRecord{
			TextRecord: &proto.TextRecord{Text: r.Text},
		},
	}
}

func (r *textRecord) ProvideToClient(printer printer) error {
	printer.Printf("Text:\n %s", r.Text)
	return nil
}
