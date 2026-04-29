package domain

type AuthUsecase interface {
    Register(user User) error
    Login(noTelp, kataSandi string) (User, error)
}