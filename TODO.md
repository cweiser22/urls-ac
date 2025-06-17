# TODO

- [x] Set up database migrations
- [x] Set up Prometheus and Grafana for monitoring
- [ ] Set up K8s
- [ ] Proper env configs
- [ ] Proper logging
- [ ] Distributed Redis
- [ ] Add more Prometheus metrics
- [ ] Use standard Go project layout
  env:
  PROJECT_ID: 'urls-ac' # TODO: update to your Google Cloud project ID
  GAR_LOCATION: 'us-west1' # TODO: update to your region
  GKE_CLUSTER: 'urls-ac-cluster-1' # TODO: update to your cluster name
  GKE_ZONE: 'us-west1' # TODO: update to your cluster zone
  DEPLOYMENT_NAME: 'urls-ac-staging-deployment' # TODO: update to your deployment name
  REPOSITORY: 'urls-ac-container-registry' # TODO: update to your Artifact Registry docker repository name
  IMAGE: 'app'
  WORKLOAD_IDENTITY_PROVIDER: 'projects/315102483322/locations/global/workloadIdentityPools/urls-ac-gke-deploy/providers/github-actions-provider' # TODO: update to your workload identity provider
