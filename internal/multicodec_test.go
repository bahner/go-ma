package internal

import (
	"reflect"
	"testing"
)

func TestMulticodecEncodeDecode(t *testing.T) {
	tests := []struct {
		name      string
		codecName string
		payload   []byte
		wantErr   bool
	}{
		{
			name:      "Test with private codec",
			codecName: "ECDHX25519ChaCha20Poly1305BLAKE3",
			payload:   []byte("test payload"),
		},
		{
			name:      "Test with standard codec",
			codecName: "blake3",
			payload:   []byte("test payload"),
		},
		{
			name:      "Test with unknown codec",
			codecName: "unknown",
			payload:   []byte("test payload"),
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := MulticodecEncode(tt.codecName, tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("MulticodecEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return // skip decoding test if encoding failed
			}

			decodedCodecName, decodedPayload, err := MulticodecDecode(encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("MulticodecDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if decodedCodecName != tt.codecName || !reflect.DeepEqual(decodedPayload, tt.payload) {
				t.Errorf("MulticodecDecode() got = %v, %v, want %v, %v", decodedCodecName, decodedPayload, tt.codecName, tt.payload)
			}
		})
	}
}

func TestPrivateCodecMapping(t *testing.T) {
	tests := []struct {
		name      string
		codecName string
		codecCode uint64
		wantErr   bool
	}{
		{
			name:      "Test with valid private codec",
			codecName: "ECDHX25519ChaCha20Poly1305BLAKE3",
			codecCode: 0x300100, // Replace with the actual code
		},
		{
			name:      "Test with invalid private codec",
			codecName: "InvalidCodec",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPrivateCodecValue(tt.codecName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrivateCodecValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.codecCode {
				t.Errorf("GetPrivateCodecValue() got = %v, want %v", got, tt.codecCode)
			}

			if tt.wantErr {
				return // skip name test if value test failed or was expected to fail
			}

			gotName, err := GetPrivateCodecName(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrivateCodecName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotName != tt.codecName {
				t.Errorf("GetPrivateCodecName() got = %v, want %v", gotName, tt.codecName)
			}
		})
	}
}
