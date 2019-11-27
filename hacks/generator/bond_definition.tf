// Code generated; DANGER ZONE FOR EDITS

module bond_definition {
  source = "./definition"

  target = var.target

  name = "bond"
  structs_content = join("\n", [
    module.bond_bond_definition.rendered,
  ])

}

module bond_bond_definition {
  source = "./struct"

  name = "bond Definition"
  attributes_content = join("\n", [
    module.bond_bond_definition_blacklisted.rendered,
    module.bond_bond_definition_display_properties.rendered,
    module.bond_bond_definition_hash.rendered,
    module.bond_bond_definition_index.rendered,
    module.bond_bond_definition_provided_unlock_hash.rendered,
    module.bond_bond_definition_provided_unlock_value_hash.rendered,
    module.bond_bond_definition_redacted.rendered,
  ])
}

module bond_bond_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module bond_bond_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module bond_bond_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module bond_bond_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module bond_bond_definition_provided_unlock_hash {
  source = "./attribute"

  name = "provided Unlock Hash"
  type = "int"
}
module bond_bond_definition_provided_unlock_value_hash {
  source = "./attribute"

  name = "provided Unlock Value Hash"
  type = "int"
}
module bond_bond_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


