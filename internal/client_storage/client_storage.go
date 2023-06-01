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
		records map[string]record.Record
		toSync  map[string]record.Record
	}
)

func New(filepath string) *clientStorage {
	return &clientStorage{filepath: filepath}
}

func (c *clientStorage) SetRecord(r record.Record) error {
	c.mx.Lock()
	defer c.mx.Unlock()

	state, err := c.getState()
	if err != nil {
		return err
	}

	state.records[r.GetId()] = r
	state.toSync[r.GetId()] = r

	return nil
}

func (c *clientStorage) Close() error {
	file, err := os.Create(c.filepath)
	if err != nil {
		return errors.Wrapf(err, "failed to open file %q for writing", c.filepath)
	}
	defer file.Close()

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
		records: make(map[string]record.Record),
		toSync:  make(map[string]record.Record),
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
