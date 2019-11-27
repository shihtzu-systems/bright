// Code generated; DANGER ZONE FOR EDITS

module achievement_definition {
  source = "./definition"

  target = var.target

  name = "achievement"
  structs_content = join("\n", [
    module.achievement_achievement_definition.rendered,
  ])

}

module achievement_achievement_definition {
  source = "./struct"

  name = "achievement Definition"
  attributes_content = join("\n", [
    module.achievement_achievement_definition_acccumulator_threshold.rendered,
    module.achievement_achievement_definition_blacklisted.rendered,
    module.achievement_achievement_definition_display_properties.rendered,
    module.achievement_achievement_definition_hash.rendered,
    module.achievement_achievement_definition_index.rendered,
    module.achievement_achievement_definition_platform_index.rendered,
    module.achievement_achievement_definition_redacted.rendered,
  ])
}

module achievement_achievement_definition_acccumulator_threshold {
  source = "./attribute"

  name = "acccumulator Threshold"
  type = "int"
}
module achievement_achievement_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module achievement_achievement_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module achievement_achievement_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module achievement_achievement_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module achievement_achievement_definition_platform_index {
  source = "./attribute"

  name = "platform Index"
  type = "int"
}
module achievement_achievement_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


