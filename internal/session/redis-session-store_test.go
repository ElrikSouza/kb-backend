package session

import (
	"testing"

	"github.com/go-redis/redis/v9"
)

func TestSessionCreation(t *testing.T) {
	redisStore := NewRedisSessionStore(&redis.Options{Addr: "localhost:6969"}, 32, 380000)

	token, err := redisStore.SaveSession(SessionPayload{
		Id:       "an id",
		Username: "Redis",
		Email:    "redis@email.com",
	})

	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error("Invalid token")
	}
}

func TestGetSessionPayload(t *testing.T) {
	redisStore := NewRedisSessionStore(&redis.Options{Addr: "localhost:6969"}, 32, 380000)

	payload := SessionPayload{
		Id:       "an id",
		Username: "Redis",
		Email:    "redis@email.com",
	}

	token, _ := redisStore.SaveSession(payload)

	retrievedPayload, _ := redisStore.RetrieveSession(token)

	if payload.Email != retrievedPayload.Email || payload.Id != retrievedPayload.Id || retrievedPayload.Username != retrievedPayload.Username {
		t.Log(payload)
		t.Log(retrievedPayload)
		t.Error("Payload doesn't match")
	}

}
