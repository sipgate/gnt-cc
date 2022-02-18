package auth

import (
	"gnt-cc/config"

	"github.com/jtblin/go-ldap-client"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Username string
}

func validateUser(userID string, password string) bool {
	switch config.Get().AuthenticationMethod {
	case "builtin":
		return validateLocalUser(userID, password)
	case "ldap":
		return validateLdapUser(userID, password)
	}
	return false
}

func validateLocalUser(userID string, password string) bool {
	for _, el := range config.Get().Users {
		userSet := config.UserConfig(el)
		if userSet.Username == userID && userSet.Password == password {
			return true
		}
	}
	return false
}

func validateLdapUser(userID string, password string) bool {
	client := &ldap.LDAPClient{
		Base:               config.Get().LDAPConfig.BaseDN,
		Host:               config.Get().LDAPConfig.Host,
		ServerName:         config.Get().LDAPConfig.Host,
		Port:               config.Get().LDAPConfig.Port,
		InsecureSkipVerify: config.Get().LDAPConfig.SkipCertificateVerify,
		UserFilter:         config.Get().LDAPConfig.UserFilter,
		GroupFilter:        config.Get().LDAPConfig.GroupFilter,
	}
	defer client.Close()
	ok, _, err := client.Authenticate(userID, password)
	if err != nil {
		log.Errorf("Error authenticating user '%s': %+v", userID, err)
		return false
	}
	if !ok {
		log.Warningf("Authenticating failed for user '%s'", userID)
	}

	groups, err := client.GetGroupsOfUser(userID)
	if err != nil {
		log.Errorf("Error getting groups for user %s: %+v", "username", err)
	}

	if len(groups) > 0 {
		log.Debugf("Authentication for user '%s' successful", userID)
		return true
	}

	log.Warningf("User '%s' does not belong to LDAP groups matching the search filter", userID)
	return false
}
