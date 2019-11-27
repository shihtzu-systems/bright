// Code generated; DANGER ZONE FOR EDITS

module season_pass_definition {
  source = "./definition"

  target = var.target

  name = "season Pass"
  structs_content = join("\n", [
    module.season_pass_season_pass_definition.rendered,
  ])

}

module season_pass_season_pass_definition {
  source = "./struct"

  name = "season Pass Definition"
  attributes_content = join("\n", [
    module.season_pass_season_pass_definition_blacklisted.rendered,
    module.season_pass_season_pass_definition_display_properties.rendered,
    module.season_pass_season_pass_definition_hash.rendered,
    module.season_pass_season_pass_definition_index.rendered,
    module.season_pass_season_pass_definition_prestige_progression_hash.rendered,
    module.season_pass_season_pass_definition_redacted.rendered,
    module.season_pass_season_pass_definition_reward_progression_hash.rendered,
  ])
}

module season_pass_season_pass_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module season_pass_season_pass_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module season_pass_season_pass_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module season_pass_season_pass_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module season_pass_season_pass_definition_prestige_progression_hash {
  source = "./attribute"

  name = "prestige Progression Hash"
  type = "int64"
}
module season_pass_season_pass_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module season_pass_season_pass_definition_reward_progression_hash {
  source = "./attribute"

  name = "reward Progression Hash"
  type = "int"
}


