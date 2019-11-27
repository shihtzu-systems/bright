// Code generated; DANGER ZONE FOR EDITS

module character_customization_category_definition {
  source = "./definition"

  target = var.target

  name = "character Customization Category"
  structs_content = join("\n", [
    module.character_customization_category_character_customization_category_definition.rendered,
  ])

}

module character_customization_category_character_customization_category_definition {
  source = "./struct"

  name = "character Customization Category Definition"
  attributes_content = join("\n", [
    module.character_customization_category_character_customization_category_definition_blacklisted.rendered,
    module.character_customization_category_character_customization_category_definition_display_properties.rendered,
    module.character_customization_category_character_customization_category_definition_hash.rendered,
    module.character_customization_category_character_customization_category_definition_index.rendered,
    module.character_customization_category_character_customization_category_definition_redacted.rendered,
  ])
}

module character_customization_category_character_customization_category_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module character_customization_category_character_customization_category_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module character_customization_category_character_customization_category_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module character_customization_category_character_customization_category_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module character_customization_category_character_customization_category_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


