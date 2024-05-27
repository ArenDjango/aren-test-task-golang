package logger

type Config struct {
	Level                       string `validate:"required"`
	SkipFrameCount              int
	Structured                  bool
	ShowUnknownErrorsInResponse bool
	SecureReqJsonPaths          []string `validate:"required"`
	SecureResJsonPaths          []string `validate:"required"`
	MaxHTTPBodySize             int      `validate:"required"`
}
