// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	pipeline "github.com/jgautheron/provider-signoz/internal/controller/namespaced/log/pipeline"
	channel "github.com/jgautheron/provider-signoz/internal/controller/namespaced/notification/channel"
	providerconfig "github.com/jgautheron/provider-signoz/internal/controller/namespaced/providerconfig"
	view "github.com/jgautheron/provider-signoz/internal/controller/namespaced/saved/view"
	alert "github.com/jgautheron/provider-signoz/internal/controller/namespaced/signoz/alert"
	dashboard "github.com/jgautheron/provider-signoz/internal/controller/namespaced/signoz/dashboard"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		pipeline.Setup,
		channel.Setup,
		providerconfig.Setup,
		view.Setup,
		alert.Setup,
		dashboard.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		pipeline.SetupGated,
		channel.SetupGated,
		providerconfig.SetupGated,
		view.SetupGated,
		alert.SetupGated,
		dashboard.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
