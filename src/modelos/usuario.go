package modelos

import "time"

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}
