package err

// AppErrCode defines the type of Application error code
// Application error codes are made of 4 digits.
// The first 2 digits denotes the component while
// the last 2 digits is the error index
type AppErrCode int

// 10: common
const (
	ErrUnknown AppErrCode = iota + 1000
)

// 11: model
const (
	ErrValidation AppErrCode = iota + 1100
)

// 12: database
const (
	ErrDatabase AppErrCode = iota + 1200
)

// 13: business logic
const (
	ErrTimerNotFound AppErrCode = iota + 1300
	ErrTimerAlreadyExists
)
