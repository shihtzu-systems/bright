// Code generated; DANGER ZONE FOR EDITS

module unlock_expression_mapping_definition {
  source = "./definition"

  target = var.target

  name = "unlock Expression Mapping"
  structs_content = join("\n", [
    module.unlock_expression_mapping_unlock_expression_mapping_definition.rendered,
  ])

}

module unlock_expression_mapping_unlock_expression_mapping_definition {
  source = "./struct"

  name = "unlock Expression Mapping Definition"
  attributes_content = join("\n", [
    module.unlock_expression_mapping_unlock_expression_mapping_definition_blacklisted.rendered,
    module.unlock_expression_mapping_unlock_expression_mapping_definition_hash.rendered,
    module.unlock_expression_mapping_unlock_expression_mapping_definition_index.rendered,
    module.unlock_expression_mapping_unlock_expression_mapping_definition_redacted.rendered,
  ])
}

module unlock_expression_mapping_unlock_expression_mapping_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module unlock_expression_mapping_unlock_expression_mapping_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module unlock_expression_mapping_unlock_expression_mapping_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module unlock_expression_mapping_unlock_expression_mapping_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


