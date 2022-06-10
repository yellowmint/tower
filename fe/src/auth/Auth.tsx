import {firebaseAuth} from '../firebase/firebase'
import React, {useEffect, useState} from "react"
import {EmailAuthProvider, FacebookAuthProvider, GoogleAuthProvider} from "firebase/auth"
import {StyledFirebaseAuth} from "./StyledFirebaseAuth"

const uiConfig = {
    signInFlow: 'popup',
    signInSuccessUrl: '/signedIn',
    signInOptions: [
        EmailAuthProvider.PROVIDER_ID,
        GoogleAuthProvider.PROVIDER_ID,
        FacebookAuthProvider.PROVIDER_ID,
    ],
}

export const Auth: React.FC = () => {
    const [isSignedIn, setIsSignedIn] = useState(false)

    useEffect(() => {
        const unregisterAuthObserver = firebaseAuth.onAuthStateChanged(user => {
            setIsSignedIn(!!user)
        })
        return () => unregisterAuthObserver()
    }, [])

    if (!isSignedIn) {
        return (
            <div>
                <p>Please sign-in:</p>
                <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebaseAuth}/>
            </div>
        )
    }

    return (
        <div>
            <p>Welcome {firebaseAuth.currentUser!.displayName}!</p>
            <button onClick={() => firebaseAuth.signOut()}>Sign-out</button>
        </div>
    )
}
