package config

import (
	"strings"

	"github.com/ktr0731/toml"
	"go.uber.org/zap"
)

// ServiceConfig is the top-level for the service configuration file.
type ServiceConfig struct {
	Services map[string]ServiceDetails
}

// ServiceDetails holds the details for a single service definition.
type ServiceDetails struct {
	EndpointURL                 string   `toml:"endpoint-url"`
	AuthorizationModel          string   `toml:"authorization-model"`
	AuthorizationWhitelist      []string `toml:"authorization-whitelist"`
	DocumentationURL            string   `toml:"documentation-url"`              // Config parameter will be moved to directory admin interface
	APISpecificationType        string   `toml:"api-specification-type"`         // Config parameter will be moved to directory admin interface
	APISpecificationDocumentURL string   `toml:"api-specification-document-url"` // Config parameter will be moved to directory admin interface
}

// LoadServiceConfig reads the service config from disk and returns.
func LoadServiceConfig(logger *zap.Logger, serviceConfigLocation string) *ServiceConfig {
	serviceConfig := &ServiceConfig{}
	tomlMetaData, err := toml.DecodeFile(serviceConfigLocation, serviceConfig)
	if err != nil {
		logger.Fatal("failed to load service config", zap.Error(err))
	}
	if len(tomlMetaData.Undecoded()) > 0 {
		logger.Fatal("unsupported values in toml", zap.String("key", strings.Join(tomlMetaData.Undecoded()[0], ">")))
	}
	return serviceConfig
}
