terraform {
  source = "my-terraform-module"
}

include "root" {
  path = find_in_parent_folders()
}

inputs = {
  instance_type = "t2.medium"

  min_size = 3
  max_size = 3
}
