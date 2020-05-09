output "kubeconfig" {
  description = "kubectl config file contents for this kubernetes cluster"
  value       = digitalocean_kubernetes_cluster.primary.kube_config.0.raw_config
  sensitive   = true
}
