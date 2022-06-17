import {DeleteForever} from "@mui/icons-material"
import {LoadingButton} from "@mui/lab"
import {
    Box,
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    TextField,
} from "@mui/material"
import {useSnackbar} from "notistack"
import {useState} from "react"
import {useNavigate} from "react-router-dom"
import {useBackend} from "../backend/BackendContextProvider"
import {handleCommonErrors} from "../backend/errors"
import {DeleteMyAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"

type DeleteAccountProps = {
    accountName: string
}

export const DeleteAccount = (props: DeleteAccountProps) => {
    const backend = useBackend()
    const navigate = useNavigate()
    const {enqueueSnackbar} = useSnackbar()

    const [isDialogOpen, setIsDialogOpen] = useState<boolean>(false)
    const [confirmationValue, setConfirmationValue] = useState<string>("")
    const [isProcessing, setIsProcessing] = useState<boolean>(false)

    const openDialog = () => {
        setIsDialogOpen(true)
    }

    const handleClose = (confirmed: boolean) => {
        if (!confirmed) {
            setIsDialogOpen(false)
            setConfirmationValue("")
            return
        }

        deleteAccount()
    }

    const deleteAccount = () => {
        setIsProcessing(true)

        const request = new DeleteMyAccountRequest()

        backend.services.accounts.deleteMyAccount(request, backend.headers, (err, _) => {
            if (err) return handleCommonErrors(err, enqueueSnackbar, "delete account")

            enqueueSnackbar("account deleted", {variant: "success"})
            navigate("/sign-out")
        })
    }

    return (
        <Box sx={{marginTop: 10, marginBottom: 10}}>
            <Button
                variant="outlined"
                color="error"
                onClick={openDialog}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                Delete account
            </Button>
            <Dialog open={isDialogOpen} onClose={() => handleClose(false)}>
                <DialogTitle id="alert-dialog-title">
                    Are you sure you want to delete your account?
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        You will lose access to your account and your account will not be visible in
                        the system. It is not possible to undo this operation.<br/><br/>
                        Enter your account name to confirm delete action:
                    </DialogContentText>
                    <TextField
                        label="Your account name"
                        value={confirmationValue}
                        onChange={e => setConfirmationValue(e.target.value)}
                        fullWidth
                        sx={{mt: 2}}
                    />
                </DialogContent>
                <DialogActions>
                    <Button
                        onClick={() => handleClose(false)}
                        autoFocus
                        disabled={isProcessing}
                    >
                        Close
                    </Button>
                    <LoadingButton
                        onClick={() => handleClose(true)}
                        color="error"
                        loading={isProcessing}
                        disabled={confirmationValue !== props.accountName}
                        loadingPosition="end"
                        endIcon={<DeleteForever/>}
                    >
                        Confirm and delete
                    </LoadingButton>
                </DialogActions>
            </Dialog>
        </Box>
    )
}
