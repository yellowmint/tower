import {Box, CircularProgress, Typography} from "@mui/material"
import {useEffect, useState} from "react"
import {useNavigate} from "react-router-dom"
import {firebaseAuth} from "../firebase/firebase"

export const SignOut = () => {
    const navigate = useNavigate()
    const [status, setStatus] = useState<string>("signing out...")

    useEffect(() => {
        let timer: NodeJS.Timeout

        firebaseAuth.signOut()
            .then(() => {
                setStatus("success, redirecting...")
                timer = setTimeout(() => navigate("/"), 1000)
            })
            .catch(() => setStatus("failed to sign out"))

        return () => {
            if (timer) clearTimeout(timer)
        }
    }, [navigate])

    return (
        <Box sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            marginTop: 10,
        }}>
            <Typography component="h1" variant="h5">Sign out</Typography>
            <br/>
            <CircularProgress/>
            <p>{status}</p>
        </Box>
    )
}
