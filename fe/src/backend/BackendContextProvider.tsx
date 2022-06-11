import {createContext, Dispatch, ReactNode, useContext, useReducer} from "react"
import {AccountsServiceClient} from "../contracts/accounts/rpcpublic/v1/accounts_pb_service"
import {BrowserHeaders} from "browser-headers"


type BackendContextType = {
    headers: BrowserHeaders,
    services: {
        accounts: AccountsServiceClient
    },
    dispatch: Dispatch<BackendContextAction> | undefined
} | undefined

const initialState: BackendContextType = {
    headers: new BrowserHeaders(),
    services: {
        accounts: new AccountsServiceClient(process.env.REACT_APP_ACCOUNT_SERVICE_URL!)
    },
    dispatch: undefined,
}

enum BackendContextActionKind {
    AuthChanged = "AUTH_CHANGED",
}

type BackendContextAction = { type: string, payload: any }
    | { type: BackendContextActionKind.AuthChanged, payload: { jwt: string } }

const backendReducer = (state: BackendContextType, action: BackendContextAction): BackendContextType => {
    switch (action.type) {
        case BackendContextActionKind.AuthChanged:
            console.log("Change", action.payload.jwt)
            return state

        default:
            throw new Error(`Unhandled action type: ${action.type}`)
    }
}


const BackendContext = createContext<BackendContextType>(undefined)

export const BackendContextProvider = ({children}: { children: ReactNode }) => {
    const [reducerState, dispatch] = useReducer(backendReducer, initialState)
    const {headers, services} = reducerState!

    return (
        <BackendContext.Provider value={{headers, services, dispatch}}>
            <div>
                {children}
            </div>
            <div>
                <button onClick={() => dispatch({type: BackendContextActionKind.AuthChanged, payload: {jwt: "123"}})}>
                    ABC
                </button>
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
