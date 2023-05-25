package record

import (
	"github.com/denistakeda/mpass/internal/ports"
	"github.com/denistakeda/mpass/proto"
)

var _ ports.Record = (*loginPasswordRecord)(nil)

type loginPasswordRecord struct {
	id string

	login    string
	password string
}

func loginPasswordRecordFromProto(id string, p *proto.LoginPasswordRecord) *loginPasswordRecord {
	return &loginPasswordRecord{
		id:       id,
		login:    p.Login,
		password: p.Password,
	}
}
