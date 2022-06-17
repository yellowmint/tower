import {Box, Link, Typography} from "@mui/material"
import {EmailAuthProvider, FacebookAuthProvider, GoogleAuthProvider, IdTokenResult} from "firebase/auth"
import {useEffect, useState} from "react"
import {Link as RouterLink, useLocation, useNavigate} from "react-router-dom"
import {BackendContextActions, useBackend} from "../backend/BackendContextProvider"
import {firebaseAuth} from "../firebase/firebase"
import {Registration, RegistrationProps} from "./Registration"
import {StyledFirebaseAuth} from "./StyledFirebaseAuth"


const uiConfig = {
    signInFlow: 'popup',
    signInOptions: [
        GoogleAuthProvider.PROVIDER_ID,
        FacebookAuthProvider.PROVIDER_ID,
        {
            provider: EmailAuthProvider.PROVIDER_ID,
            requireDisplayName: false,
        },
    ],
}

export const SignIn = () => {
    const backend = useBackend()
    const navigate = useNavigate()
    const location = useLocation()

    const [registrationData, setRegistrationData] = useState<RegistrationProps | null>(null)

    const locationState = location.state as { from: { pathname: string } }
    const locationFrom = locationState?.from?.pathname || "/"

    if (backend.isAuthorized) navigate(locationFrom, {replace: true})

    useEffect(() => {
        const signIn = (tokenResult: IdTokenResult) => {
            if (tokenResult.claims["accountId"]) {
                backend.dispatch!({type: BackendContextActions.AuthChanged, payload: {jwt: tokenResult.token}})
                navigate(locationFrom, {replace: true})
                return
            }

            setRegistrationData({
                token: tokenResult.token,
                initName: firebaseAuth.currentUser!.displayName,
                successCallback: () => {
                    firebaseAuth.currentUser!.getIdTokenResult(true).then(signIn).catch(signInFailed)
                },
            })
        }

        const signInFailed = () => {
            setRegistrationData(null)
            backend.dispatch!({type: BackendContextActions.AuthChanged, payload: {jwt: null}})
        }

        const unregisterAuthObserver = firebaseAuth.onAuthStateChanged(authUser => {
            if (!authUser) return signInFailed()

            authUser.getIdTokenResult().then(signIn).catch(signInFailed)
        })
        return () => unregisterAuthObserver()
    }, [backend.dispatch, navigate, locationFrom])

    return (
        <Box sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            marginTop: 10,
            padding: 10,
        }}>
            <Typography component="h1" variant="h4" sx={{marginBottom: 3}}>
                Sign in
            </Typography>

            {locationFrom !== "/" && <p>You must sign in to view the page at {locationFrom}</p>}

            <Box sx={{marginTop: 2, width: "100%", maxWidth: 800}}>
                {registrationData === null ?
                    <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebaseAuth}/>
                    :
                    <Registration {...registrationData} />
                }
            </Box>
            <Box sx={{marginTop: 10}}>
                <Link component={RouterLink} to="/">Back to home page</Link>
            </Box>
        </Box>
    )
}
