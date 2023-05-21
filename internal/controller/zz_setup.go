/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	providerconfig "github.com/max4t/provider-cloudflare/internal/controller/providerconfig"
	authenticatedoriginpulls "github.com/max4t/provider-cloudflare/internal/controller/zone/authenticatedoriginpulls"
	authenticatedoriginpullscertificate "github.com/max4t/provider-cloudflare/internal/controller/zone/authenticatedoriginpullscertificate"
	record "github.com/max4t/provider-cloudflare/internal/controller/zone/record"
	settingsoverride "github.com/max4t/provider-cloudflare/internal/controller/zone/settingsoverride"
	zone "github.com/max4t/provider-cloudflare/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		authenticatedoriginpulls.Setup,
		authenticatedoriginpullscertificate.Setup,
		record.Setup,
		settingsoverride.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
