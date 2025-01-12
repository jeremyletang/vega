package jsonrpc

// Result is just a nicer way to describe what's expected to be returned by the
// handlers.
type Result interface{}

type Response struct {
	// Version specifies the version of the JSON-RPC protocol.
	// MUST be exactly "2.0".
	Version string `json:"jsonrpc"`

	// Result is REQUIRED on success. This member MUST NOT exist if there was an
	// error invoking the method.
	Result Result `json:"result,omitempty"`

	// Error is REQUIRED on error. This member MUST NOT exist if there was no
	// error triggered during invocation.
	Error *ErrorDetails `json:"error,omitempty"`

	// ID is an identifier established by the Client that MUST contain a String.
	// This member is REQUIRED. It MUST be the same as the value of the id member
	// in the Request Object.
	// If there was an error in detecting the id in the Request object (e.g.
	// Parse error/Invalid Request), it MUST be empty.
	ID string `json:"id,omitempty"`
}

type ErrorCode int16

const (
	// ErrorCodeParseError Invalid JSON was received by the server. An error
	// occurred on the server while parsing the JSON text.
	ErrorCodeParseError ErrorCode = -32700
	// ErrorCodeInvalidRequest The JSON sent is not a valid Request object.
	ErrorCodeInvalidRequest ErrorCode = -32600
	// ErrorCodeMethodNotFound The method does not exist / is not available.
	ErrorCodeMethodNotFound ErrorCode = -32601
	// ErrorCodeInvalidParams Invalid method parameter(s).
	ErrorCodeInvalidParams ErrorCode = -32602
	// ErrorCodeInternalError Internal JSON-RPC error.
	ErrorCodeInternalError ErrorCode = -32603

	// Implementation-defined server-errors.
	// -32000 to -32099 codes are reserved for implementation-defined server-errors.
	// See https://www.jsonrpc.org/specification#error_object for more information.

	// ErrorCodeRequestAlreadyBeingProcessed is a custom server implementation
	// error indicating that a request is already being processed. The server
	// doesn't accept concurrent requests.
	ErrorCodeRequestAlreadyBeingProcessed ErrorCode = -32000
	// ErrorCodeRequestHasBeenInterrupted refers to a request that has been
	// interrupted by the server or the third-party application. It could
	// originate from a timeout or an explicit cancellation.
	ErrorCodeRequestHasBeenInterrupted ErrorCode = -32001
)

// ErrorDetails is returned when an RPC call encounters an error.
type ErrorDetails struct {
	// Code indicates the error type that occurred.
	Code ErrorCode `json:"code"`

	// Message provides a short description of the error.
	// The message SHOULD be limited to a concise single sentence.
	Message string `json:"message"`

	// Data is a primitive or a structured value that contains additional
	// information about the error. This may be omitted.
	// The value of this member is defined by the Server (e.g. detailed error
	// information, nested errors etc.).
	Data string `json:"data,omitempty"`
}

func (d ErrorDetails) IsInternalError() bool {
	return d.Message == "Internal error"
}

func NewParseError(data error) *ErrorDetails {
	return &ErrorDetails{
		Code:    ErrorCodeParseError,
		Message: "Parse error",
		Data:    data.Error(),
	}
}

func NewInvalidRequest(data error) *ErrorDetails {
	return &ErrorDetails{
		Code:    ErrorCodeInvalidRequest,
		Message: "Invalid Request",
		Data:    data.Error(),
	}
}

func NewMethodNotFound(data error) *ErrorDetails {
	return &ErrorDetails{
		Code:    ErrorCodeMethodNotFound,
		Message: "Method not found",
		Data:    data.Error(),
	}
}

func NewInvalidParams(data error) *ErrorDetails {
	return &ErrorDetails{
		Code:    ErrorCodeInvalidParams,
		Message: "Invalid params",
		Data:    data.Error(),
	}
}

func NewInternalError(data error) *ErrorDetails {
	return &ErrorDetails{
		Code:    ErrorCodeInternalError,
		Message: "Internal error",
		Data:    data.Error(),
	}
}

func NewServerError(code ErrorCode, data error) *ErrorDetails {
	if code > -32000 || code < -32099 {
		panic("server error code should be between [-32000, -32099]")
	}
	return &ErrorDetails{
		Code:    code,
		Message: "Server error",
		Data:    data.Error(),
	}
}

func NewCustomError(code ErrorCode, message string, data error) *ErrorDetails {
	if code <= -32000 {
		panic("custom error code should be greater than -32000")
	}
	return &ErrorDetails{
		Code:    code,
		Message: message,
		Data:    data.Error(),
	}
}

func NewErrorResponse(id string, details *ErrorDetails) *Response {
	return &Response{
		Version: JSONRPC2,
		Error:   details,
		ID:      id,
	}
}

func NewSuccessfulResponse(id string, result Result) *Response {
	return &Response{
		Version: JSONRPC2,
		Result:  result,
		ID:      id,
	}
}
