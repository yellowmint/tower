import {Box, Button, FormGroup, FormHelperText, FormLabel, Input} from "@mui/material"
import {useState} from "react"
import {Controller, SubmitHandler, useForm} from "react-hook-form"
import {useBackend} from "../backend/BackendContextProvider"
import {GetAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"

type Inputs = {
    accountId: string,
}

export const GetAccountDetails = () => {
    const be = useBackend()
    const [serverStatus, setServerStatus] = useState<string | null>(null)
    const {control, handleSubmit, formState: {errors}} = useForm<Inputs>()

    const onSubmit: SubmitHandler<Inputs> = (data) => {
        setServerStatus("processing...")

        const request = new GetAccountRequest()
        request.setAccountId(data.accountId.toLowerCase())

        be.services.accounts.getAccount(request, be.headers, (err, response) => {
            if (err?.message === "account not found") {
                setServerStatus("account not found")
                return
            }
            if (err) {
                console.log(err)
                setServerStatus("server error")
                return
            }

            setServerStatus(`Account name: ${response?.getName()}`)
        })
    }

    if (!be.isAuthorized) return <></>

    return (
        <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{mt: 1, marginTop: 10}}>
            <FormGroup>
                <FormLabel>
                    Account Id
                </FormLabel>
                <Controller
                    name="accountId"
                    control={control}
                    defaultValue={""}
                    rules={{
                        required: true,
                        pattern: /^[a-fA-F\d]{8}-[a-fA-F\d]{4}-[a-fA-F\d]{4}-[a-fA-F\d]{4}-[a-fA-F\d]{12}$/,
                    }}
                    render={({field}) => <Input {...field} />}
                />
                <FormHelperText error>
                    {errors.accountId?.type === "required" && "Field is required"}
                    {errors.accountId?.type === "pattern" && "Needs to be a valid UUID"}
                </FormHelperText>
            </FormGroup>
            <Button type="submit" fullWidth sx={{mt: 3, mb: 2}}>Get account info</Button>
            <FormHelperText>{serverStatus}</FormHelperText>
        </Box>
    )
}
