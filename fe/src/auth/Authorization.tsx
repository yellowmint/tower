import {IdTokenResult} from "firebase/auth"
import {useEffect} from "react"
import {BackendContextActions, useBackend} from "../backend/BackendContextProvider"
import {firebaseAuth} from "../firebase/firebase"
import {SignInButton} from "./SignInButton"
import {SignOutButton} from "./SignOutButton"


export const Authorization = () => {
    const backend = useBackend()

    const signIn = (tokenResult: IdTokenResult) => {
        if (!tokenResult.claims["accountId"]) return signOut()

        backend.dispatch!({type: BackendContextActions.AuthChanged, payload: {jwt: tokenResult.token}})
    }

    const signOut = () => {
        backend.dispatch!({type: BackendContextActions.AuthChanged, payload: {jwt: null}})
    }

    useEffect(() => {
        const unregisterAuthObserver = firebaseAuth.onAuthStateChanged(authUser => {
            if (!authUser) return signOut()

            authUser.getIdTokenResult().then(signIn).catch(signOut)
        })
        return () => unregisterAuthObserver()
    }, [])

    return (
        <>
            {backend.isAuthorized ? <SignOutButton/> : <SignInButton/>}
        </>
    )
}
