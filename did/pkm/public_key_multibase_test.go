package pkm

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"reflect"
	"testing"
)

func TestNewAndDecode(t *testing.T) {
	tests := []struct {
		name string
		key  interface{}
	}{
		{
			name: "RSA Private Key",
			key:  generateRSAPrivateKey(t),
		},
		{
			name: "RSA Public Key",
			key:  generateRSAPrivateKey(t).Public(),
		},
		{
			name: "Ed25519 Private Key",
			key:  generateEd25519PrivateKey(t),
		},
		{
			name: "Ed25519 Public Key",
			key:  generateEd25519PrivateKey(t).Public().(ed25519.PublicKey),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkmb, err := New(tt.key)
			if err != nil {
				t.Fatalf("failed to create PublicKeyMultibase: %v", err)
			}

			decodedPKMB, err := Parse(pkmb.PublicKeyMultibase)
			if err != nil {
				t.Fatalf("failed to decode PublicKeyMultibase: %v", err)
			}

			if pkmb.PublicKeyMultibase != decodedPKMB.PublicKeyMultibase {
				t.Errorf("PublicKeyMultibase mismatch: got %v, want %v", decodedPKMB.PublicKeyMultibase, pkmb.PublicKeyMultibase)
			}

			if pkmb.MulticodecCodeString != decodedPKMB.MulticodecCodeString {
				t.Errorf("MulticodecCodeString mismatch: got %v, want %v", decodedPKMB.MulticodecCodeString, pkmb.MulticodecCodeString)
			}

			if !reflect.DeepEqual(pkmb.PublicKey, decodedPKMB.PublicKey) {
				t.Errorf("PublicKey mismatch: got %v, want %v", decodedPKMB.PublicKey, pkmb.PublicKey)
			}
		})
	}
}

func generateRSAPrivateKey(t *testing.T) *rsa.PrivateKey {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate RSA private key: %v", err)
	}
	return key
}

func generateEd25519PrivateKey(t *testing.T) ed25519.PrivateKey {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate Ed25519 private key: %v", err)
	}
	return priv
}
