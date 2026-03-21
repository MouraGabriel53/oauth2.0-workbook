package authrepository

func NewAuthenticationRepositoryInterface() *authenticationRepositoryInterface {
	return &authenticationRepositoryInterface{}
}

type AuthenticationRepositoryInterface interface {
	StoreVerifier(state, verifier string) (success bool, err error)
}

type authenticationRepositoryInterface struct{}
