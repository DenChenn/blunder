package blunder

type Error interface {
	Error() string
	GetId() string
	GetHttpStatusCode() int
	GetGrpcStatusCode() int
	GetErrorCode() string
	GetMessage() string
	WithCustomMessage(msg string) Error
}

type Unexpected struct {
	Id             string `json:"id"`
	HttpStatusCode int    `json:"http_status_code"`
	GrpcStatusCode int    `json:"grpc_status_code"`
	ErrorCode      string `json:"error_code"`
	Message        string `json:"message"`
}

var (
	ErrUndefined = &Unexpected{
		Id:             "undefined_error",
		HttpStatusCode: 500,
		GrpcStatusCode: 13,
		ErrorCode:      "UNDEFINED_ERROR",
		Message:        "Unexpected error happened.",
	}
	_ Error = ErrUndefined
)

func (e *Unexpected) Error() string {
	return e.Message
}

func (e *Unexpected) GetId() string {
	return e.Id
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

func (e *Unexpected) WithCustomMessage(msg string) Error {
	newE := *e
	newE.Message = msg
	return &newE
}
