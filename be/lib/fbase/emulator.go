package fbase

import "os"

type EmulatorConfig struct {
	Enable        bool
	AuthHost      string
	FirestoreHost string
}

func ConfigureEmulator(emulatorConfig EmulatorConfig) error {
	if !emulatorConfig.Enable {
		return nil
	}

	err := os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", emulatorConfig.AuthHost)
	if err != nil {
		return err
	}

	err = os.Setenv("FIRESTORE_EMULATOR_HOST", emulatorConfig.FirestoreHost)
	if err != nil {
		return err
	}

	return nil
}
