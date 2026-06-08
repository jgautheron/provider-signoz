# Provider SigNoz

A [Crossplane](https://crossplane.io/) provider for self-hosted
[SigNoz](https://signoz.io), generated with [Upjet](https://github.com/crossplane/upjet)
from the [SigNoz Terraform provider](https://github.com/jgautheron/terraform-provider-signoz).

Manage SigNoz observability config as Kubernetes Managed Resources: dashboards,
alerts, notification channels, saved views, and log pipelines — reconciled by a
controller, GitOps-friendly.

> **Status: early (v0.x).** Targets SigNoz Community >= 0.125. Both
> cluster-scoped (`signoz.crossplane.io`) and namespaced (`signoz.m.crossplane.io`)
> Managed Resources are generated.

## Managed Resources

| Kind | API group | SigNoz resource |
| --- | --- | --- |
| `Dashboard` | `signoz.signoz.crossplane.io` | Dashboards (full definition as a `data` JSON string) |
| `Alert` | `signoz.signoz.crossplane.io` | Alert rules (v5/v2alpha1) |
| `Channel` | `notification.signoz.crossplane.io` | Notification channels (**admin token**) |
| `View` | `saved.signoz.crossplane.io` | Logs/traces explorer saved views |
| `Pipeline` | `log.signoz.crossplane.io` | Log-processing pipeline set (singleton) |

Complex bodies (alert condition, dashboard definition, …) are plain JSON
strings on `spec.forProvider`, mirroring the Terraform provider's churn-resistant
design.

## Usage

```yaml
# 1. Credentials Secret (SigNoz Service Account token; admin role for channels)
apiVersion: v1
kind: Secret
metadata:
  name: signoz-creds
  namespace: crossplane-system
stringData:
  credentials: |
    {"access_token":"<SIGNOZ-API-KEY>","endpoint":"https://signoz.example.com"}
---
# 2. ProviderConfig referencing it
apiVersion: signoz.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: signoz-creds
      key: credentials
---
# 3. A managed resource
apiVersion: signoz.signoz.crossplane.io/v1alpha1
kind: Dashboard
metadata:
  name: overview
spec:
  forProvider:
    data: |
      {"title":"App Overview","tags":["managed-by:crossplane"],"layout":[],"widgets":[]}
  providerConfigRef:
    name: default
```

## Install

```bash
crossplane xpkg install provider xpkg.upbound.io/jgautheron/provider-signoz:v0.1.0
```

## Developing

```bash
make generate              # regenerate CRDs + controllers from the TF provider schema
make build                 # build controller image + xpkg
make local-deploy          # kind + Crossplane + the locally-built provider
```

The runtime bundles Terraform + `jgautheron/signoz` as an offline provider
mirror; the controller reconciles each MR by running that provider.

## Related

- [jgautheron/terraform-provider-signoz](https://github.com/jgautheron/terraform-provider-signoz) — the wrapped Terraform provider
- [`@jooon/pulumi-signoz`](https://www.npmjs.com/package/@jooon/pulumi-signoz) — the Pulumi bridge over the same provider
- [Upjet](https://github.com/crossplane/upjet)

## License

Apache-2.0.
