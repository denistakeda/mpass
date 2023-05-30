package record

import (
	"time"

	"github.com/denistakeda/mpass/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Record = (*loginPasswordRecord)(nil)

type loginPasswordRecord struct {
	id             string
	lastUpdateDate time.Time

	login    string
	password string
}

func NewLoginPasswordRecord(login, password string) *loginPasswordRecord {
	return &loginPasswordRecord{
		id:             login,
		lastUpdateDate: time.Now(),

		login:    login,
		password: password,
	}
}

func loginPasswordRecordFromProto(id string, lastUpdateDate time.Time, p *proto.LoginPasswordRecord) *loginPasswordRecord {
	return &loginPasswordRecord{
		id:             id,
		lastUpdateDate: lastUpdateDate,

		login:    p.Login,
		password: p.Password,
	}
}

func (r *loginPasswordRecord) GetId() string {
	return r.id
}

func (r *loginPasswordRecord) GetLastUpdateDate() time.Time {
	return r.lastUpdateDate
}

func (r *loginPasswordRecord) ToProto() *proto.Record {
	return &proto.Record{
		Id:             r.id,
		LastUpdateDate: timestamppb.New(r.lastUpdateDate),

		Record: &proto.Record_LoginPasswordRecord{
			LoginPasswordRecord: &proto.LoginPasswordRecord{
				Login:    r.login,
				Password: r.password,
			},
		},
	}
}
