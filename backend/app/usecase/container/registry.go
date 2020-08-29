package container

import "fwcli/app/entity"

type Registry interface {
	GetNamespaceSecret() entity.Secret
	GetAuthSecrets() []entity.Secret
	CreateWebHook(namespace string, imageName string, webHook entity.WebHook) error
}
