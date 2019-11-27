// Code generated; DANGER ZONE FOR EDITS

module damage_type_definition {
  source = "./definition"

  target = var.target

  name = "damage Type"
  structs_content = join("\n", [
    module.damage_type_damage_type_definition.rendered,
  ])

}

module damage_type_damage_type_definition {
  source = "./struct"

  name = "damage Type Definition"
  attributes_content = join("\n", [
    module.damage_type_damage_type_definition_blacklisted.rendered,
    module.damage_type_damage_type_definition_display_properties.rendered,
    module.damage_type_damage_type_definition_enum_value.rendered,
    module.damage_type_damage_type_definition_hash.rendered,
    module.damage_type_damage_type_definition_index.rendered,
    module.damage_type_damage_type_definition_redacted.rendered,
    module.damage_type_damage_type_definition_show_icon.rendered,
    module.damage_type_damage_type_definition_transparent_icon_path.rendered,
  ])
}

module damage_type_damage_type_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module damage_type_damage_type_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module damage_type_damage_type_definition_enum_value {
  source = "./attribute"

  name = "enum Value"
  type = "int"
}
module damage_type_damage_type_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module damage_type_damage_type_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module damage_type_damage_type_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module damage_type_damage_type_definition_show_icon {
  source = "./attribute"

  name = "show Icon"
  type = "bool"
}
module damage_type_damage_type_definition_transparent_icon_path {
  source = "./attribute"

  name = "transparent Icon Path"
  type = "string"
}


