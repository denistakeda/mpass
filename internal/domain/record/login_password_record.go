package record

import (
	"encoding/gob"
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Record = (*loginPasswordRecord)(nil)

func init() {
	gob.Register(&binaryRecord{})
}

func init() {
	gob.Register(&loginPasswordRecord{})
}

type loginPasswordRecord struct {
	ID             string    `db:"id"`
	LastUpdateDate time.Time `db:"last_update_date"`

	Login    string `db:"login"`
	Password string `db:"password"`
}

func NewLoginPasswordRecord(login, password string) *loginPasswordRecord {
	return &loginPasswordRecord{
		ID:             login,
		LastUpdateDate: time.Now(),

		Login:    login,
		Password: password,
	}
}

func loginPasswordRecordFromProto(id string, lastUpdateDate time.Time, p *proto.LoginPasswordRecord) *loginPasswordRecord {
	return &loginPasswordRecord{
		ID:             id,
		LastUpdateDate: lastUpdateDate,

		Login:    p.Login,
		Password: p.Password,
	}
}

func (r *loginPasswordRecord) GetId() string {
	return r.ID
}

func (r *loginPasswordRecord) GetLastUpdateDate() time.Time {
	return r.LastUpdateDate
}

func (r *loginPasswordRecord) ToProto() *proto.Record {
	return &proto.Record{
		Id:             r.ID,
		LastUpdateDate: timestamppb.New(r.LastUpdateDate),

		Record: &proto.Record_LoginPasswordRecord{
			LoginPasswordRecord: &proto.LoginPasswordRecord{
				Login:    r.Login,
				Password: r.Password,
			},
		},
	}
}

// ProvideToClient implements Record
func (r *loginPasswordRecord) ProvideToClient(printer printer) error {
	printer.Printf("Password: %s\n", r.Password)
	return nil
}
