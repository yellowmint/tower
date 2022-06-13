import {firebaseAuth} from "../firebase/firebase"
import {useEffect, useState} from "react"
import {EmailAuthProvider, FacebookAuthProvider, GoogleAuthProvider, IdTokenResult} from "firebase/auth"
import {StyledFirebaseAuth} from "./StyledFirebaseAuth"
import {BackendContextActions, useBackend} from "../backend/BackendContextProvider"
import {Registration, RegistrationProps} from "./Registration"

const uiConfig = {
    signInFlow: 'popup',
    signInSuccessUrl: '/signedIn',
    signInOptions: [
        GoogleAuthProvider.PROVIDER_ID,
        FacebookAuthProvider.PROVIDER_ID,
        {
            provider: EmailAuthProvider.PROVIDER_ID,
            requireDisplayName: false,
        },
    ],
}

enum authStatuses {
    Loading = "LOADING",
    SignedIn = "SIGNED_IN",
    SignedOut = "SIGNED_OUT",
    Registration = "REGISTRATION",
}

export const SignIn = () => {
    const backend = useBackend()
    const [authStatus, setAuthStatus] = useState<authStatuses>(authStatuses.Loading)
    const [registrationData, setRegistrationData] = useState<RegistrationProps | null>(null)

    const signIn = (tokenResult: IdTokenResult) => {
        if (tokenResult.claims["accountId"]) {
            setAuthStatus(authStatuses.SignedIn)
            backend.dispatch!({type: BackendContextActions.AuthChanged, payload: {jwt: tokenResult.token}})
            return
        }

        setAuthStatus(authStatuses.Registration)
        setRegistrationData({
            token: tokenResult.token,
            initName: firebaseAuth.currentUser!.displayName,
            successCallback: () => {
                firebaseAuth.currentUser!.getIdTokenResult(true).then(signIn).catch(signOut)
            }
        })
    }

    const signOut = () => {
        setAuthStatus(authStatuses.SignedOut)
        backend.dispatch!({type: BackendContextActions.AuthChanged, payload: {jwt: null}})
    }

    useEffect(() => {
        const unregisterAuthObserver = firebaseAuth.onAuthStateChanged(authUser => {
            if (!authUser) return signOut()

            authUser.getIdTokenResult().then(signIn).catch(signOut)
        })
        return () => unregisterAuthObserver()
    })

    switch (authStatus) {
        case authStatuses.Loading:
            return (
                <div>Loading...</div>
            )

        case authStatuses.SignedOut:
            return (
                <div>
                    <p>Please sign-in:</p>
                    <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebaseAuth}/>
                </div>
            )

        case authStatuses.SignedIn:
            return (
                <div>
                    <p>Welcome!</p>
                    <button onClick={() => firebaseAuth.signOut()}>Sign-out</button>
                </div>
            )

        case authStatuses.Registration:
            return <Registration {...registrationData!} />
    }
}
