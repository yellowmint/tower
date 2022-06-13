import {SubmitHandler, useForm} from "react-hook-form"
import {BrowserHeaders} from "browser-headers"
import {CreateMyAccountRequest} from "../contracts/accounts/rpcpublic/v1/accounts_pb"
import {firebaseAuth} from "../firebase/firebase"

type Inputs = {
    example: string,
    exampleRequired: string,
};

export const Registration = () => {
    const {register, handleSubmit, watch, formState: {errors}} = useForm<Inputs>()
    const onSubmit: SubmitHandler<Inputs> = data => console.log(data)

    console.log(watch("example")) // watch input value by passing the name of it

    const register55 = () => {
        const headers = new BrowserHeaders({"authorization": `bearer ${registrationToken}`})

        const request = new CreateMyAccountRequest()
        request.setName(firebaseAuth!.currentUser!.displayName!)

        backend.services.accounts.createMyAccount(request, headers, (err, result) => {
            console.log(err)
            console.log(result)
        })
    }

    return (
        /* "handleSubmit" will validate your inputs before invoking "onSubmit" */
        <form onSubmit={handleSubmit(onSubmit)}>
            {/* register your input into the hook by invoking the "register" function */}
            <input defaultValue="test" {...register("example")} />

            {/* include validation with required or other standard HTML validation rules */}
            <input {...register("exampleRequired", {required: true})} />
            {/* errors will return when field validation fails  */}
            {errors.exampleRequired && <span>This field is required</span>}

            <input type="submit"/>
        </form>
    )
}