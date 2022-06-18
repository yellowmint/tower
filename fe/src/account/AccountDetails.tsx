export type AccountData = {
    accountId: string
    name: AccountName
}

export type AccountName = {
    base: string
    number: number
}

export const fullName = (name: AccountName): string => {
    return `${name.base}/${name.number}`
}

export const AccountDetails = ({accountData}: { accountData: AccountData }) => {
    return (
        <>
            <p>AccountId: {accountData.accountId}</p>
            <p>Name: {fullName(accountData.name)}</p>
        </>
    )
}
