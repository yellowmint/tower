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
    Typography,
} from "@mui/material"
import {useState} from "react"
import {useNavigate} from "react-router-dom"
import {useBackend} from "../backend/BackendContextProvider"
import {DeleteMyAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"

export const DeleteAccount = () => {
    const backend = useBackend()
    const navigate = useNavigate()
    const [openDialog, setOpenDialog] = useState<boolean>(false)
    const [serverStatus, setServerStatus] = useState<string | null>(null)
    const [isProcessing, setIsProcessing] = useState<boolean>(false)

    const handleClickOpen = () => {
        setOpenDialog(true)
    }

    const handleClose = (confirmed: boolean) => {
        if (!confirmed) return setOpenDialog(false)

        setIsProcessing(true)

        const request = new DeleteMyAccountRequest()

        backend.services.accounts.deleteMyAccount(request, backend.headers, (err, _) => {
            if (err) {
                console.log(err)
                setServerStatus("server error")
                return
            }

            setIsProcessing(false)
            setServerStatus(null)
            navigate("/sign-out")
        })
    }

    return (
        <Box sx={{marginTop: 10, marginBottom: 10}}>
            <Button
                variant="outlined"
                color="error"
                onClick={handleClickOpen}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                Delete account
            </Button>
            <Dialog open={openDialog} onClose={() => handleClose(false)}>
                <DialogTitle id="alert-dialog-title">
                    Are you sure you want to delete your account?
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        After doing this, you will lose access to your account and your account will not be visible in
                        the system.
                        However, the complete deletion of the account and anonymization of your data will take place
                        after 30 days.
                        {serverStatus && <Typography component="span" color="orange">
                            <br/><br/>{serverStatus}
                        </Typography>}
                    </DialogContentText>
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
