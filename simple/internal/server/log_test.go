package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLog(t *testing.T) {
	log := NewLog()
	assert.NotNil(t, log)
}

func TestAppend(t *testing.T) {
	log := NewLog()
	record := Record{
		Value: []byte("hello"),
	}
	offset, err := log.Append(record)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), offset)
}

func TestRead(t *testing.T) {
	log := NewLog()
	record := Record{
		Value: []byte("hello"),
	}
	offset, err := log.Append(record)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), offset)

	readRecord, err := log.Read(offset)
	assert.NoError(t, err)
	assert.Equal(t, record, readRecord)
}

func TestReadErrOutofRange(t *testing.T) {
	log := NewLog()
	record := Record{
		Value: []byte("hello"),
	}
	offset, err := log.Append(record)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), offset)

	_, err = log.Read(offset + 1)
	assert.Error(t, err)
}
