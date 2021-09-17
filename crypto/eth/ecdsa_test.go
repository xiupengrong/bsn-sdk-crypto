package eth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestKey(t *testing.T) {

	key := "-----BEGIN PRIVATE KEY-----\r\nMIGNAgEAMBAGByqGSM49AgEGBSuBBAAKBHYwdAIBAQQgVxFlqM9/e2SnAOTcBvix\r\nPv+ZG21ku/M5crUavirBLJygBwYFK4EEAAqhRANCAAS6BLPEyNPXvsgE3Z6AC6Qd\r\nb3o3PH9NJHZRlA2m/XFVe0MQ9lY8ZnzpFBVUfQMdi8TZg73I3bCTHGd38++MCLg/\r\n-----END PRIVATE KEY-----"

	pk, err := LoadPrivateKey([]byte(key))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(pk.D.String())

	data := "123456"

	r, s, v, sig, err := SignData(pk, Hash1([]byte(data)))

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r.String())
	fmt.Println(s.String())
	fmt.Println(v.String())
	fmt.Println(hex.EncodeToString(sig))
}

func Hash1(msg []byte) []byte {

	h := sha256.New()

	h.Write([]byte(msg))
	hash := h.Sum(nil)

	return hash
}
