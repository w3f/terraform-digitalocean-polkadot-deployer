variable "do_token" {
  description = "Digital Ocean API token"
  type        = string
}

variable "cluster_name" {
  description = "Name of the Kubernetes cluster"
  default     = "polkadot-deployer"
  type        = string
}

variable "deployment_name" {
  description = "Name of this polkadot deployment"
  default     = "polkadot"
  type        = string
}

variable "location" {
  description = "Digital Ocean region"
  type        = string
}

variable "node_count" {
  description = "Size of EKS cluster"
  default     = 2
  type        = number
}

variable "machine_type" {
  description = "Type of Droplet instances used for the cluster"
  default     = "s-4vcpu-8gb"
  type        = string
}

variable "k8s_version" {
  description = "Kubernetes version to use for the cluster"
  default     = "1.14.3-do.0"
  type        = string
}

variable "gcp_project_id" {
  description = "Google cloud project id used for terraform state"
  type        = string
}

variable "gcp_credentials" {
  description = "Either the path to or the contents of a GCP service account key file in JSON format"
  type        = string
}
