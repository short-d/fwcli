package entity

type Secret struct {
	Key                    string
	Value                  string
	AvailableInPullRequest bool
}
