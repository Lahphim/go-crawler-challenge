package initializers

import (
	_ "go-crawler-challenge/routers"
)

func init() {
	SetUpDatabase()
}
