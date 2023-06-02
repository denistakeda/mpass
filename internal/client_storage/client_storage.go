package client_storage

import (
	"encoding/gob"
	"os"
	"sync"

	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/pkg/errors"
)

type (
	clientStorage struct {
		filepath string
		mx       sync.Mutex
		state    *state
	}

	state struct {
		Token string

		Records map[string]record.Record
		ToSync  map[string]record.Record
	}
)

func New(filepath string) *clientStorage {
	return &clientStorage{filepath: filepath}
}

func (c *clientStorage) GetToken() (string, error) {
	c.mx.Lock()
	defer c.mx.Unlock()

	state, err := c.getState()
	if err != nil {
		return "", err
	}

	return state.Token, nil
}

func (c *clientStorage) SetToken(t string) error {
	c.mx.Lock()
	defer c.mx.Unlock()

	state, err := c.getState()
	if err != nil {
		return err
	}

	state.Token = t

	return nil
}

func (c *clientStorage) SetRecord(r record.Record) error {
	c.mx.Lock()
	defer c.mx.Unlock()

	state, err := c.getState()
	if err != nil {
		return err
	}

	state.Records[r.GetId()] = r
	state.ToSync[r.GetId()] = r

	return nil
}

func (c *clientStorage) GetRecord(key string) (record.Record, error) {
	c.mx.Lock()
	defer c.mx.Unlock()

	state, err := c.getState()
	if err != nil {
		return nil, err
	}

	rec, ok := state.Records[key]
	if !ok {
		return nil, errors.Errorf("no record with key %q", key)
	}

	return rec, nil
}

func (c *clientStorage) ItemsToSync() ([]record.Record, error) {
	c.mx.Lock()
	defer c.mx.Unlock()

	state, err := c.getState()
	if err != nil {
		return nil, err
	}

	res := make([]record.Record, 0, len(state.ToSync))
	for _, item := range state.ToSync {
		res = append(res, item)
	}

	return res, nil
}

func (c *clientStorage) SyncRecords(records []record.Record) error {
	c.mx.Lock()
	defer c.mx.Unlock()

	state, err := c.getState()
	if err != nil {
		return err
	}

	state.ToSync = make(map[string]record.Record)
	state.Records = make(map[string]record.Record)

	for _, item := range records {
		state.Records[item.GetId()] = item
	}

	return nil
}

func (c *clientStorage) Close() error {
	// only store the state if it was loaded before
	if c.state == nil {
		return nil
	}

	file, err := os.Create(c.filepath)
	if err != nil {
		return errors.Wrapf(err, "failed to open file %q for writing", c.filepath)
	}
	defer file.Close()

	gob.Register(map[string]record.Record{})
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(c.state)
	if err != nil {
		return errors.Wrap(err, "failed to encode the state")
	}

	return nil
}

func (c *clientStorage) getState() (*state, error) {
	if c.state == nil {
		err := c.loadStateFromFile()
		return c.state, err
	}

	return c.state, nil
}

func (c *clientStorage) loadStateFromFile() error {
	c.state = &state{
		Records: make(map[string]record.Record),
		ToSync:  make(map[string]record.Record),
	}

	file, err := os.Open(c.filepath)
	if err != nil {
		// the file does not exists, returns the default state
		return nil
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(c.state)
	if err != nil {
		return errors.Wrapf(err, "failed to decode the content of file %q", c.filepath)
	}

	return nil
}
