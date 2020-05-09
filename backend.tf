provider "google" {
  project     = var.gcp_project_id
  credentials = var.gcp_credentials
}

terraform {
  backend "gcs" {
    bucket  = "pd-tf-state-${var.deployment_name}"
    prefix  = "terraform/state/${var.cluster_name}"
  }
}
