resource "vault_generic_secret" "secret_user_password" {
  path = "secret/user/password"

  data_json = <<EOT
{
  "user_password":   "myS3cr3t"
}
EOT
}

