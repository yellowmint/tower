import {grpc} from "@improbable-eng/grpc-web"
import {Send} from "@mui/icons-material"
import {LoadingButton} from "@mui/lab"
import {Box, FormGroup, FormHelperText, FormLabel, Input, Typography} from "@mui/material"
import {useSnackbar} from "notistack"
import {useState} from "react"
import {Controller, SubmitHandler, useForm} from "react-hook-form"
import {fullName} from "../account/AccountDetails"
import {useBackend} from "../backend/BackendContextProvider"
import {checkError, handleCommonErrors} from "../backend/errors"
import {GetAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"

type Inputs = {
    accountId: string,
}

export const GetAccountDetails = () => {
    const backend = useBackend()
    const {enqueueSnackbar} = useSnackbar()

    const [serverFeedback, setServerFeedback] = useState<string | null>(null)
    const {control, handleSubmit, formState: {errors}} = useForm<Inputs>()

    const onSubmit: SubmitHandler<Inputs> = (data) => {
        setServerFeedback("processing...")

        const request = new GetAccountRequest()
        request.setAccountId(data.accountId.toLowerCase())

        backend.services.accounts.getAccount(request, backend.headers, (err, response) => {
            if (checkError(err, grpc.Code.NotFound, "account not found")) {
                setServerFeedback("account not found")
                return
            }
            if (err) {
                setServerFeedback("error")
                handleCommonErrors(err, enqueueSnackbar, "find account")
                return
            }

            const name = fullName({
                base: response!.getName()!.getBase(),
                number: response!.getName()!.getNumber(),
            })

            setServerFeedback(`Account name: ${name}`)
        })
    }

    if (!backend.isAuthorized) return <></>

    return (
        <Box component="form" onSubmit={handleSubmit(onSubmit)}>
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
                loading={serverFeedback === "processing..."}
                loadingPosition="end"
                endIcon={<Send/>}
                sx={{mt: 3, mb: 2}}
            >
                Get account info
            </LoadingButton>
            <FormHelperText>{serverFeedback}</FormHelperText>
        </Box>
    )
}
