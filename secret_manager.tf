provider "aws" {
  version = "~> 3.0"
  region  = "ap-south-1"
}
variable "tags" {
    type = "map"
 default = {}
}
resource "aws_secretsmanager_secret" "example" {
  name = "redshift79"
  tags = var.tags
}

output "name" {
  value = aws_secretsmanager_secret.example.name
}

