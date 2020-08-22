package resolver

// Mutation represents GraphQL mutation resolver
type Mutation struct {
}

// AuthMutationArgs represents possible parameters for AuthMutation endpoint
type AuthMutationArgs struct {
	AuthToken       *string
	CaptchaResponse string
}

// AuthMutation extracts user information from authentication token
func (m Mutation) AuthMutation(args *AuthMutationArgs) (*AuthMutation, error) {
	authMutation := newAuthMutation(
		args.AuthToken,
	)
	return &authMutation, nil
}

func newMutation() Mutation {
	return Mutation{}
}
