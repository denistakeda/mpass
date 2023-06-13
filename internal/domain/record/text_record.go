package record

import (
	"encoding/gob"
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Record = (*TextRecord)(nil)

func init() {
	gob.Register(&TextRecord{})
}

type TextRecord struct {
	ID             string    `db:"id"`
	LastUpdateDate time.Time `db:"last_update_date"`

	Text string `db:"text"`
}

func NewTextRecord(key string, text string) *TextRecord {
	return &TextRecord{
		ID:             key,
		LastUpdateDate: time.Now(),

		Text: text,
	}
}

func textRecordFromProto(id string, lastUpdateDate time.Time, p *proto.TextRecord) *TextRecord {
	return &TextRecord{
		ID:             id,
		LastUpdateDate: lastUpdateDate,

		Text: p.Text,
	}
}

func (r *TextRecord) GetId() string {
	return r.ID
}

func (r *TextRecord) GetLastUpdateDate() time.Time {
	return r.LastUpdateDate
}

func (r *TextRecord) ToProto() *proto.Record {
	return &proto.Record{
		Id:             r.ID,
		LastUpdateDate: timestamppb.New(r.LastUpdateDate),

		Record: &proto.Record_TextRecord{
			TextRecord: &proto.TextRecord{Text: r.Text},
		},
	}
}

func (r *TextRecord) ProvideToClient(printer printer) error {
	printer.Printf("Text:\n %s", r.Text)
	return nil
}
