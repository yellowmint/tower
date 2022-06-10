import {connectAuthEmulator, getAuth} from "firebase/auth"
import {FirebaseApp, FirebaseOptions, initializeApp} from 'firebase/app'

const firebaseConfig: FirebaseOptions = {
    apiKey: process.env.REACT_APP_FIREBASE_API_KEY,
    authDomain: process.env.REACT_APP_FIREBASE_AUTH_DOMAIN,
    projectId: process.env.REACT_APP_FIREBASE_PROJECT_ID,
    storageBucket: process.env.REACT_APP_FIREBASE_STORAGE_BUCKET,
    messagingSenderId: process.env.REACT_APP_FIREBASE_MESSAGING_SENDER_ID,
    appId: process.env.REACT_APP_FIREBASE_APP_ID
}

const firebaseApp = initializeApp(firebaseConfig)
connectEmulators(firebaseApp)

export const firebaseAuth = getAuth(firebaseApp)


function connectEmulators(app: FirebaseApp) {
    if (process.env.REACT_APP_FIREBASE_EMULATOR_DISABLE === "1") return

    connectAuthEmulator(getAuth(app), process.env.REACT_APP_FIREBASE_AUTH_EMULATOR_URL!)
}
