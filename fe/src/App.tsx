import React from 'react'
import logo from './logo.svg'
import './App.css'
import {GetAccountDetails} from "./accounts/GetAccountDetails"
import {SignIn} from "./auth/SignIn"
import {BackendContextProvider} from "./backend/BackendContextProvider"
import {AuthContextProvider} from "./auth/AuthContextProvider"

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo"/>
                <AuthContextProvider>
                    <SignIn/>

                    <BackendContextProvider>
                        <GetAccountDetails/>
                    </BackendContextProvider>
                </AuthContextProvider>
            </header>
        </div>
    )
}

export default App
