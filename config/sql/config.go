package compute

import (
	"context"
	"fmt"

	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint: gocyclo
	// Note(turkenh): We ignore gocyclo in this function since it configures
	//  all resources separately and no complex logic here.

	p.AddResourceConfigurator("google_sql_ssl_cert", func(r *config.Resource) {
		r.Kind = "SslCert"
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			project, err := common.GetField(providerConfig, common.KeyProject)
			if err != nil {
				return "", err
			}
			instance, err := common.GetField(parameters, "instance")
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("projects/%s/instances/%s/sslCerts/%s", project, instance, externalName), nil
		}
	})
}
