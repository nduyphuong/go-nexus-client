package maven

import (
	"testing"

	"github.com/nduyphuong/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestMavenGroupRepository(name string) repository.MavenGroupRepository {
	return repository.MavenGroupRepository{
		Name:   name,
		Online: true,

		Group: repository.Group{
			MemberNames: []string{},
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
	}
}

func TestMavenGroupRepository(t *testing.T) {
	service := getTestService()
	repo := getTestMavenGroupRepository("test-maven-repo-group")

	testProxyRepo := getTestMavenProxyRepository("test-maven-group-proxy")
	defer service.Proxy.Delete(testProxyRepo.Name)
	err := service.Proxy.Create(testProxyRepo)
	assert.Nil(t, err)
	repo.Group.MemberNames = append(repo.Group.MemberNames, testProxyRepo.Name)

	err = service.Group.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Group.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Group, generatedRepo.Group)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Group.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Group.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	service.Group.Delete(repo.Name)
	assert.Nil(t, err)
}
