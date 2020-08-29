package ci

import "fwcli/app/entity"

type CI interface {
	EnableRepo(namespace string, name string) error
	UpdateRepo(namespace string, repo entity.Repo) error
	CreateSecret(namespace string, repoName string, secret entity.Secret)
}
