package api

import (
	"time"

	"github.com/iaping/exgo"
)

type Signature struct {
	Key, Timestamp, Method, Path, Body string
}

func NewSignature(key, method, path, body string) *Signature {
	return &Signature{
		Key:    key,
		Method: method,
		Path:   path,
		Body:   body,
	}
}

// The Base64-encoded signature (see Signing Messages subsection for details).
func (s *Signature) Build() string {
	if s.Timestamp == "" {
		s.setTimestamp()
	}
	data := s.Timestamp + s.Method + s.Path + s.Body
	return exgo.HmacSHA256([]byte(data), []byte(s.Key))
}

// The timestamp of your request.e.g : 2020-12-08T09:08:57.715Z
func (s *Signature) setTimestamp() {
	s.Timestamp = time.Now().UTC().Format(time.RFC3339)
}
