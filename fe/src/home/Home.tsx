import {Typography} from "@mui/material"
import {GetAccountDetails} from "./GetAccountDetails"

export const Home = () => {
    return (
        <>
            <Typography component="h2" variant="h4" sx={{marginBottom: 3}}>
                Home
            </Typography>
            <GetAccountDetails/>
        </>
    )
}
