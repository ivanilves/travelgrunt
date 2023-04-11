terraform {
  source = "my-terraform-module"
}

include "root" {
  path = find_in_parent_folders()
}

inputs = {
  foo = "bar"
}
