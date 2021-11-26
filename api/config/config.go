package config

import (
	"errors"
	"fmt"
	"gnt-cc/utils"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ClusterConfig struct {
	Name        string
	Hostname    string
	Port        int
	Description string
	Username    string
	Password    string
	SSL         bool
}

type UserConfig struct {
	Username string
	Password string
}

type LDAPConfig struct {
	Host                  string
	Port                  int
	SkipCertificateVerify bool
	BaseDN                string
	UserFilter            string
	GroupFilter           string
}

type RapiConfig struct {
	SkipCertificateVerify bool
}

type Config struct {
	Bind                 string
	Port                 int
	DevelopmentMode      bool
	PublicUrl            string
	JwtSigningKey        string
	JwtExpire            time.Duration
	AuthenticationMethod string
	Loglevel             string
	Users                []UserConfig
	Clusters             []ClusterConfig
	LDAPConfig           LDAPConfig
	RapiConfig           RapiConfig
}

const (
	AuthMethodBuiltin = "builtin"
	AuthMethodLDAP    = "ldap"

	ConfigFileEnv = "GNT_CC_CONFIG"
)

var defaultConfigPaths = [...]string{
	"./config.yaml",
	"/etc/gnt-cc/config.yaml",
}

var c Config

func ClusterExists(clusterName string) bool {
	for _, cluster := range c.Clusters {
		if cluster.Name == clusterName {
			return true
		}
	}
	return false
}

func Get() Config {
	return c
}

func Init() {
	path, err := getConfigPath()

	if err != nil {
		panic(err)
	}

	log.Info("Using config file ", path)

	Parse(path)
}

func Parse(configPath string) {
	var config Config

	viper.SetConfigFile(configPath)

	viper.SetDefault("bind", "127.0.0.1")
	viper.SetDefault("port", "8000")
	viper.SetDefault("developmentMode", "false")
	viper.SetDefault("jwtExpire", "1m")
	viper.SetDefault("AuthenticationMethod", "builtin")
	viper.SetDefault("logLevel", "warning")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	log.SetLevel(parseLogLevel(config.Loglevel))

	err = validateConfig(&config)

	if err != nil {
		panic(err)
	}

	c = config
}

type clusterNotFoundError struct {
	clusterName string
}

func (e *clusterNotFoundError) Error() string {
	return fmt.Sprintf("Cluster '%s' not found", e.clusterName)
}

func GetClusterConfig(clusterName string) (ClusterConfig, *clusterNotFoundError) {
	for _, cluster := range c.Clusters {
		if cluster.Name == clusterName {
			return cluster, nil
		}
	}

	return ClusterConfig{}, &clusterNotFoundError{clusterName: clusterName}
}

func getConfigDefaultPath() (string, error) {
	for _, path := range defaultConfigPaths {
		if utils.FileExists(path) {
			return path, nil
		}
	}

	looked := strings.Join(defaultConfigPaths[:], "\n")

	return "", fmt.Errorf("cannot find a config file in the following places:\n\n%s", looked)
}

func getConfigPath() (string, error) {
	overridePath, overrideSet := os.LookupEnv(ConfigFileEnv)

	if overrideSet {
		if !utils.FileExists(overridePath) {
			return "", fmt.Errorf("config override specified, but cannot be found: %s", overridePath)
		}

		return overridePath, nil
	}

	return getConfigDefaultPath()
}

func validateConfig(config *Config) error {
	if config.JwtSigningKey == "" {
		return errors.New("No JWT signing key is set")
	}

	if config.PublicUrl == "" {
		return errors.New("no public URL specified")
	}

	switch config.AuthenticationMethod {
	case AuthMethodBuiltin:
		if len(config.Users) == 0 {
			return errors.New("Authentication Method has been set to 'builtin' but no user is specified")
		}
	case AuthMethodLDAP:
		if config.LDAPConfig.Host == "" {
			return errors.New("Authentication Method has been set to 'ldap' but no LDAP host is specified")
		}
	default:
		return errors.New(fmt.Sprintf("Invalid authentication method '%s'", config.AuthenticationMethod))
	}

	if len(config.Clusters) == 0 {
		return errors.New("No Ganeti clusters have been specified in the configuration file")
	}

	return nil
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
	log.Warnf("Invalid loglevel given: '%s', falling back to 'warning'", logLevel)
	return log.WarnLevel
}
