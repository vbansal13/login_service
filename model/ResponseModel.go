package model

//ResponseResultModel describes generic response message data model
type ResponseResultModel struct {
	Error  string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
}
