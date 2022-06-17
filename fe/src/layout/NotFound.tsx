import {Box, Typography} from "@mui/material"

export const NotFound = () => {
    return (
        <Box sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
        }}>
            <Typography component="h2" variant="h5">
                Not Found
            </Typography>
        </Box>
    )
}
