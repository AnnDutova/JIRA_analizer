package models

type IIssue interface {
	Error() string
}

type ConnectorError struct {
	ErrorMessage string `json:"error"`
}

func (connectorError ConnectorError) Error() string {
	return connectorError.ErrorMessage
}
