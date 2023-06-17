package record

import (
	"encoding/gob"
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Record = (*LoginPasswordRecord)(nil)

func init() {
	gob.Register(&BinaryRecord{})
}

func init() {
	gob.Register(&LoginPasswordRecord{})
}

type LoginPasswordRecord struct {
	ID             string    `db:"id"`
	LastUpdateDate time.Time `db:"last_update_date"`

	Login    string `db:"login"`
	Password string `db:"password"`
}

func NewLoginPasswordRecord(login, password string) *LoginPasswordRecord {
	return &LoginPasswordRecord{
		ID:             login,
		LastUpdateDate: time.Now(),

		Login:    login,
		Password: password,
	}
}

func loginPasswordRecordFromProto(id string, lastUpdateDate time.Time, p *proto.LoginPasswordRecord) *LoginPasswordRecord {
	return &LoginPasswordRecord{
		ID:             id,
		LastUpdateDate: lastUpdateDate,

		Login:    p.Login,
		Password: p.Password,
	}
}

func (r *LoginPasswordRecord) GetId() string {
	return r.ID
}

func (r *LoginPasswordRecord) GetLastUpdateDate() time.Time {
	return r.LastUpdateDate
}

func (r *LoginPasswordRecord) ToProto() *proto.Record {
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
func (r *LoginPasswordRecord) ProvideToClient(printer printer) error {
	printer.Printf("Password: %s\n", r.Password)
	return nil
}
