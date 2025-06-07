package logger

type LogTags = string

const (
	UserIDLabel    LogTags = "user_id"
	ErrorLabel     LogTags = "error"
	RequestIDLabel LogTags = "req_id"
	TraceIDLabel   LogTags = "trace_id"
)
