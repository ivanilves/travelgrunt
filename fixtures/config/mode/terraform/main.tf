resource "aws_vpc" "default" {
  cidr_block = var.vpc_cidr
  tags = merge(
    local.common-tags,
    map(
      "Description", "VPC for creating resources",
    )
  )
}
