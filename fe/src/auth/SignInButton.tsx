import {Login} from "@mui/icons-material"
import {Button} from "@mui/material"
import {Link} from "react-router-dom"

export const SignInButton = () => (
    <Button to="/sign-in" component={Link} variant="outlined" endIcon={<Login/>}>
        Sign in
    </Button>
)
