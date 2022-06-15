import {CircularProgress, Typography} from "@mui/material"
import {useEffect, useState} from "react"
import {useBackend} from "../backend/BackendContextProvider"
import {GetMyAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"
import {AccountData, AccountDetails} from "./AccountDetails"
import {DeleteAccount} from "./DeleteAccount"

export const Account = () => {
    const backend = useBackend()
    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [serverStatus, setServerStatus] = useState<string | null>(null)
    const [accountData, setAccountData] = useState<AccountData | null>(null)

    useEffect(() => {
        const request = new GetMyAccountRequest()

        backend.services.accounts.getMyAccount(request, backend.headers, (err, response) => {
            setIsLoading(false)

            if (err) {
                console.log(err)
                setServerStatus("server error")
                return
            }

            setServerStatus(null)
            setAccountData({name: response!.getName(), accountId: response!.getAccountId()})
        })
    }, [])

    return (
        <>
            <Typography component="h2" variant="h4" sx={{marginBottom: 3}}>
                Account
            </Typography>
            {isLoading ?
                <CircularProgress/>
                :
                <>
                    {accountData && <AccountDetails {...{accountData}} />}
                    <DeleteAccount/>
                    <p>{serverStatus}</p>
                </>
            }
        </>
    )
}
