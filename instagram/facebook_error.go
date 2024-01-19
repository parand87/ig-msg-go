package instagram

type FacebookError struct {
	Message      string `json:"message,omitempty"`
	Type         string `json:"type,omitempty"`
	Code         int    `json:"code,omitempty"`
	ErrorSubcode int    `json:"error_subcode,omitempty"`
	FbtraceId    string `json:"fbtrace_id,omitempty"`
}

type FacebookErrorResponse struct {
	Error FacebookError `json:"error,omitempty"`
}

func (e *FacebookError) Error() string {
	return e.Message
}
