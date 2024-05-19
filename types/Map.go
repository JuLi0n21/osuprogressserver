package types

import (
	"errors"
	"sync"
)

type Mapstruct struct {
	sync.Mutex
	m map[string]UserContext
}

func (m *Mapstruct) Read(key string) (UserContext, error) {
	m.Lock()
	defer m.Unlock()
	if key == "" {
		return UserContext{}, errors.New("invalid key provied")
	}

	if v, ok := m.m[key]; ok {
		return v, nil
	}

	return UserContext{}, nil
}

func (m *Mapstruct) Write(key string, value UserContext) error {
	m.Lock()
	defer m.Unlock()
	if key == "" {
		return errors.New("invalid key provided")
	}

	m.m[key] = value
	return nil
}

func NewUserSessionMap() *Mapstruct {
	return &Mapstruct{
		m: make(map[string]UserContext),
	}
}
