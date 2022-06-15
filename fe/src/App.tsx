import {createTheme, CssBaseline, ThemeProvider} from "@mui/material"
import {BrowserRouter, Route, Routes} from "react-router-dom"
import {Account} from "./account/Account"
import {RequireAuth} from "./auth/RequireAuthRoute"
import {SignIn} from "./auth/SignIn"
import {SignOut} from "./auth/SignOut"
import {BackendContextProvider} from "./backend/BackendContextProvider"
import {Home} from "./home/Home"
import {Layout} from "./layout/Layout"
import {NotFound} from "./layout/NotFound"

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
    },
})

export const App = () => {
    return (
        <ThemeProvider theme={darkTheme}>
            <CssBaseline/>
            <BackendContextProvider>
                <BrowserRouter>
                    <Routes>
                        <Route path="/sign-in" element={<SignIn/>}/>
                        <Route path="/sign-out" element={<SignOut/>}/>

                        <Route path="/" element={<Layout/>}>
                            <Route index element={<Home/>}/>
                            <Route path="account" element={<RequireAuth><Account/></RequireAuth>}/>
                            <Route path="*" element={<NotFound/>}/>
                        </Route>
                    </Routes>
                </BrowserRouter>
            </BackendContextProvider>
        </ThemeProvider>
    )
}
