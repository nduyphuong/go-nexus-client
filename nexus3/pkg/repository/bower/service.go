package bower

import (
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/client"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	repositoryBowerAPIEndpoint = common.RepositoryAPIEndpoint + "/bower"
)
type RepositoryBowerService struct {
	client *client.Client

	Group  *RepositoryBowerGroupService
	Hosted *RepositoryBowerHostedService
	Proxy  *RepositoryBowerProxyService
}

func NewRepositoryBowerService(c *client.Client) *RepositoryBowerService {
	return &RepositoryBowerService{
		client: c,

		Group:  NewRepositoryBowerGroupService(c),
		Hosted: NewRepositoryBowerHostedService(c),
		Proxy:  NewRepositoryBowerProxyService(c),
	}
}
