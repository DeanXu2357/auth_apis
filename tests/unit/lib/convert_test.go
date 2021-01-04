package lib

import (
	"auth/internal/helpers"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_UuidToString(t *testing.T) {
	uuidBytes := []byte{0x12, 0x3e, 0x45, 0x67, 0xe8, 0x9b, 0x12, 0xd3, 0xa4, 0x56, 0x42, 0x66, 0x55, 0x44, 0x00, 0x00}
	shortString := "Ej5FZ90ibEtOkVkJmVUQAAA"
	uuidObject, _ := uuid.FromBytes(uuidBytes)

	s := helpers.UuidToShortString(uuidObject)

	assert.Equal(t, s, shortString)
}

func Test_StringToUuid(t *testing.T) {
	uuidBytes := []byte{0x12, 0x3e, 0x45, 0x67, 0xe8, 0x9b, 0x12, 0xd3, 0xa4, 0x56, 0x42, 0x66, 0x55, 0x44, 0x00, 0x00}
	shortString := "Ej5FZ90ibEtOkVkJmVUQAAA"
	uuidObject, _ := uuid.FromBytes(uuidBytes)

	u, err := helpers.ShortStringToUuid(shortString)

	assert.Nil(t, err)
	assert.Equal(t, u, uuidObject)
}
