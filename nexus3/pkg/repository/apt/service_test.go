package apt_test

import (
	"testing"

	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/client"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/repository/apt"
	"github.com/nduyphuong/go-nexus-client/nexus3/pkg/tools"
	"github.com/stretchr/testify/assert"
)

var (
	testClient *client.Client = nil
)

func getTestClient() *client.Client {
	if testClient != nil {
		return testClient
	}
	return client.NewClient(getDefaultConfig())
}

func getTestService() *apt.RepositoryAptService {
	return apt.NewRepositoryAptService(getTestClient())
}

func getDefaultConfig() client.Config {
	return client.Config{
		Insecure: tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USRNAME", "admin").(string),
	}
}

func TestNewRepositoryService(t *testing.T) {
	s := getTestService()

	assert.NotNil(t, s, "NewRepositoryService() must not return nil")
}
