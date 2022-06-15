import {Box, Button, FormGroup, FormHelperText, FormLabel, Input, Typography} from "@mui/material"
import {BrowserHeaders} from "browser-headers"
import {useState} from "react"
import {Controller, SubmitHandler, useForm} from "react-hook-form"
import packageJSON from "../../package.json"
import {useBackend} from "../backend/BackendContextProvider"
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
    const [serverStatus, setServerStatus] = useState<string | null>(null)
    const {control, handleSubmit, formState: {errors}} = useForm<Inputs>()

    const onSubmit: SubmitHandler<Inputs> = (data) => {
        setServerStatus("processing...")

        const headers = new BrowserHeaders({
            "app-version": `tower-spa:v${packageJSON.version}`,
            "authorization": `bearer ${props.token}`,
        })

        const request = new CreateMyAccountRequest()
        request.setName(data.name)

        backend.services.accounts.createMyAccount(request, headers, (err, result) => {
            if (err?.message === "account already created") {
                setServerStatus("account already created")
                return
            }
            if (err) {
                console.log(err)
                setServerStatus("server error")
                return
            }

            setServerStatus(null)
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
            <Button type="submit" fullWidth variant="contained" sx={{mt: 3, mb: 2}}>Submit</Button>
            <FormHelperText>{serverStatus}</FormHelperText>

            <Box sx={{marginTop: 10}}>
                <SignOutButton/>
            </Box>
        </Box>
    )
}
