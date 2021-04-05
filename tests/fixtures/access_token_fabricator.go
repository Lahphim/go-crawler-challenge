package fixtures

import (
	"context"
	"fmt"
	"time"

	. "go-crawler-challenge/services/oauth"

	"github.com/bxcodec/faker/v3"
	"github.com/go-oauth2/oauth2/v4/models"
	. "github.com/onsi/ginkgo"
)

func FabricatorAccessToken(clientId string, userId int64) (accessToken string) {
	tokenInfo := &models.Token{
		ClientID:         clientId,
		UserID:           fmt.Sprint(userId),
		Access:           faker.Password(),
		AccessCreateAt:   time.Now().Local(),
		AccessExpiresIn:  time.Hour * 2,
		Refresh:          faker.Password(),
		RefreshCreateAt:  time.Now().Local(),
		RefreshExpiresIn: time.Hour * 2,
	}

	err := TokenStore.Create(context.Background(), tokenInfo)
	if err != nil {
		Fail("Add TokenInfo failed: " + err.Error())
	}

	return tokenInfo.GetAccess()
}
