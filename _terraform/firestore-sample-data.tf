
resource "google_firestore_document" "store_1" {
  count = var.generate-sample-data ? 1 : 0
  project = var.gcp_project
  collection = "store"
  document_id = "b13ee4d6-d1b2-403e-bd3d-3779760fa2a9"
  fields      = <<EOT
  {
      "name":{
          "stringValue": "Day by day"
      },
      "city":{
          "stringValue": "Croix"
      },
      "zipcode":{
          "stringValue": "59170"
      }
  }
  EOT
}

resource "google_firestore_document" "store_2" {
  count = var.generate-sample-data ? 1 : 0
  project = var.gcp_project
  collection = "store"
  document_id = "3a2444a2-2657-4fd2-a90e-d2a92cdb895a"
  fields      = <<EOT
  {
      "name":{
          "stringValue": "Aldi"
      },
      "city":{
          "stringValue": "Croix"
      },
      "zipcode":{
          "stringValue": "59170"
      }
  }
  EOT
}

resource "google_firestore_document" "store_3" {
  count = var.generate-sample-data ? 1 : 0
  project = var.gcp_project
  collection = "store"
  document_id = "5c690c91-2518-4044-a63b-f95de3252fb0"
  fields      = <<EOT
  {
      "name":{
          "stringValue": "Intermarché"
      },
      "city":{
          "stringValue": "Croix"
      },
      "zipcode":{
          "stringValue": "59170"
      }
  }
  EOT
}

resource "google_firestore_document" "store_4" {
  count = var.generate-sample-data ? 1 : 0
  project = var.gcp_project
  collection = "store"
  document_id = "5c690c91-2518-4044-a63b-f95df3252fb0"
  fields      = <<EOT
  {
      "name":{
          "stringValue": "Monoprix"
      },
      "city":{
          "stringValue": "Lille"
      },
      "zipcode":{
          "stringValue": "59000"
      }
  }
  EOT
}

resource "google_firestore_document" "store_5" {
  count = var.generate-sample-data ? 1 : 0
  project = var.gcp_project
  collection = "store"
  document_id = "5c690c91-2518-4044-a63b-f95dd3252fb0"
  fields      = <<EOT
  {
      "name":{
          "stringValue": "Leclerc"
      },
      "city":{
          "stringValue": "Nantes"
      },
      "zipcode":{
          "stringValue": "44300"
      }
  }
  EOT
}

resource "google_firestore_document" "store_6" {
  count = var.generate-sample-data ? 1 : 0
  project = var.gcp_project
  collection = "store"
  document_id = "5c691c91-2518-4044-a63b-f95dd3252fb0"
  fields      = <<EOT
  {
      "name":{
          "stringValue": "Carrefour"
      },
      "city":{
          "stringValue": "Wasquehal"
      },
      "zipcode":{
          "stringValue": "59290"
      }
  }
  EOT
}
