import {HomeOutlined, Person} from "@mui/icons-material"
import {Box, Button} from "@mui/material"
import {NavLink} from "react-router-dom"
import {Authorization} from "../auth/Authorization"
import style from "./navigation.module.scss"

export const Navigation = () => {
    return (
        <Box
            sx={{
                marginTop: 5,
                marginBottom: 5,
                display: "flex",
                flexDirection: "row",
                justifyContent: "space-evenly",
                width: "100%",
                maxWidth: 800,
            }}
            component="nav"
            className={style.navigation}
        >
            <NavLink to="/" children={({isActive}) => (
                <Button variant={isActive ? "contained" : "outlined"} endIcon={<HomeOutlined/>}>
                    Home
                </Button>
            )}/>
            <NavLink to="/account" children={({isActive}) => (
                <Button variant={isActive ? "contained" : "outlined"} endIcon={<Person/>}>
                    Account
                </Button>
            )}/>
            <Authorization/>
        </Box>
    )
}
