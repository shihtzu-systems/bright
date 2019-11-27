// Code generated; DANGER ZONE FOR EDITS

module unlock_count_mapping_definition {
  source = "./definition"

  target = var.target

  name = "unlock Count Mapping"
  structs_content = join("\n", [
    module.unlock_count_mapping_unlock_count_mapping_definition.rendered,
  ])

}

module unlock_count_mapping_unlock_count_mapping_definition {
  source = "./struct"

  name = "unlock Count Mapping Definition"
  attributes_content = join("\n", [
    module.unlock_count_mapping_unlock_count_mapping_definition_blacklisted.rendered,
    module.unlock_count_mapping_unlock_count_mapping_definition_hash.rendered,
    module.unlock_count_mapping_unlock_count_mapping_definition_index.rendered,
    module.unlock_count_mapping_unlock_count_mapping_definition_redacted.rendered,
    module.unlock_count_mapping_unlock_count_mapping_definition_unlock_value_hash.rendered,
  ])
}

module unlock_count_mapping_unlock_count_mapping_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module unlock_count_mapping_unlock_count_mapping_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module unlock_count_mapping_unlock_count_mapping_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module unlock_count_mapping_unlock_count_mapping_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module unlock_count_mapping_unlock_count_mapping_definition_unlock_value_hash {
  source = "./attribute"

  name = "unlock Value Hash"
  type = "int"
}


