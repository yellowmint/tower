import {firebaseAuth} from '../firebase/firebase'
import {useEffect} from "react"
import {EmailAuthProvider, FacebookAuthProvider, GoogleAuthProvider} from "firebase/auth"
import {StyledFirebaseAuth} from "./StyledFirebaseAuth"
import {AuthContextActionKind, useAuth} from "./AuthContextProvider"

const uiConfig = {
    signInFlow: 'popup',
    signInSuccessUrl: '/signedIn',
    signInOptions: [
        EmailAuthProvider.PROVIDER_ID,
        GoogleAuthProvider.PROVIDER_ID,
        FacebookAuthProvider.PROVIDER_ID,
    ],
}

export const SignIn = () => {
    const auth = useAuth()

    useEffect(() => {
        const unregisterAuthObserver = firebaseAuth.onAuthStateChanged(user => {
            if (user) {
                auth.dispatch!({type: AuthContextActionKind.SignedIn, payload: {user: user}})
            } else {
                auth.dispatch!({type: AuthContextActionKind.SignedOut})
            }
        })
        return () => unregisterAuthObserver()
    }, [auth.dispatch])

    if (!auth.user) {
        return (
            <div>
                <p>Please sign-in:</p>
                <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebaseAuth}/>
            </div>
        )
    }

    return (
        <div>
            <p>Welcome {auth.user.displayName}!</p>
            <button onClick={() => firebaseAuth.signOut()}>Sign-out</button>
        </div>
    )
}
