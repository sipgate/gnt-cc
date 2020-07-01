package config_test

import (
	"gnt-cc/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestIsvalidCluster(t *testing.T) {
	config.Parse("../testfiles/config.default.test.yaml")

	validClusterName := "test"
	assert.True(t, config.ClusterExists(validClusterName))

	invalidClusterName := "invalid-cluster"
	assert.False(t, config.ClusterExists(invalidClusterName))
}

func TestConfig_it_should_panic_if_auth_method_is_set_to_builtin_and_users_are_missing(t *testing.T) {
	assert.PanicsWithError(t, "Authentication Method has been set to 'builtin' but no user is specified", func() {
		config.Parse("../testfiles/config.missing-users.test.yaml")
	})
}

func TestConfig_it_should_panic_with_an_invalid_auth_method(t *testing.T) {
	assert.PanicsWithError(t, "Invalid authentication method 'test'", func() {
		config.Parse("../testfiles/config.invalid-auth-method.test.yaml")
	})
}

func TestConfig_it_should_panic_with_an_empty_config(t *testing.T) {
	assert.PanicsWithError(t, "No JWT signing key is set", func() {
		config.Parse("../testfiles/config.empty.test.yaml")
	})
}

func TestConfig_it_should_panic_if_auth_method_is_set_to_ldap_without_ldap_config(t *testing.T) {
	assert.PanicsWithError(t, "Authentication Method has been set to 'ldap' but no LDAP host is specified", func() {
		config.Parse("../testfiles/config.invalid-ldap.test.yaml")
	})
}
