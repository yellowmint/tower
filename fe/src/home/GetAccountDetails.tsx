import {Send} from "@mui/icons-material"
import {LoadingButton} from "@mui/lab"
import {Box, FormGroup, FormHelperText, FormLabel, Input, Typography} from "@mui/material"
import {useState} from "react"
import {Controller, SubmitHandler, useForm} from "react-hook-form"
import {useBackend} from "../backend/BackendContextProvider"
import {GetAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"

type Inputs = {
    accountId: string,
}

export const GetAccountDetails = () => {
    const backend = useBackend()
    const [serverStatus, setServerStatus] = useState<string | null>(null)
    const {control, handleSubmit, formState: {errors}} = useForm<Inputs>()

    const onSubmit: SubmitHandler<Inputs> = (data) => {
        setServerStatus("processing...")

        const request = new GetAccountRequest()
        request.setAccountId(data.accountId.toLowerCase())

        backend.services.accounts.getAccount(request, backend.headers, (err, response) => {
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

    if (!backend.isAuthorized) return <></>

    return (
        <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{mt: 1, marginTop: 10}}>
            <Typography component="h3" variant="h6" sx={{marginBottom: 3}}>
                Find account
            </Typography>

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

            <LoadingButton
                type="submit"
                fullWidth
                loading={serverStatus === "processing..."}
                loadingPosition="end"
                endIcon={<Send/>}
                sx={{mt: 3, mb: 2}}
            >
                Get account info
            </LoadingButton>
            <FormHelperText>{serverStatus}</FormHelperText>
        </Box>
    )
}
