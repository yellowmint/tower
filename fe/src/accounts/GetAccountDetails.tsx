import {GetAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"
import {useBackend} from "../backend/BackendContextProvider"

export const GetAccountDetails = () => {
    const be = useBackend()

    const handleClick = () => {
        const req = new GetAccountRequest()
        be.services.accounts.getAccount(req, be.headers, (err, resp) => {
            console.log("err: ", err)
            console.log("resp: ", resp)
        })
    }

    return (
        <div>
            <button onClick={handleClick}>Get account info</button>
        </div>
    )
}
