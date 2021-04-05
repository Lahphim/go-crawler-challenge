package v1serializers

import (
	"time"
)

type TokenInformation struct {
	AccessToken  string        `json:"access_token"`
	ExpiresIn    time.Duration `json:"expires_in"`
	RefreshToken string        `json:"refresh_token"`
	TokenType    string        `json:"token_type"`
}

type TokenError struct {
	Code   string `json:"error"`
	Detail string `json:"error_description"`
}

type TokenInformationResponse struct {
	Id           int           `jsonapi:"primary,token_information"`
	AccessToken  string        `jsonapi:"attr,access_token"`
	ExpiresIn    time.Duration `jsonapi:"attr,expires_in"`
	RefreshToken string        `jsonapi:"attr,refresh_token"`
	TokenType    string        `jsonapi:"attr,token_type"`
}

func (serializer *TokenInformation) Data() (data *TokenInformationResponse) {
	data = &TokenInformationResponse{
		AccessToken:  serializer.AccessToken,
		TokenType:    serializer.TokenType,
		RefreshToken: serializer.RefreshToken,
		ExpiresIn:    serializer.ExpiresIn,
	}

	return data
}
