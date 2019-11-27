// Code generated; DANGER ZONE FOR EDITS

module enemy_race_definition {
  source = "./definition"

  target = var.target

  name = "enemy Race"
  structs_content = join("\n", [
    module.enemy_race_enemy_race_definition.rendered,
  ])

}

module enemy_race_enemy_race_definition {
  source = "./struct"

  name = "enemy Race Definition"
  attributes_content = join("\n", [
    module.enemy_race_enemy_race_definition_blacklisted.rendered,
    module.enemy_race_enemy_race_definition_display_properties.rendered,
    module.enemy_race_enemy_race_definition_hash.rendered,
    module.enemy_race_enemy_race_definition_index.rendered,
    module.enemy_race_enemy_race_definition_redacted.rendered,
  ])
}

module enemy_race_enemy_race_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module enemy_race_enemy_race_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module enemy_race_enemy_race_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module enemy_race_enemy_race_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module enemy_race_enemy_race_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


