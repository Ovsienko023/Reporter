package core

import (
	"time"
)

type AuthState struct {
	ttl          time.Time
	ServerHost   string
	ClientOrigin string
}

type (
	Cache struct {
		// Хранит поле state в oauth2 авторизации
		AuthState map[string]AuthState
	}
)

func newCache() *Cache {
	return &Cache{
		AuthState: make(map[string]AuthState),
	}
}

func (c *Cache) Clean() {
	for range time.Tick(1 * time.Second) {
		for key, state := range c.AuthState {
			if state.ttl.Add(time.Second*5).Unix() < time.Now().Unix() {
				delete(c.AuthState, key)
			}
		}
	}
}
