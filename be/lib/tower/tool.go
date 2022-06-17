package tower

import "context"

type AppTool interface {
	GetAuth() Auth
}

type Auth interface {
	SetCustomUserClaims(ctx context.Context, uid string, customClaims map[string]interface{}) error
	DeleteUser(ctx context.Context, uid string) error
}

func (a *App) GetAuth() Auth {
	return a.FirebaseClients.Auth
}
