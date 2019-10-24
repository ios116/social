package exceptions

type DomainError string

func (ee DomainError) Error() string {
	return string(ee)
}

var (
	Auth                = DomainError("Password or login is not valid")
	LoginRequired       = DomainError("Login required")
	FirstNameRequired       = DomainError("FirstName required")
	LastNameRequired       = DomainError("LastName required")
	EmailRequired       = DomainError("Email required")
	PasswordRequired    = DomainError("Password required")
	ObjectDoesNotExist  = DomainError("Object does not exist")
	InternalServerError = DomainError("Internal Server Error")
)
