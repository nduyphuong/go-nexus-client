package repository

import (
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/client"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/apt"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/bower"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/cocoapods"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/conan"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/conda"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/docker"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/gitlfs"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/golang"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/helm"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/legacy"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/maven"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/npm"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/nuget"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/p2"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/pypi"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/r"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/raw"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/rubygems"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/yum"
	"github.com/nduyphuong/go-nexus-client/nexus3/schema/repository"
)

type RepositoryService struct {
	client *client.Client

	// API Services
	Apt       *apt.RepositoryAptService
	Bower     *bower.RepositoryBowerService
	Cocoapods *cocoapods.RepositoryCocoapodsService
	Conan     *conan.RepositoryConanService
	Conda     *conda.RepositoryCondaService
	Docker    *docker.RepositoryDockerService
	GitLfs    *gitlfs.RepositoryGitLfsService
	Go        *golang.RepositoryGoService
	Helm      *helm.RepositoryHelmService
	Maven     *maven.RepositoryMavenService
	Npm       *npm.RepositoryNpmService
	Nuget     *nuget.RepositoryNugetService
	P2        *p2.RepositoryP2Service
	Pypi      *pypi.RepositoryPypiService
	R         *r.RepositoryRService
	Raw       *raw.RepositoryRawService
	RubyGems  *rubygems.RepositoryRubyGemsService
	Yum       *yum.RepositoryYumService
	Legacy    *legacy.RepositoryLegacyService
}

func NewRepositoryService(c *client.Client) *RepositoryService {
	return &RepositoryService{
		client: c,

		Apt:       apt.NewRepositoryAptService(c),
		Bower:     bower.NewRepositoryBowerService(c),
		Cocoapods: cocoapods.NewRepositoryCocoapodsService(c),
		Conan:     conan.NewRepositoryConanService(c),
		Conda:     conda.NewRepositoryCondaService(c),
		Docker:    docker.NewRepositoryDockerService(c),
		GitLfs:    gitlfs.NewRepositoryGitLfsService(c),
		Go:        golang.NewRepositoryGoService(c),
		Helm:      helm.NewRepositoryHelmService(c),
		Maven:     maven.NewRepositoryMavenService(c),
		Npm:       npm.NewRepositoryNpmService(c),
		Nuget:     nuget.NewRepositoryNugetService(c),
		P2:        p2.NewRepositoryP2Service(c),
		Pypi:      pypi.NewRepositoryPypiService(c),
		R:         r.NewRepositoryRService(c),
		Raw:       raw.NewRepositoryRawService(c),
		RubyGems:  rubygems.NewRepositoryRubyGemsService(c),
		Yum:       yum.NewRepositoryYumService(c),
		Legacy:    legacy.NewRepositoryLegacyService(c),
	}
}

func (s *RepositoryService) List() ([]repository.RepositoryInfo, error) {
	return common.ListRepositories(s.client)
}
