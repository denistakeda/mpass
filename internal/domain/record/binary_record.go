package record

import (
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Record = (*binaryRecord)(nil)

type binaryRecord struct {
	id             string
	lastUpdateDate time.Time

	binary []byte
}

func binaryRecordFromProto(id string, lastUpdateDate time.Time, p *proto.BinaryRecord) *binaryRecord {
	return &binaryRecord{
		id:             id,
		lastUpdateDate: lastUpdateDate,

		binary: p.Binary,
	}
}

func (r *binaryRecord) GetId() string {
	return r.id
}

func (r *binaryRecord) GetLastUpdateDate() time.Time {
	return r.lastUpdateDate
}

func (r *binaryRecord) ToProto() *proto.Record {
	return &proto.Record{
		Id:             r.id,
		LastUpdateDate: timestamppb.New(r.lastUpdateDate),

		Record: &proto.Record_BinaryRecord{
			BinaryRecord: &proto.BinaryRecord{Binary: r.binary},
		},
	}
}
