package tower

type Err struct {
	ErrorValue     error
	EndUserMessage string
}

func (e Err) Error() string {
	return e.ErrorValue.Error()
}

func UnhandledError(err error) Err {
	return Err{
		ErrorValue:     err,
		EndUserMessage: "unhandled error",
	}
}
