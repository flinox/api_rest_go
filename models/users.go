package models

type User struct {
	ID       string  `json:"id,omitempty"`
	Login    string  `json:"login,omitempty"`
	Password string  `json:"password,omitempty"`
	People   *People `json:"people,omitempty"`
}
type People struct {
	ID   string `json:"id,omitempty"`
	Cpf  string `json:"cpf,omitempty"`
	Name string `json:"name,omitempty"`
}
