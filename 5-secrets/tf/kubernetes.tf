data "kubernetes_service_account" "payments_service_account" {
  metadata {
    name = "payments"
    namespace = "filhodanuvem-services"
  }
}

data "kubernetes_secret" "payments_service_account_jwt_secret" {
  metadata {
    name = "${data.kubernetes_service_account.payments_service_account.default_secret_name}"
    namespace = "filhodanuvem-services"
  }
}

resource "vault_auth_backend" "kubernetes" {
  type = "kubernetes"
}

resource "vault_kubernetes_auth_backend_config" "example" {
  backend                = vault_auth_backend.kubernetes.path
  kubernetes_host        = "https://kubernetes.default.svc"
  token_reviewer_jwt     = data.kubernetes_secret.payments_service_account_jwt_secret.data.token
}

resource "vault_policy" "policy_user_password" {
  name = "dev-team"

  policy = <<EOT
path "secret/data/user/password" {
  capabilities = ["read"]
}
EOT
}

resource "vault_kubernetes_auth_backend_role" "app" {
  backend                          = vault_auth_backend.kubernetes.path
  role_name                        = "app"
  bound_service_account_names      = ["payments"]
  bound_service_account_namespaces = ["*"]
  token_policies =  [vault_policy.policy_user_password.name]
  token_ttl = 36000
}


