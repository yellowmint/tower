import {grpc} from "@improbable-eng/grpc-web"
import {Send} from "@mui/icons-material"
import {LoadingButton} from "@mui/lab"
import {Box, FormGroup, FormHelperText, FormLabel, Input, Typography} from "@mui/material"
import {BrowserHeaders} from "browser-headers"
import {useSnackbar} from "notistack"
import {useState} from "react"
import {Controller, SubmitHandler, useForm} from "react-hook-form"
import packageJSON from "../../package.json"
import {useBackend} from "../backend/BackendContextProvider"
import {checkError, handleCommonErrors} from "../backend/errors"
import {CreateMyAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"
import {SignOutButton} from "./SignOutButton"


export type RegistrationProps = {
    token: string,
    initName: string | undefined | null
    successCallback: () => void
}

type Inputs = {
    name: string,
}

export const Registration = (props: RegistrationProps) => {
    const backend = useBackend()
    const {enqueueSnackbar} = useSnackbar()

    const [serverFeedback, setServerFeedback] = useState<string | null>(null)
    const {control, handleSubmit, formState: {errors}} = useForm<Inputs>()

    const onSubmit: SubmitHandler<Inputs> = (data) => {
        setServerFeedback("processing...")

        const headers = new BrowserHeaders({
            "app-version": `tower-spa:v${packageJSON.version}`,
            "authorization": `bearer ${props.token}`,
        })

        const request = new CreateMyAccountRequest()
        request.setName(data.name)

        backend.services.accounts.createMyAccount(request, headers, (err, _) => {
            if (checkError(err, grpc.Code.AlreadyExists, "account already created")) {
                setServerFeedback("account already created")
                return
            }
            if (err) {
                setServerFeedback("error")
                handleCommonErrors(err, enqueueSnackbar, "register")
                return
            }

            setServerFeedback(null)
            enqueueSnackbar("account created", {variant: "success"})
            props.successCallback()
        })
    }

    return (
        <Box component="form" onSubmit={handleSubmit(onSubmit)}>
            <Typography component="h2" variant="h6" sx={{marginBottom: 5}}>
                Please finish registration
            </Typography>
            <FormGroup>
                <FormLabel>
                    Name
                </FormLabel>
                <Controller
                    name="name"
                    control={control}
                    defaultValue={props.initName || ""}
                    rules={{
                        required: true,
                        minLength: 6,
                        maxLength: 16,
                        pattern: /^[a-zA-Z\d]+$/,
                    }}
                    render={({field}) => <Input {...field} />}
                />
                <FormHelperText error>
                    {errors.name?.type === "required" && "Field is required"}
                    {errors.name?.type === "minLength" && "Minimum length is 6"}
                    {errors.name?.type === "maxLength" && "Max length is 16"}
                    {errors.name?.type === "pattern" && "Can contain only letters and digits"}
                </FormHelperText>
            </FormGroup>
            <LoadingButton
                type="submit"
                fullWidth
                variant="contained"
                loading={serverFeedback === "processing..."}
                loadingPosition="end"
                endIcon={<Send/>}
                sx={{mt: 3, mb: 2}}
            >
                Register
            </LoadingButton>
            <FormHelperText>{serverFeedback}</FormHelperText>

            <Box sx={{marginTop: 10}}>
                <SignOutButton/>
            </Box>
        </Box>
    )
}
