import Logout from "@mui/icons-material/Logout"
import {Button} from "@mui/material"
import {Link} from "react-router-dom"

export const SignOutButton = () => (
    <Button to="/sign-out" component={Link} variant="outlined" endIcon={<Logout/>}>
        Sign out
    </Button>
)
