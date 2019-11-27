// Code generated; DANGER ZONE FOR EDITS

module season_definition {
  source = "./definition"

  target = var.target

  name = "season"
  structs_content = join("\n", [
    module.season_season_definition.rendered,
  ])

}

module season_season_definition {
  source = "./struct"

  name = "season Definition"
  attributes_content = join("\n", [
    module.season_season_definition_blacklisted.rendered,
    module.season_season_definition_display_properties.rendered,
    module.season_season_definition_hash.rendered,
    module.season_season_definition_index.rendered,
    module.season_season_definition_redacted.rendered,
    module.season_season_definition_season_number.rendered,
    module.season_season_definition_season_pass_progression_hash.rendered,
    module.season_season_definition_season_pass_unlock_hash.rendered,
    module.season_season_definition_start_time_in_seconds.rendered,
  ])
}

module season_season_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module season_season_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module season_season_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module season_season_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module season_season_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module season_season_definition_season_number {
  source = "./attribute"

  name = "season Number"
  type = "int"
}
module season_season_definition_season_pass_progression_hash {
  source = "./attribute"

  name = "season Pass Progression Hash"
  type = "int"
}
module season_season_definition_season_pass_unlock_hash {
  source = "./attribute"

  name = "season Pass Unlock Hash"
  type = "int"
}
module season_season_definition_start_time_in_seconds {
  source = "./attribute"

  name = "start Time In Seconds"
  type = "string"
}


