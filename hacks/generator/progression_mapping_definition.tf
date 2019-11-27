// Code generated; DANGER ZONE FOR EDITS

module progression_mapping_definition {
  source = "./definition"

  target = var.target

  name = "progression Mapping"
  structs_content = join("\n", [
    module.progression_mapping_progression_mapping_definition.rendered,
  ])

}

module progression_mapping_progression_mapping_definition {
  source = "./struct"

  name = "progression Mapping Definition"
  attributes_content = join("\n", [
    module.progression_mapping_progression_mapping_definition_blacklisted.rendered,
    module.progression_mapping_progression_mapping_definition_display_properties.rendered,
    module.progression_mapping_progression_mapping_definition_hash.rendered,
    module.progression_mapping_progression_mapping_definition_index.rendered,
    module.progression_mapping_progression_mapping_definition_redacted.rendered,
  ])
}

module progression_mapping_progression_mapping_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module progression_mapping_progression_mapping_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module progression_mapping_progression_mapping_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module progression_mapping_progression_mapping_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module progression_mapping_progression_mapping_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


