package helpers

import (
	"encoding/base64"
	"github.com/satori/uuid"
	"strings"
)

var escaper = strings.NewReplacer("9", "99", "-", "90", "_", "91")
var unescaper = strings.NewReplacer("99", "9", "90", "-", "91", "_")

func UuidToShortString(u uuid.UUID) string {
	return escaper.Replace(base64.RawURLEncoding.EncodeToString(u.Bytes()))
}

func ShortStringToUuid(s string) (uuid.UUID, error) {
	dec, err := base64.RawURLEncoding.DecodeString(unescaper.Replace(s))
	if err != nil {
		return uuid.UUID{}, err
	}

	u, err := uuid.FromBytes(dec)
	if err != nil {
		return uuid.UUID{}, err
	}

	return u, nil
}
