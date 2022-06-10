import React from "react"
import {AccountsServiceClient} from "../contracts/accounts/rpcpublic/v1/accounts_pb_service"
import {GetAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"

export const GetAccountDetails: React.FC = () => {
    const handleClick = () => {
        console.log("info")
        const client = new AccountsServiceClient("http://localhost:4040")
        const req = new GetAccountRequest()

        client.getAccount(req, (err, resp) => {
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
