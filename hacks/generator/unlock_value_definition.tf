// Code generated; DANGER ZONE FOR EDITS

module unlock_value_definition {
  source = "./definition"

  target = var.target

  name = "unlock Value"
  structs_content = join("\n", [
    module.unlock_value_unlock_value_definition.rendered,
  ])

}

module unlock_value_unlock_value_definition {
  source = "./struct"

  name = "unlock Value Definition"
  attributes_content = join("\n", [
    module.unlock_value_unlock_value_definition_aggregation_type.rendered,
    module.unlock_value_unlock_value_definition_blacklisted.rendered,
    module.unlock_value_unlock_value_definition_hash.rendered,
    module.unlock_value_unlock_value_definition_index.rendered,
    module.unlock_value_unlock_value_definition_mapping_index.rendered,
    module.unlock_value_unlock_value_definition_redacted.rendered,
    module.unlock_value_unlock_value_definition_scope.rendered,
  ])
}

module unlock_value_unlock_value_definition_aggregation_type {
  source = "./attribute"

  name = "aggregation Type"
  type = "int"
}
module unlock_value_unlock_value_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module unlock_value_unlock_value_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module unlock_value_unlock_value_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module unlock_value_unlock_value_definition_mapping_index {
  source = "./attribute"

  name = "mapping Index"
  type = "int"
}
module unlock_value_unlock_value_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module unlock_value_unlock_value_definition_scope {
  source = "./attribute"

  name = "scope"
  type = "int"
}


