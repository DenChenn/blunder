package blunder

type Unexpected struct {
	HttpStatusCode int    `json:"http_status_code"`
	GrpcStatusCode int    `json:"grpc_status_code"`
	ErrorCode      string `json:"error_code"`
	Message        string `json:"message"`
	Err            error  `json:"err"`
}

var (
	ErrUndefined = &Unexpected{
		HttpStatusCode: 500,
		GrpcStatusCode: 13,
		ErrorCode:      "UNDEFINED_ERROR",
		Message:        "Unexpected error happened.",
	}
	_ error = ErrUndefined
)

func (e *Unexpected) Error() string {
	return e.Message
}

func (e *Unexpected) GetHttpStatusCode() int {
	return e.HttpStatusCode
}

func (e *Unexpected) GetGrpcStatusCode() int {
	return e.GrpcStatusCode
}

func (e *Unexpected) GetErrorCode() string {
	return e.ErrorCode
}

func (e *Unexpected) GetMessage() string {
	return e.Message
}

func (e *Unexpected) Wrap(err error) {
	e.Err = err
}

func (e *Unexpected) Unwrap() error {
	return e.Err
}
