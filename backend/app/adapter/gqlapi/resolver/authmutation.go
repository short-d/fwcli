package resolver

// AuthMutation represents GraphQL mutation resolver that acts differently based
// on the identify of the user
type AuthMutation struct {
	authToken *string
}

func newAuthMutation(
	authToken *string,
) AuthMutation {
	return AuthMutation{
		authToken: authToken,
	}
}
