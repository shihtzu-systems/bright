// Code generated; DANGER ZONE FOR EDITS

module activity_modifier_definition {
  source = "./definition"

  target = var.target

  name = "activity Modifier"
  structs_content = join("\n", [
    module.activity_modifier_activity_modifier_definition.rendered,
  ])

}

module activity_modifier_activity_modifier_definition {
  source = "./struct"

  name = "activity Modifier Definition"
  attributes_content = join("\n", [
    module.activity_modifier_activity_modifier_definition_blacklisted.rendered,
    module.activity_modifier_activity_modifier_definition_display_properties.rendered,
    module.activity_modifier_activity_modifier_definition_hash.rendered,
    module.activity_modifier_activity_modifier_definition_index.rendered,
    module.activity_modifier_activity_modifier_definition_redacted.rendered,
  ])
}

module activity_modifier_activity_modifier_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module activity_modifier_activity_modifier_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module activity_modifier_activity_modifier_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module activity_modifier_activity_modifier_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module activity_modifier_activity_modifier_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


