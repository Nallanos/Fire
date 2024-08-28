package errdefs

type Error struct {
	code int
	msg  string
}

var _ error = (*Error)(nil)

func (e Error) Error() string {
	return e.msg
}

func (e Error) HTTPCode() int {
	return e.code
}

func New(code int, msg string) error {
	return &Error{code: code, msg: msg}
}

var (
	ErrAlreadyExist = New(409, "already exist")
	ErrNotFound     = New(404, "not found")
	ErrInternal     = New(500, "internal server error")
	ErrBadRequest   = New(400, "bad request")
	ErrUnauthorized = New(401, "unauthorized")
	ErrForbidden    = New(403, "forbidden")
)
