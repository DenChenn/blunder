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

var (
	ErrTargetIsNotGrpcStandardStatus = &Unexpected{
		Id:             "blunder:target_is_not_grpc_standard_status",
		HttpStatusCode: 500,
		GrpcStatusCode: 13,
		ErrorCode:      "TARGET_IS_NOT_GRPC_STANDARD_STATUS",
		Message:        "[blunder error]: Given target error is not a grpc standard status error",
	}
	ErrComparativeIsNotBlunderError = &Unexpected{
		Id:             "blunder:comparative_is_not_blunder_error",
		HttpStatusCode: 500,
		GrpcStatusCode: 13,
		ErrorCode:      "COMPARATIVE_IS_NOT_BLUNDER_ERROR",
		Message:        "[blunder error]: Some of the given comparative error is not a blunder error, specifying FromGrpc is not allowed",
	}
	ErrUndefined = &Unexpected{
		Id:             "undefined_error",
		HttpStatusCode: 500,
		GrpcStatusCode: 13,
		ErrorCode:      "UNDEFINED_ERROR",
		Message:        "Unexpected error happened.",
	}
)

var (
	_ Error = ErrTargetIsNotGrpcStandardStatus
	_ Error = ErrComparativeIsNotBlunderError
)

type Unexpected struct {
	Id             string `json:"id"`
	HttpStatusCode int    `json:"http_status_code"`
	GrpcStatusCode int    `json:"grpc_status_code"`
	ErrorCode      string `json:"error_code"`
	Message        string `json:"message"`
}

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
