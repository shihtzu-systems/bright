// Code generated; DANGER ZONE FOR EDITS

module sandbox_perk_definition {
  source = "./definition"

  target = var.target

  name = "sandbox Perk"
  structs_content = join("\n", [
    module.sandbox_perk_sandbox_perk_definition.rendered,
  ])

}

module sandbox_perk_sandbox_perk_definition {
  source = "./struct"

  name = "sandbox Perk Definition"
  attributes_content = join("\n", [
    module.sandbox_perk_sandbox_perk_definition_blacklisted.rendered,
    module.sandbox_perk_sandbox_perk_definition_damage_type.rendered,
    module.sandbox_perk_sandbox_perk_definition_damage_type_hash.rendered,
    module.sandbox_perk_sandbox_perk_definition_display_properties.rendered,
    module.sandbox_perk_sandbox_perk_definition_hash.rendered,
    module.sandbox_perk_sandbox_perk_definition_index.rendered,
    module.sandbox_perk_sandbox_perk_definition_is_displayable.rendered,
    module.sandbox_perk_sandbox_perk_definition_redacted.rendered,
  ])
}

module sandbox_perk_sandbox_perk_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module sandbox_perk_sandbox_perk_definition_damage_type {
  source = "./attribute"

  name = "damage Type"
  type = "int"
}
module sandbox_perk_sandbox_perk_definition_damage_type_hash {
  source = "./attribute"

  name = "damage Type Hash"
  type = "int64"
}
module sandbox_perk_sandbox_perk_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module sandbox_perk_sandbox_perk_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module sandbox_perk_sandbox_perk_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module sandbox_perk_sandbox_perk_definition_is_displayable {
  source = "./attribute"

  name = "is Displayable"
  type = "bool"
}
module sandbox_perk_sandbox_perk_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


