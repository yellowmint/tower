import {BrowserHeaders} from "browser-headers"
import {createContext, Dispatch, ReactNode, useContext, useReducer} from "react"
import packageJSON from "../../package.json"
import {AccountsServiceClient} from "../contracts/accounts/rpcpublic/v1/accounts_pb_service"

type BackendContextType = {
    headers: BrowserHeaders,
    services: {
        accounts: AccountsServiceClient
    },
    isAuthorized: boolean,
    dispatch: Dispatch<BackendContextAction> | undefined
} | undefined

const initialState: BackendContextType = {
    headers: new BrowserHeaders({"app-version": `tower-spa:v${packageJSON.version}`}),
    services: {
        accounts: new AccountsServiceClient(process.env.REACT_APP_ACCOUNT_SERVICE_URL!),
    },
    isAuthorized: false,
    dispatch: undefined,
}

export enum BackendContextActions {
    AuthChanged = "AUTH_CHANGED",
}

type BackendContextAction = { type: string, payload: any }
    | { type: BackendContextActions.AuthChanged, payload: { jwt: string } }

const backendReducer = (state: BackendContextType, action: BackendContextAction): BackendContextType => {
    switch (action.type) {
        case BackendContextActions.AuthChanged:
            return authChanged(state, action.payload.jwt)

        default:
            throw new Error(`Unhandled action type: ${action.type}`)
    }
}

const authChanged = (state: BackendContextType, jwt: string | null): BackendContextType => {
    if (jwt === null || jwt === "") {
        state!.headers.set("authorization", "")
        return {...state!, isAuthorized: false}
    }

    state!.headers.set("authorization", "bearer " + jwt)
    return {...state!, isAuthorized: true}
}


const BackendContext = createContext<BackendContextType>(undefined)

export const BackendContextProvider = ({children}: { children: ReactNode }) => {
    const [reducerState, dispatch] = useReducer(backendReducer, initialState)
    const {headers, services, isAuthorized} = reducerState!

    return (
        <BackendContext.Provider value={{headers, services, isAuthorized, dispatch}}>
            <div>
                {children}
            </div>
        </BackendContext.Provider>
    )
}

export const useBackend = () => {
    const context = useContext(BackendContext)
    if (context === undefined) {
        throw new Error('useBackend must be used within a BackendContextProvider')
    }

    return context
}
