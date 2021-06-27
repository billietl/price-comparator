# resource "google_firestore_index" "my-index" {
#   project = var.gcp_project
#   collection = "store"

#   fields {
#     field_path = "name"
#     order      = "ASCENDING"
#   }

#   fields {
#     field_path = "city"
#     order      = "ASCENDING"
#   }

#   fields {
#     field_path = "zipcode"
#     order      = "ASCENDING"
#   }

# }
