package towermock

import (
	"context"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower"
)

type AppToolMock struct {
	Auth AuthMock
}

type AuthMock struct{}

func NewTowerMockApp() tower.AppTool {
	return &AppToolMock{}
}

func (a *AppToolMock) GetAuth() tower.Auth {
	return a.Auth
}

func (a AuthMock) SetCustomUserClaims(_ context.Context, _ string, _ map[string]interface{}) error {
	return nil
}

func (a AuthMock) DeleteUser(_ context.Context, _ string) error {
	return nil
}
