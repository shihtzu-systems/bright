// Code generated; DANGER ZONE FOR EDITS

module place_definition {
  source = "./definition"

  target = var.target

  name = "place"
  structs_content = join("\n", [
    module.place_place_definition.rendered,
  ])

}

module place_place_definition {
  source = "./struct"

  name = "place Definition"
  attributes_content = join("\n", [
    module.place_place_definition_blacklisted.rendered,
    module.place_place_definition_display_properties.rendered,
    module.place_place_definition_hash.rendered,
    module.place_place_definition_index.rendered,
    module.place_place_definition_redacted.rendered,
  ])
}

module place_place_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module place_place_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module place_place_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module place_place_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module place_place_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


