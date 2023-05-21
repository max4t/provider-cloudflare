/*
Copyright 2021 Upbound Inc.
*/

package cloudflare

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const zoneGroup = "zone"

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cloudflare_zone", func(r *ujconfig.Resource) {
		r.ShortGroup = zoneGroup
	})
	p.AddResourceConfigurator("cloudflare_zone_settings_override", func(r *ujconfig.Resource) {
		r.ShortGroup = zoneGroup
		r.Kind = "SettingsOverride"
		r.References["zone_id"] = ujconfig.Reference{
			Type:              "Zone",
			RefFieldName:      "ZoneRef",
			SelectorFieldName: "ZoneSelector",
		}
	})
	p.AddResourceConfigurator("cloudflare_record", func(r *ujconfig.Resource) {
		r.ShortGroup = zoneGroup
		r.References["zone_id"] = ujconfig.Reference{
			Type:              "Zone",
			RefFieldName:      "ZoneRef",
			SelectorFieldName: "ZoneSelector",
		}
	})
	p.AddResourceConfigurator("cloudflare_authenticated_origin_pulls", func(r *ujconfig.Resource) {
		r.ShortGroup = zoneGroup
		r.Kind = "AuthenticatedOriginPulls"
		r.References["zone_id"] = ujconfig.Reference{
			Type:              "Zone",
			RefFieldName:      "ZoneRef",
			SelectorFieldName: "ZoneSelector",
		}
		r.References["authenticated_origin_pulls_certificate"] = ujconfig.Reference{
			Type:              "AuthenticatedOriginPullsCertificate",
			RefFieldName:      "AuthenticatedOriginPullsCertificateRef",
			SelectorFieldName: "AuthenticatedOriginPullsCertificateSelector",
		}
	})
	p.AddResourceConfigurator("cloudflare_authenticated_origin_pulls_certificate", func(r *ujconfig.Resource) {
		r.ShortGroup = zoneGroup
		r.Kind = "AuthenticatedOriginPullsCertificate"
		r.References["zone_id"] = ujconfig.Reference{
			Type:              "Zone",
			RefFieldName:      "ZoneRef",
			SelectorFieldName: "ZoneSelector",
		}
		r.TerraformResource.Schema["certificate"].Sensitive = true
		r.TerraformResource.Schema["private_key"].Sensitive = true
	})
}
