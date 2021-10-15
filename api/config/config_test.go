package config_test

import (
	"gnt-cc/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFuncReturnsLoadedConfig(t *testing.T) {
	config.Parse("../testfiles/config.default.test.yaml")
	assert.EqualValues(t, config.Config{
		JwtSigningKey:        "test",
		AuthenticationMethod: "builtin",
		Bind:                 "127.0.0.1",
		PublicUrl:            "https://gnt-cc.example.com",
		Port:                 8000,
		JwtExpire:            60000000000,
		Loglevel:             "warning",
		Users: []config.UserConfig{{
			Username: "admin",
			Password: "test",
		}},
		Clusters: []config.ClusterConfig{{
			Name:        "test",
			Hostname:    "test-cluster.example.com",
			Port:        5080,
			Description: "Ganeti Test Cluster",
			Username:    "test",
			Password:    "supersecret",
			SSL:         true,
		}},
		LDAPConfig: config.LDAPConfig{},
	}, config.Get())
}

func TestGetClusterConfig(t *testing.T) {
	config.Parse("../testfiles/config.default.test.yaml")

	validClusterName := "test"

	cluster, err := config.GetClusterConfig(validClusterName)
	assert.Nil(t, err)
	assert.Equal(t, validClusterName, cluster.Name)

	invalidClusterName := "invalid-cluster"
	cluster, err = config.GetClusterConfig(invalidClusterName)
	assert.NotNil(t, err)
}

func TestIsValidCluster(t *testing.T) {
	config.Parse("../testfiles/config.default.test.yaml")

	validClusterName := "test"
	assert.True(t, config.ClusterExists(validClusterName))

	invalidClusterName := "invalid-cluster"
	assert.False(t, config.ClusterExists(invalidClusterName))
}

func TestParsingConfigShouldPanicIfAuthMethodIsSetToBuiltInWithMissingUsers(t *testing.T) {
	assert.PanicsWithError(t, "Authentication Method has been set to 'builtin' but no user is specified", func() {
		config.Parse("../testfiles/config.missing-users.test.yaml")
	})
}

func TestParsingConfigShouldPanicIfAuthMethodIsSetTOLDAPAndNoLDAPConfigIsPresent(t *testing.T) {
	assert.PanicsWithError(t, "Authentication Method has been set to 'ldap' but no LDAP host is specified", func() {
		config.Parse("../testfiles/config.invalid-ldap.test.yaml")
	})
}

func TestParsingConfigShouldPanicWithAnInvalidAuthMethod(t *testing.T) {
	assert.PanicsWithError(t, "Invalid authentication method 'test'", func() {
		config.Parse("../testfiles/config.invalid-auth-method.test.yaml")
	})
}

func TestParsingConfigShouldPanicWithEmptyConfig(t *testing.T) {
	assert.PanicsWithError(t, "No JWT signing key is set", func() {
		config.Parse("../testfiles/config.empty.test.yaml")
	})
}
