/*
Copyright 2021 Upbound Inc.
*/

package cloudflare

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cloudflare_zone", func(r *ujconfig.Resource) {
		r.ShortGroup = "zone"
	})
	p.AddResourceConfigurator("cloudflare_record", func(r *ujconfig.Resource) {
		r.ShortGroup = "zone"
		r.References["zone_id"] = ujconfig.Reference{
			Type:              "Zone",
			RefFieldName:      "ZoneRef",
			SelectorFieldName: "ZoneSelector",
		}
	})
}
