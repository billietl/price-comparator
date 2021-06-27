
plan-dev:
	terraform workspace select dev
	terraform plan -var-file "tfvars/default.tfvars" -var-file "tfvars/dev.tfvars"

plan-prod:
	terraform workspace select prod
	terraform plan -var-file "tfvars/default.tfvars" -var-file "tfvars/prod.tfvars"

apply-dev:
	terraform workspace select dev
	terraform apply -var-file "tfvars/default.tfvars" -var-file "tfvars/dev.tfvars"

apply-prod:
	terraform workspace select prod
	terraform apply -var-file "tfvars/default.tfvars" -var-file "tfvars/prod.tfvars"

destroy-dev:
	terraform workspace select dev
	terraform destroy -var-file "tfvars/default.tfvars" -var-file "tfvars/dev.tfvars"

destroy-prod:
	terraform workspace select prod
	terraform destroy -var-file "tfvars/default.tfvars" -var-file "tfvars/prod.tfvars"
