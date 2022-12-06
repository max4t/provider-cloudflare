/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/registry/reference"

	"github.com/max4t/provider-cloudflare/config/cloudflare"
)

const (
	resourcePrefix = "cloudflare"
	modulePath     = "github.com/max4t/provider-cloudflare"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithDefaultResourceOptions(
			groupOverrides(),
			externalNameConfig(),
			defaultVersion(),
			ExternalNameConfigurations(),
			// externalNameConfigurations(),
			descriptionOverrides(),
		),
		ujconfig.WithRootGroup("cloudflare.max4t.xyz"),
		ujconfig.WithShortName("cloudflare"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		cloudflare.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
