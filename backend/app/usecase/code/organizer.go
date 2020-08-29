package code

import "fwcli/app/entity"

type Organizer interface {
	GetRepoURL(namespace string, repoName string) (string, error)
	CreateRepo(namespace string, repoName string) (entity.Repo, error)
}
