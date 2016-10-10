---
# Google Cloud Storage
bucket_access_key: {{bucket_access_key}} # GCS interop access key
bucket_secret_key: {{bucket_secret_key}} # GCS interop secret key
bucket_name:       {{bucket_name}} # GCS bucket for semver storage

# Google service account
service_account: {{service_account}}
service_account_key_json: |
  {{service_account_key_json}}

# BOSH and Cloud Foundry
bosh_director_address: {{bosh_director_address}} # IP address of BOSH director
bosh_user:             {{bosh_user}} # Bosh admin username
bosh_password:         {{bosh_password} # Bosh password
cf_deployment_name:    {{cf_deployment_name}} # Name of CF deployment to update

# GitHub
github_deployment_key_bosh_google_cpi_release: | # GitHub deployment key for release artifacts
github_pr_access_token: # An access token with repo:status access, used to test PRs
