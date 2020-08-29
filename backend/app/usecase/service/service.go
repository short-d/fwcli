package service

import (
	"fmt"
	"fwcli/app/entity"
	"fwcli/app/usecase/ci"
	"fwcli/app/usecase/code"
	"fwcli/app/usecase/container"
	"fwcli/app/usecase/dns"
	"fwcli/app/usecase/system"
	"path/filepath"
	"sync"
)

type Config struct {
	repoOrgNamespace       string
	containerOrgNamespace  string
	stagingBranch          string
	productionBranch       string
	deploymentConfigDir    string
	imageUpdaterWebHookURL string
	rootDomain             string
	loadBalancerIP         string
}

type Controller struct {
	codeOrganizer     code.Organizer
	versionControl    code.VersionControl
	ci                ci.CI
	orchestrator      container.Orchestrator
	dns               dns.DNS
	containerRegistry container.Registry
	system            system.System
	mutex             sync.Mutex
	config            Config
}

func (c Controller) CreateService(serviceName string) (entity.Service, error) {
	repo, err := c.codeOrganizer.CreateRepo(c.config.repoOrgNamespace, serviceName)
	if err != nil {
		return entity.Service{}, err
	}
	return entity.Service{
		Repo: repo,
	}, nil
}

func (c *Controller) EnableService(serviceName string) error {
	err := c.ci.EnableRepo(c.config.repoOrgNamespace, serviceName)
	if err != nil {
		return err
	}

	repo := entity.Repo{
		Name:      serviceName,
		IsTrusted: true,
	}
	err = c.ci.UpdateRepo(c.config.repoOrgNamespace, repo)
	if err != nil {
		return err
	}

	secrets := c.containerRegistry.GetAuthSecrets()
	for _, secret := range secrets {
		c.ci.CreateSecret(c.config.repoOrgNamespace, serviceName, secret)
	}
	secret := c.containerRegistry.GetNamespaceSecret()
	c.ci.CreateSecret(c.config.repoOrgNamespace, serviceName, secret)

	repoURL, err := c.codeOrganizer.GetRepoURL(c.config.repoOrgNamespace, serviceName)
	if err != nil {
		return err
	}

	repoDir := filepath.Join("repo", c.config.repoOrgNamespace, serviceName)

	err = c.versionControl.CloneRepo(repoURL, repoDir)
	if err != nil {
		return err
	}

	c.mutex.Lock()

	err = c.system.ChangeDirectory(repoDir)
	if err != nil {
		return err
	}

	err = c.versionControl.CheckoutBranch(c.config.stagingBranch)
	if err != nil {
		return err
	}

	configDir := filepath.Join(repoDir, c.config.deploymentConfigDir)
	err = c.orchestrator.ApplyConfig(configDir)
	if err != nil {
		return err
	}

	c.mutex.Unlock()

	err = c.dns.CreateARecord(serviceName, c.config.rootDomain, c.config.loadBalancerIP)
	err = c.orchestrator.ApplyConfig(configDir)
	if err != nil {
		return err
	}

	stagingImageName := fmt.Sprintf("%s_staging", serviceName)
	webHook := entity.WebHook{
		Name: "Image Updater",
		URL:  c.config.imageUpdaterWebHookURL,
	}
	err = c.containerRegistry.CreateWebHook(
		c.config.containerOrgNamespace,
		stagingImageName,
		webHook,
	)
	if err != nil {
		return err
	}

	return c.containerRegistry.CreateWebHook(
		c.config.containerOrgNamespace,
		serviceName,
		webHook,
	)
}
