package support

import (
	"git.jetbrains.space/artdecoction/wt/tower/lib/fbase"
	"github.com/spf13/viper"
)

func createFirebaseClients() *fbase.Clients {
	cfg := fbase.EmulatorConfig{
		Enable:        true,
		AuthHost:      viper.GetString("firebaseEmulator.authHost"),
		FirestoreHost: viper.GetString("firebaseEmulator.firestoreHost"),
	}

	err := fbase.ConfigureEmulator(cfg)
	if err != nil {
		panic(err)
	}

	firebaseClients, err := fbase.NewClients(viper.GetString("gcpProjectId"))
	if err != nil {
		panic(err)
	}

	return firebaseClients
}
