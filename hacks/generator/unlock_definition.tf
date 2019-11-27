// Code generated; DANGER ZONE FOR EDITS

module unlock_definition {
  source = "./definition"

  target = var.target

  name = "unlock"
  structs_content = join("\n", [
    module.unlock_unlock_definition.rendered,
  ])

}

module unlock_unlock_definition {
  source = "./struct"

  name = "unlock Definition"
  attributes_content = join("\n", [
    module.unlock_unlock_definition_blacklisted.rendered,
    module.unlock_unlock_definition_display_properties.rendered,
    module.unlock_unlock_definition_hash.rendered,
    module.unlock_unlock_definition_index.rendered,
    module.unlock_unlock_definition_redacted.rendered,
    module.unlock_unlock_definition_scope.rendered,
    module.unlock_unlock_definition_unlock_type.rendered,
  ])
}

module unlock_unlock_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module unlock_unlock_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module unlock_unlock_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module unlock_unlock_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module unlock_unlock_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module unlock_unlock_definition_scope {
  source = "./attribute"

  name = "scope"
  type = "int"
}
module unlock_unlock_definition_unlock_type {
  source = "./attribute"

  name = "unlock Type"
  type = "int"
}


