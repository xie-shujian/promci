package main

type Config struct {
	Gitlab            Gitlab `yaml:"gitlab"`
	PromciAccessToken string `yaml:"promci_access_token"`
}
type Repository struct {
	Name        string `yaml:"name"`
	AccessToken string `yaml:"access_token"`
	Directory   string `yaml:"directory"`
}
type Gitlab struct {
	GroupURL     string       `yaml:"group_url"`
	Repositories []Repository `yaml:"repositories"`
}

func (c *Config) BuildRepositoryAccessUrl(repository_name string) (string, string) {
	for _, repository := range c.Gitlab.Repositories {
		if repository.Name == repository_name {
			repositoryAccessUrl := "http://oauth2:" + repository.AccessToken + "@" + c.Gitlab.GroupURL + "/" + repository.Name + ".git"
			return repositoryAccessUrl, repository.Directory
		}
	}
	return "", ""
}

func (c *Config) ExistRepository(repository_name string) bool {
	for _, repository := range c.Gitlab.Repositories {
		if repository.Name == repository_name {
			return true
		}
	}
	return false
}
