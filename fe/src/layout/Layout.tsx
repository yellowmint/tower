import {Box, Container, Typography} from "@mui/material"
import {Outlet} from "react-router-dom"
import {Authorization} from "../auth/Authorization"

export const Layout = () => {
    return (
        <Container component="main" maxWidth="md">
            <Box
                sx={{
                    marginTop: 15,
                    marginBottom: 10,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <Typography component="h1" variant="h5">Tower</Typography>
                <nav>
                    <Authorization/>
                </nav>
            </Box>
            <Outlet/>
        </Container>
    )
}
