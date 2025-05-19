variable "aws_region" {
  type        = string
  description = "AWS region to deploy"
  default     = "sa-east-1"
}

variable "environment" {
  type    = string
  default = "development"
}

variable "project_name" {
  type        = string
  description = "Project name"
  default     = "course"
}
