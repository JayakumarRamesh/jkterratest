provider "aws" {
  version = "~> 3.0"
  region  = "ap-south-1"
}
#resource "aws_instance" "web" {
#  ami           = "ami-04893cdb768d0f9ee"
#  instance_type = "t2.micro"

#  tags = {
#    Name = "HelloWorld"
#  }
#}

#resource "aws_secretsmanager_secret" "example" {
#  name = "example"
#}


variable "example" {
  default = {
    "username" : "awsuser",
    "password" : "***",
    "engine" : "redshift",
    "host" : "redshift-cluster-1.***.ap-south-1.redshift.amazonaws.com",
    "port" : 5439,
    "dbClusterIdentifier" : "redshift-cluster-1"
  }

  type = map(string)
}

module "secret_manager" {
  source = "./modules/secretmanager"
  tags   = var.tags
}

variable "tags" {
  type = "map"
  default = {}
}

resource "aws_secretsmanager_secret_version" "example" {
  secret_id     = module.secret_manager.secret_id
  secret_string = jsonencode(var.example)
}

output "test" {
  value = module.secret_manager.name
}

