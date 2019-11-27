// Code generated; DANGER ZONE FOR EDITS

module breaker_type_definition {
  source = "./definition"

  target = var.target

  name = "breaker Type"
  structs_content = join("\n", [
    module.breaker_type_breaker_type_definition.rendered,
  ])

}

module breaker_type_breaker_type_definition {
  source = "./struct"

  name = "breaker Type Definition"
  attributes_content = join("\n", [
    module.breaker_type_breaker_type_definition_blacklisted.rendered,
    module.breaker_type_breaker_type_definition_display_properties.rendered,
    module.breaker_type_breaker_type_definition_enum_value.rendered,
    module.breaker_type_breaker_type_definition_hash.rendered,
    module.breaker_type_breaker_type_definition_index.rendered,
    module.breaker_type_breaker_type_definition_redacted.rendered,
    module.breaker_type_breaker_type_definition_unlock_hash.rendered,
  ])
}

module breaker_type_breaker_type_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module breaker_type_breaker_type_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module breaker_type_breaker_type_definition_enum_value {
  source = "./attribute"

  name = "enum Value"
  type = "int"
}
module breaker_type_breaker_type_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module breaker_type_breaker_type_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module breaker_type_breaker_type_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module breaker_type_breaker_type_definition_unlock_hash {
  source = "./attribute"

  name = "unlock Hash"
  type = "int"
}


