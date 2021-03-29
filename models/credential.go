package models

type Credential struct {
	ClientId     string `jsonapi:"attr,client_id"`
	ClientSecret string `jsonapi:"attr,client_secret"`
	GrantType    string `jsonapi:"attr,grant_type"`
	Email        string `jsonapi:"attr,email"`
	Password     string `jsonapi:"attr,password"`
}
