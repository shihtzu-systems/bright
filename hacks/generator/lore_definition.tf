// Code generated; DANGER ZONE FOR EDITS

module lore_definition {
  source = "./definition"

  target = var.target

  name = "lore"
  structs_content = join("\n", [
    module.lore_lore_definition.rendered,
  ])

}

module lore_lore_definition {
  source = "./struct"

  name = "lore Definition"
  attributes_content = join("\n", [
    module.lore_lore_definition_blacklisted.rendered,
    module.lore_lore_definition_display_properties.rendered,
    module.lore_lore_definition_hash.rendered,
    module.lore_lore_definition_index.rendered,
    module.lore_lore_definition_redacted.rendered,
    module.lore_lore_definition_subtitle.rendered,
  ])
}

module lore_lore_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module lore_lore_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module lore_lore_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module lore_lore_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module lore_lore_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module lore_lore_definition_subtitle {
  source = "./attribute"

  name = "subtitle"
  type = "string"
}


