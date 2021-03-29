package v1serializers

import (
	"time"
)

const TokenType = "Bearer"

type TokenInformation struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Duration
}

type tokenInformationResponse struct {
	Id           int           `jsonapi:"primary,token_information"`
	AccessToken  string        `jsonapi:"attr,access_token"`
	TokenType    string        `jsonapi:"attr,token_type"`
	RefreshToken string        `jsonapi:"attr,refresh_token"`
	Expiry       time.Duration `jsonapi:"attr,expiry"`
}

func (serializer *TokenInformation) Data() (data *tokenInformationResponse) {
	data = &tokenInformationResponse{
		AccessToken:  serializer.AccessToken,
		TokenType:    TokenType,
		RefreshToken: serializer.RefreshToken,
		Expiry:       serializer.Expiry,
	}

	return data
}
