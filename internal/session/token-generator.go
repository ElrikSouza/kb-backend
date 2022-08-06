package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

type TokenGenerator struct {
	tokenBytes int
}

func (generator TokenGenerator) GenerateToken() string {
	bytes := make([]byte, generator.tokenBytes)

	io.ReadFull(rand.Reader, bytes)

	encoded_bytes := base64.StdEncoding.EncodeToString(bytes)

	return encoded_bytes
}

func NewTokenGenerator(tokenByteLength int) TokenGenerator {
	return TokenGenerator{tokenBytes: tokenByteLength}
}
