package v1

// KMSConfig defines the configuration for the KMS instance
// that will be used with KMSv2 encryption
//
// +openshift:enable:FeatureGate=KMSv2
type KMSConfig struct {
	// aws defines the key config for using an AWS KMS instance
	// for the encryption. The AWS KMS instance is managed
	// by the user outside the purview of the control plane.
	//
	// +optional
	AWS *AWSKMSConfig `json:"aws,omitempty"`
}

// AWSKMSConfig defines the KMS config specific to AWS KMS provider
//
// +openshift:enable:FeatureGate=KMSv2
type AWSKMSConfig struct {
	// keyARN is the AWS ARN for the symmetric encryption KMS key
	//
	// +kubebuilder:validation:Pattern=`^arn:aws:kms:[a-z0-9-]+:[0-9]{12}:key/[a-f0-9-]+$`
	KeyARN string `json:"keyARN"`
	// region is the AWS region where the KMS instance exists
	//
	// +kubebuilder:validation:Pattern=`^[a-z]{2}-[a-z]+-\d{1,2}$`
	Region string `json:"region"`
}
