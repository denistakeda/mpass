package record

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Record = (*BinaryRecord)(nil)

func init() {
	gob.Register(&BinaryRecord{})
}

type BinaryRecord struct {
	ID             string    `db:"id"`
	LastUpdateDate time.Time `db:"last_update_date"`

	Binary []byte `db:"binary"`
}

func NewBinaryRecord(key string, data []byte) *BinaryRecord {
	return &BinaryRecord{
		ID:             key,
		LastUpdateDate: time.Now(),

		Binary: data,
	}
}

func binaryRecordFromProto(id string, lastUpdateDate time.Time, p *proto.BinaryRecord) *BinaryRecord {
	return &BinaryRecord{
		ID:             id,
		LastUpdateDate: lastUpdateDate,

		Binary: p.Binary,
	}
}

func (r *BinaryRecord) GetId() string {
	return r.ID
}

func (r *BinaryRecord) GetLastUpdateDate() time.Time {
	return r.LastUpdateDate
}

func (r *BinaryRecord) ToProto() *proto.Record {
	return &proto.Record{
		Id:             r.ID,
		LastUpdateDate: timestamppb.New(r.LastUpdateDate),

		Record: &proto.Record_BinaryRecord{
			BinaryRecord: &proto.BinaryRecord{Binary: r.Binary},
		},
	}
}

func (r *BinaryRecord) ProvideToClient(printer printer) error {
	return os.WriteFile(fmt.Sprintf("./%s", r), r.Binary, 760)
}
