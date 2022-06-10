package support

import (
	"git.jetbrains.space/artdecoction/wt/tower/lib/fbase"
)

type Support struct {
	FirebaseClients *fbase.Clients
}

func Init() *Support {
	loadConfig()

	return &Support{
		FirebaseClients: createFirebaseClients(),
	}
}

func (s *Support) Cleanup() {
	if s.FirebaseClients != nil {
		_ = fbase.Cleanup(s.FirebaseClients)
	}

	closeGrpcClientConnections()
}
