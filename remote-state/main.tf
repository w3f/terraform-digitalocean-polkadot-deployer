provider "google" {
  project     = var.gcp_project_id
  credentials = var.gcp_credentials
}

resource "google_storage_bucket" "imagestore" {
  name          = "pd-tf-state-${var.deployment_name}"
  force_destroy = true
}
