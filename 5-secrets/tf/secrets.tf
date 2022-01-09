resource "random_string" "database_password" {
  length           = 16
  special          = true
  override_special = "!&@$"
}

resource "vault_generic_secret" "secret_database" {
  path = "secret/database/credentials"

  data_json = <<EOT
{
  "password":  "${random_string.database_password.result}"
}
EOT
}

