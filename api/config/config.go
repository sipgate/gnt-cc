package config

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type GanetiCluster struct {
	Name        string `mapstructure:"name"`
	Hostname    string `mapstructure:"hostname"`
	Port        int
	Description string
	Username    string
	Password    string
	SSL         bool
}

type UserSet struct {
	Username string
	Password string
}

type LdapConfig struct {
	Host                  string
	Port                  int
	SkipCertificateVerify bool
	BaseDN                string
	UserFilter            string
	GroupFilter           string
}

type Config struct {
	Bind                 string
	Port                 int
	DevelopmentMode      bool
	JwtSigningKey        string
	JwtExpire            string
	AuthenticationMethod string
	Loglevel             string
	DummyMode            bool
	Users                []UserSet
	Clusters             []GanetiCluster
	LDAPConfig           LdapConfig
}

var (
	ValidAuthMethods = []string{
		"builtin",
		"ldap",
	}

	c Config
)

func Get() Config {
	return c
}

func Parse() {
	viper.SetConfigName("config")
	alternateConfigDir := os.Getenv("GNT_CC_CONFIG_DIR")
	if alternateConfigDir != "" {
		viper.AddConfigPath(alternateConfigDir)
	} else {
		viper.AddConfigPath(".")
	}

	viper.SetDefault("bind", "127.0.0.1")
	viper.SetDefault("port", "8000")
	viper.SetDefault("developmentMode", "false")
	viper.SetDefault("jwtExpire", "1m")
	viper.SetDefault("AuthenticationMethod", "builtin")
	viper.SetDefault("logLevel", "warning")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&c)

	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	log.SetLevel(parseLogLevel(c.Loglevel))

	validateConfig()
}

func GetClusterConfig(clusterName string) GanetiCluster {
	for _, cluster := range c.Clusters {
		if cluster.Name == clusterName {
			return cluster
		}
	}
	// TODO: panic() is probably not a great reaction to querying a non-existant cluster
	panic(fmt.Sprintf("Could not find requested config for ganeti cluster '%s'", clusterName))
}

func validateConfig() bool {
	if !isInSlice(c.AuthenticationMethod, ValidAuthMethods) {
		panic(fmt.Sprintf("'%s' is not a valid Authentication Method (available methods: %v)", c.AuthenticationMethod, ValidAuthMethods))
	}

	switch c.AuthenticationMethod {
	case "builtin":
		if len(c.Users) == 0 {
			panic(fmt.Sprintf("Authentication Method has been set to 'builtin' but no users have been specified."))
		}
	}

	if len(c.Clusters) == 0 {
		panic(fmt.Sprintf("No Ganeti clusters have been specified in the configuration file."))
	}

	return true
}

func isInSlice(needle string, list []string) bool {
	for _, entry := range list {
		if entry == needle {
			return true
		}
	}
	return false
}

func parseLogLevel(logLevel string) log.Level {
	switch strings.ToLower(logLevel) {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warning":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	}
	log.Errorf("Invalid loglevel given: '%s', falling back to 'warning'", logLevel)
	return log.WarnLevel
}
