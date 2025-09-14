package keycloak

import (
	"errors"
	"fmt"
	"net/url"
	"os"
)

const keycloakUrlKey string = "KeycloakUrl"
const keycloakRealmKey string = "KeycloakRealm"
const keycloakClientIdKey string = "KeycloakClientId"

type OicdConfiguration struct {
	keycloakUrl      *url.URL
	keycloakRealm    string
	keycloakClientId string
}

func NewOicdConfiguration() (*OicdConfiguration, error) {
	u, err := loadUrl()
	if err != nil {
		return nil, err
	}

	realm, err := loadRealm()
	if err != nil {
		return nil, err
	}

	clientId, err := loadClientId()
	if err != nil {
		return nil, err
	}

	return &OicdConfiguration{u, realm, clientId}, nil

}

func loadUrl() (*url.URL, error) {
	stringifiedUrl := os.Getenv(keycloakUrlKey)
	if stringifiedUrl == "" {
		return nil, errors.New(fmt.Sprintf("Could not load environment variable '%v'", keycloakUrlKey))
	}

	parsedUrl, err := url.ParseRequestURI(stringifiedUrl)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("The environment variable '%v' is not an url: '%v'", keycloakUrlKey, url))
	}

	return parsedUrl, nil
}

func loadRealm() (string, error) {
	realm := os.Getenv(keycloakRealmKey)
	if realm == "" {
		return "", errors.New(fmt.Sprintf("Could not load environment variable '%v'", keycloakRealmKey))
	}

	return realm, nil
}

func loadClientId() (string, error) {
	clientId := os.Getenv(keycloakClientIdKey)
	if clientId == "" {
		return "", errors.New(fmt.Sprintf("Could not load environment variable '%v'", keycloakClientIdKey))
	}

	return clientId, nil
}

func (configuration *OicdConfiguration) GetUrl() *url.URL {
	return configuration.keycloakUrl
}

func (configuration *OicdConfiguration) GetUrlStringified() string {
	return configuration.keycloakUrl.String()
}

func (configuration *OicdConfiguration) GetRealm() string {
	return configuration.keycloakRealm
}

func (configuration *OicdConfiguration) GetClientId() string {
	return configuration.keycloakClientId
}
