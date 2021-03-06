package sending_email

import (
	"encoding/json"
	"errors"
)

const msgErrEmpty = "empty content input"

func NewMsg(content string) (*Msg, error) {
	if content == "" {
		return nil, errors.New(msgErrEmpty)
	}

	var c msgContent
	err := json.Unmarshal([]byte(content), &c)
	if err != nil {
		return nil, err
	}

	return &Msg{content, c}, nil
}

type msgContent struct {
	Email string `json:"email"`
}

type Msg struct {
	original string
	content  msgContent
}

func (m *Msg) OriginJson() string {
	return m.original
}
