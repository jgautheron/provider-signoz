package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider. Every SigNoz resource's identifier is assigned by the server (a
// UUID for dashboards/alerts/channels/views; a fixed id for the singleton log
// pipeline set), so they all use IdentifierFromProvider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"signoz_dashboard":            config.IdentifierFromProvider,
	"signoz_alert":                config.IdentifierFromProvider,
	"signoz_notification_channel": config.IdentifierFromProvider,
	"signoz_saved_view":           config.IdentifierFromProvider,
	"signoz_log_pipeline":         config.IdentifierFromProvider,
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
