// Code generated; DANGER ZONE FOR EDITS

module platform_bucket_mapping_definition {
  source = "./definition"

  target = var.target

  name = "platform Bucket Mapping"
  structs_content = join("\n", [
    module.platform_bucket_mapping_platform_bucket_mapping_definition.rendered,
  ])

}

module platform_bucket_mapping_platform_bucket_mapping_definition {
  source = "./struct"

  name = "platform Bucket Mapping Definition"
  attributes_content = join("\n", [
    module.platform_bucket_mapping_platform_bucket_mapping_definition_blacklisted.rendered,
    module.platform_bucket_mapping_platform_bucket_mapping_definition_bucket_hash.rendered,
    module.platform_bucket_mapping_platform_bucket_mapping_definition_hash.rendered,
    module.platform_bucket_mapping_platform_bucket_mapping_definition_index.rendered,
    module.platform_bucket_mapping_platform_bucket_mapping_definition_membership_type.rendered,
    module.platform_bucket_mapping_platform_bucket_mapping_definition_redacted.rendered,
  ])
}

module platform_bucket_mapping_platform_bucket_mapping_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module platform_bucket_mapping_platform_bucket_mapping_definition_bucket_hash {
  source = "./attribute"

  name = "bucket Hash"
  type = "int"
}
module platform_bucket_mapping_platform_bucket_mapping_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module platform_bucket_mapping_platform_bucket_mapping_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module platform_bucket_mapping_platform_bucket_mapping_definition_membership_type {
  source = "./attribute"

  name = "membership Type"
  type = "int"
}
module platform_bucket_mapping_platform_bucket_mapping_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


