package server

import (
	"errors"
	"sync"
)

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrOffsetNotFound = errors.New("offset not found")

type Log struct {
	Mx      sync.Mutex
	Records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (l *Log) Append(record Record) (uint64, error) {
	l.Mx.Lock()
	defer l.Mx.Unlock()
	record.Offset = uint64(len(record.Value))
	l.Records = append(l.Records, record)
	return record.Offset, nil
}

func (l *Log) Read(offset uint64) (*Record, error) {
	l.Mx.Lock()
	defer l.Mx.Unlock()
	if offset > uint64(len(l.Records)) {
		return &Record{}, ErrOffsetNotFound
	}
	return &l.Records[offset], nil
}
