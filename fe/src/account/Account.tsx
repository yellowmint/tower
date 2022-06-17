import {CircularProgress, Typography} from "@mui/material"
import {useSnackbar} from "notistack"
import {useEffect, useState} from "react"
import {useBackend} from "../backend/BackendContextProvider"
import {handleCommonErrors} from "../backend/errors"
import {GetMyAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"
import {AccountData, AccountDetails} from "./AccountDetails"
import {DeleteAccount} from "./DeleteAccount"

export const Account = () => {
    const backend = useBackend()
    const {enqueueSnackbar} = useSnackbar()

    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [accountData, setAccountData] = useState<AccountData | null>(null)

    useEffect(() => {
        const request = new GetMyAccountRequest()

        backend.services.accounts.getMyAccount(request, backend.headers, (err, response) => {
            setIsLoading(false)

            if (err) return handleCommonErrors(err, enqueueSnackbar, "load account")

            setAccountData({name: response!.getName(), accountId: response!.getAccountId()})
        })
    }, [backend, enqueueSnackbar])

    return (
        <>
            <Typography component="h2" variant="h4" sx={{marginBottom: 3}}>
                Account
            </Typography>
            {isLoading && <CircularProgress/>}
            {accountData &&
                <>
                    <AccountDetails {...{accountData}} />
                    <DeleteAccount accountName={accountData.name}/>
                </>
            }
        </>
    )
}
