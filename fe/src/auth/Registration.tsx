import {Box, Button, FormGroup, FormHelperText, FormLabel, Input} from "@mui/material"
import {BrowserHeaders} from "browser-headers"
import {useState} from "react"
import {Controller, SubmitHandler, useForm} from "react-hook-form"
import {useBackend} from "../backend/BackendContextProvider"
import {CreateMyAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"


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

        const headers = new BrowserHeaders({"authorization": `bearer ${props.token}`})

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
        <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{mt: 1}}>
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
                        pattern: /^[a-zA-Z0-9]+$/,
                    }}
                    render={({field}) => <Input {...field} />}
                />
                <FormHelperText>
                    {errors.name?.type === "required" && "Name is required"}
                    {errors.name?.type === "minLength" && "Minimum length is 6"}
                    {errors.name?.type === "maxLength" && "Max length is 16"}
                    {errors.name?.type === "pattern" && "Name can contain only letters and digits"}
                </FormHelperText>
            </FormGroup>
            <Button type="submit" fullWidth variant="contained" sx={{mt: 3, mb: 2}}>Submit</Button>
            <FormHelperText>{serverStatus}</FormHelperText>
        </Box>
    )
}
