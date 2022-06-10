import React from 'react'
import logo from './logo.svg'
import './App.css'
import {GetAccountDetails} from "./accounts/GetAccountDetails"
import {Auth} from "./auth/Auth"

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo"/>
                <Auth/>
                <GetAccountDetails/>
            </header>
        </div>
    )
}

export default App
