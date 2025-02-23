package handlers

import (
	"errors"
	"fmt"
	"time"
)

type MockRedis struct {
	data map[string]string
}

func NewMockRedis() *MockRedis {
	return &MockRedis{data: map[string]string{}}
}

func (m *MockRedis) Get(key string) (string, error) {
	value, exist := m.data[key]
	if !exist {
		return "", errors.New("data doesn't exist")
	}
	return value, nil

}

func (m *MockRedis) Set(key string, value interface{}, expiration time.Duration) error {
	fmt.Println(value)
	m.data[key] = string(value.([]byte))
	fmt.Println(value)
	return nil
}
