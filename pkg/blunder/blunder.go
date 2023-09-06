package blunder

type OrdinaryError interface {
	Error() string
	GetId() string
	GetHttpStatusCode() int
	GetGrpcStatusCode() int
	GetErrorCode() string
	GetMessage() string
	Wrap(err error)
	Unwrap() error
	Is(err error) bool
}
