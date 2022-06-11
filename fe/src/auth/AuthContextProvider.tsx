import {createContext, Dispatch, ReactNode, useContext, useReducer} from "react"
import {User} from "firebase/auth"

type AuthContextType = {
    user: User | null,
    dispatch: Dispatch<AuthContextAction> | undefined
} | undefined

const initialState: AuthContextType = {
    user: null,
    dispatch: undefined,
}

export enum AuthContextActionKind {
    SignedIn = "SIGNED_IN",
    SignedOut = "SIGNED_OUT",
}

type AuthContextAction =
    | { type: AuthContextActionKind.SignedIn, payload: { user: User } }
    | { type: AuthContextActionKind.SignedOut }

const authReducer = (state: AuthContextType, action: AuthContextAction): AuthContextType => {
    switch (action.type) {
        case AuthContextActionKind.SignedIn:
            return {...state!, user: action.payload.user}

        case AuthContextActionKind.SignedOut:
            return {...state!, user: null}

        default:
            throw new Error(`Unhandled action type: ${action['type']}`)
    }
}


const AuthContext = createContext<AuthContextType>(undefined)

export const AuthContextProvider = ({children}: { children: ReactNode }) => {
    const [reducerState, dispatch] = useReducer(authReducer, initialState)
    const {user} = reducerState!

    return (
        <AuthContext.Provider value={{user, dispatch}}>
            {children}
        </AuthContext.Provider>
    )
}

export const useAuth = () => {
    const context = useContext(AuthContext)
    if (context === undefined) {
        throw new Error('useAuth must be used within a AuthContextProvider')
    }

    return context
}
