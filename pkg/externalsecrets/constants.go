package externalsecrets

const (
	// ChartPullSecretPrefix is the prefix with which the registry pull secret
	// will be prefixed in the tenant namespace in the platform cluster
	ChartPullSecretPrefix = "sp-eso-"
	// DefaultNamespace is the default namespace in which the external-secrets-operator will be deployed in the mcp
	DefaultNamespace = "external-secrets"
	// HelmReleaseName is the name of the External Secrets Operator HelmRelease resource
	HelmReleaseName = "external-secrets"
	// OCIRepositoryName is the name of the External Secrets Operator OCIRepository resource
	OCIRepositoryName = "external-secrets"
)
