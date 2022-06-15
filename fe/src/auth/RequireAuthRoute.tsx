import {Navigate, useLocation} from "react-router-dom"
import {useBackend} from "../backend/BackendContextProvider"

export const RequireAuth = ({children}: { children: JSX.Element }) => {
    const backend = useBackend()
    const location = useLocation()

    if (!backend.isAuthorized) {
        return <Navigate to="/sign-in" state={{from: location}} replace/>
    }

    return children
}
