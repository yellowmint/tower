package tower

type Err struct {
	ErrorValue     error
	EndUserMessage string
}

func (e Err) Error() string {
	return e.ErrorValue.Error()
}
