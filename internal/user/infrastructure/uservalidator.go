package infrastructure

type IValidatorRepository interface {
	CreateUser(phoneNumber string) (bool, error)
}

type UserValidator struct {
}

func NewValidator() IValidatorRepository {
	return UserValidator{}
}

func (r UserValidator) CreateUser(phoneNumber string) (bool, error) {
	//TODO implement me

	panic("implement me")
}
