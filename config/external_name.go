/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/types/name"
)

// VersionV1Beta1 is used to signify that the resource has been tested and external name configured
const VersionV1Beta1 = "v1beta1"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"cloudflare_zone":                                   config.IdentifierFromProvider,
	"cloudflare_zone_settings_override":                 config.IdentifierFromProvider,
	"cloudflare_record":                                 config.IdentifierFromProvider,
	"cloudflare_authenticated_origin_pulls":             config.IdentifierFromProvider,
	"cloudflare_authenticated_origin_pulls_certificate": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

func externalNameConfig() config.ResourceOption {
	return func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	}
}

func defaultVersion() config.ResourceOption {
	return func(r *config.Resource) {
		r.Version = VersionV1Beta1
	}
}

func descriptionOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		config.ManipulateEveryField(r.TerraformResource, func(sch *schema.Schema) {
			sch.Description = ""
		})
	}
}

func groupOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = name.NewFromSnake(strings.TrimPrefix(r.Name, "cloudflare_")).Camel
	}
}
