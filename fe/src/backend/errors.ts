import {grpc} from "@improbable-eng/grpc-web"
import {OptionsObject, SnackbarKey, SnackbarMessage} from "notistack"
import {ServiceError} from "../contracts/accounts/rpcpublic/v1/accounts_pb_service"

type errors = ServiceError
type enqueueSnackbarType = (message: SnackbarMessage, options?: OptionsObject) => SnackbarKey

export const checkError = (err: errors | null, code: grpc.Code, message: string): boolean => {
    return err?.code === code && err?.message === message
}

export const handleCommonErrors = (err: errors, enqueueSnackbar: enqueueSnackbarType, action: string) => {
    if (err.code === grpc.Code.Unauthenticated) {
        enqueueSnackbar(`failed to ${action} - authentication error`, {variant: "error"})
        return
    }

    console.log(err)
    enqueueSnackbar(`failed to ${action} - server error`, {variant: "error"})
}
