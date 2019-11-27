// Code generated; DANGER ZONE FOR EDITS

module energy_type_definition {
  source = "./definition"

  target = var.target

  name = "energy Type"
  structs_content = join("\n", [
    module.energy_type_energy_type_definition.rendered,
  ])

}

module energy_type_energy_type_definition {
  source = "./struct"

  name = "energy Type Definition"
  attributes_content = join("\n", [
    module.energy_type_energy_type_definition_blacklisted.rendered,
    module.energy_type_energy_type_definition_capacity_stat_hash.rendered,
    module.energy_type_energy_type_definition_cost_stat_hash.rendered,
    module.energy_type_energy_type_definition_display_properties.rendered,
    module.energy_type_energy_type_definition_enum_value.rendered,
    module.energy_type_energy_type_definition_hash.rendered,
    module.energy_type_energy_type_definition_index.rendered,
    module.energy_type_energy_type_definition_redacted.rendered,
    module.energy_type_energy_type_definition_show_icon.rendered,
    module.energy_type_energy_type_definition_transparent_icon_path.rendered,
  ])
}

module energy_type_energy_type_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module energy_type_energy_type_definition_capacity_stat_hash {
  source = "./attribute"

  name = "capacity Stat Hash"
  type = "int"
}
module energy_type_energy_type_definition_cost_stat_hash {
  source = "./attribute"

  name = "cost Stat Hash"
  type = "int64"
}
module energy_type_energy_type_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module energy_type_energy_type_definition_enum_value {
  source = "./attribute"

  name = "enum Value"
  type = "int"
}
module energy_type_energy_type_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module energy_type_energy_type_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module energy_type_energy_type_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module energy_type_energy_type_definition_show_icon {
  source = "./attribute"

  name = "show Icon"
  type = "bool"
}
module energy_type_energy_type_definition_transparent_icon_path {
  source = "./attribute"

  name = "transparent Icon Path"
  type = "string"
}


