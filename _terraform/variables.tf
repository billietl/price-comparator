
variable "gcp_project" {
    type = string
    description = "Selected gcp project"
    default = ""
}

variable "generate-sample-data" {
    type = bool
    description = "Make terraform generate sample data"
    default = false
}
