import {firebaseAuth} from '../firebase/firebase'
import {useEffect, useState} from "react"
import {EmailAuthProvider, FacebookAuthProvider, GoogleAuthProvider} from "firebase/auth"
import {StyledFirebaseAuth} from "./StyledFirebaseAuth"
import {BackendContextActionKind, useBackend} from "../backend/BackendContextProvider"

const uiConfig = {
    signInFlow: 'popup',
    signInSuccessUrl: '/signedIn',
    signInOptions: [
        EmailAuthProvider.PROVIDER_ID,
        GoogleAuthProvider.PROVIDER_ID,
        FacebookAuthProvider.PROVIDER_ID,
    ],
}

enum SignInState {
    Loading = "LOADING",
    SignedIn = "SIGNED_IN",
    SignedOut = "SIGNED_OUT"
}

export const SignIn = () => {
    const [isSignedIn, setSignedIn] = useState(SignInState.Loading)
    const backend = useBackend()

    const signIn = (token: string) => {
        setSignedIn(SignInState.SignedIn)
        backend.dispatch!({type: BackendContextActionKind.AuthChanged, payload: {jwt: token}})
    }

    const signOut = () => {
        setSignedIn(SignInState.SignedOut)
        backend.dispatch!({type: BackendContextActionKind.AuthChanged, payload: {jwt: null}})
    }

    console.log("headers", backend.headers)

    useEffect(() => {
        const unregisterAuthObserver = firebaseAuth.onAuthStateChanged(authUser => {
            if (!authUser) return signOut()

            authUser.getIdToken()
                .then(signIn)
                .catch(signOut)
        })
        return () => unregisterAuthObserver()
    })

    switch (isSignedIn) {
        case SignInState.Loading:
            return (
                <div>Loading...</div>
            )

        case SignInState.SignedOut:
            return (
                <div>
                    <p>Please sign-in:</p>
                    <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebaseAuth}/>
                </div>
            )

        case SignInState.SignedIn:
            return (
                <div>
                    <p>Welcome!</p>
                    <button onClick={() => firebaseAuth.signOut()}>Sign-out</button>
                </div>
            )
    }
}
