package record

import (
	"encoding/gob"
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Record = (*binaryRecord)(nil)

func init() {
	gob.Register(&binaryRecord{})
}

type binaryRecord struct {
	ID             string    `db:"id"`
	LastUpdateDate time.Time `db:"last_update_date"`

	Binary []byte `db:"binary"`
}

func NewBinaryRecord(key string, data []byte) *binaryRecord {
	return &binaryRecord{
		ID:             key,
		LastUpdateDate: time.Now(),

		Binary: data,
	}
}

func binaryRecordFromProto(id string, lastUpdateDate time.Time, p *proto.BinaryRecord) *binaryRecord {
	return &binaryRecord{
		ID:             id,
		LastUpdateDate: lastUpdateDate,

		Binary: p.Binary,
	}
}

func (r *binaryRecord) GetId() string {
	return r.ID
}

func (r *binaryRecord) GetLastUpdateDate() time.Time {
	return r.LastUpdateDate
}

func (r *binaryRecord) ToProto() *proto.Record {
	return &proto.Record{
		Id:             r.ID,
		LastUpdateDate: timestamppb.New(r.LastUpdateDate),

		Record: &proto.Record_BinaryRecord{
			BinaryRecord: &proto.BinaryRecord{Binary: r.Binary},
		},
	}
}

// ProvideToClient implements Record
func (*binaryRecord) ProvideToClient(printer printer) error {
	panic("unimplemented")
}
