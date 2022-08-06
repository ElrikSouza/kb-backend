package session

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/vmihailenco/msgpack/v5"
)

type RedisSessionStore struct {
	client          *redis.Client
	tokenGenerator  TokenGenerator
	sessionDuration int
}

func (store *RedisSessionStore) SaveSession(session SessionPayload) (string, error) {
	marshledPayload, err := msgpack.Marshal(session)

	if err != nil {
		return "", err
	}

	sessionToken := store.tokenGenerator.GenerateToken()

	err = store.client.Set(context.Background(), sessionToken, marshledPayload, time.Duration(store.sessionDuration)).Err()

	if err != nil {
		return "", err
	}

	return sessionToken, nil
}

func (store *RedisSessionStore) RetrieveSession(sessionToken string) (SessionPayload, error) {
	marshledPayload, err := store.client.Get(context.Background(), sessionToken).Result()

	if err != nil {
		return SessionPayload{}, err
	}

	var payload SessionPayload

	err = msgpack.Unmarshal([]byte(marshledPayload), &payload)

	if err != nil {
		return SessionPayload{}, err
	}

	return payload, nil
}

func NewRedisSessionStore(options *redis.Options, tokenByteLength, sessionDuration int) *RedisSessionStore {
	client := redis.NewClient(options)
	tokenGenerator := NewTokenGenerator(tokenByteLength)

	return &RedisSessionStore{client, tokenGenerator, sessionDuration}
}
