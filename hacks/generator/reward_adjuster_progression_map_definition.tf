// Code generated; DANGER ZONE FOR EDITS

module reward_adjuster_progression_map_definition {
  source = "./definition"

  target = var.target

  name = "reward Adjuster Progression Map"
  structs_content = join("\n", [
    module.reward_adjuster_progression_map_reward_adjuster_progression_map_definition.rendered,
  ])

}

module reward_adjuster_progression_map_reward_adjuster_progression_map_definition {
  source = "./struct"

  name = "reward Adjuster Progression Map Definition"
  attributes_content = join("\n", [
    module.reward_adjuster_progression_map_reward_adjuster_progression_map_definition_blacklisted.rendered,
    module.reward_adjuster_progression_map_reward_adjuster_progression_map_definition_hash.rendered,
    module.reward_adjuster_progression_map_reward_adjuster_progression_map_definition_index.rendered,
    module.reward_adjuster_progression_map_reward_adjuster_progression_map_definition_is_additive.rendered,
    module.reward_adjuster_progression_map_reward_adjuster_progression_map_definition_redacted.rendered,
  ])
}

module reward_adjuster_progression_map_reward_adjuster_progression_map_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module reward_adjuster_progression_map_reward_adjuster_progression_map_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module reward_adjuster_progression_map_reward_adjuster_progression_map_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module reward_adjuster_progression_map_reward_adjuster_progression_map_definition_is_additive {
  source = "./attribute"

  name = "is Additive"
  type = "bool"
}
module reward_adjuster_progression_map_reward_adjuster_progression_map_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


