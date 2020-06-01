package jwks

import (
	"gopkg.in/square/go-jose.v2"
)

type DummySource struct {
	Jwks *jose.JSONWebKeySet
}

func NewDummySource(jwks *jose.JSONWebKeySet) *DummySource {
	return &DummySource{Jwks: jwks}
}

func (s *DummySource) JSONWebKeySet() (*jose.JSONWebKeySet, error) {
	return s.Jwks, nil
}
