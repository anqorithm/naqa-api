variable "project_id" {
  type = string
}

variable "region" {
  type    = string
  default = "us-central1"
}

variable "service_name" {
  type    = string
  default = "naqa-api"
}

variable "mongo_uri" {
  type      = string
  sensitive = true
}

variable "mongo_database" {
  type    = string
  default = "naqa"
}