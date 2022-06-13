import {Box, Container, createTheme, CssBaseline, ThemeProvider, Typography} from "@mui/material"
import React from "react"
import {GetAccountDetails} from "./accounts/GetAccountDetails"
import {SignIn} from "./auth/SignIn"
import {BackendContextProvider} from "./backend/BackendContextProvider"

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
    },
})

export const App = () => {
    return (
        <>
            <ThemeProvider theme={darkTheme}>
                <CssBaseline/>
                <BackendContextProvider>
                    <Container component="main" maxWidth="md">
                        <Box
                            sx={{
                                marginTop: 15,
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'center',
                            }}
                        >
                            <Typography component="h1" variant="h5">Tower</Typography>
                        </Box>
                        <SignIn/>
                        <GetAccountDetails/>
                    </Container>
                </BackendContextProvider>
            </ThemeProvider>
        </>
    )
}
