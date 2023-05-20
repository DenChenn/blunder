package model

type Blunder struct {
	Details []Detail `yaml:"details"`
}

type Detail struct {
	Package string             `yaml:"package"`
	Errors  []ErrorDescription `yaml:"errors"`
}

type ErrorDescription struct {
	Id             string `yaml:"id"`
	Code           string `yaml:"code"`
	HttpStatusCode int    `yaml:"http_status_code"`
	GrpcStatusCode int    `yaml:"grpc_status_code"`
	Message        string `yaml:"message"`
}
