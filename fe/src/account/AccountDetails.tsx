export type AccountData = {
    name: string
    accountId: string
}

export const AccountDetails = ({accountData}: { accountData: AccountData }) => {
    return (
        <>
            <p>AccountId: {accountData.accountId}</p>
            <p>Name: {accountData.name}</p>
        </>
    )
}
