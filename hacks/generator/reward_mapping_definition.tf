// Code generated; DANGER ZONE FOR EDITS

module reward_mapping_definition {
  source = "./definition"

  target = var.target

  name = "reward Mapping"
  structs_content = join("\n", [
    module.reward_mapping_reward_mapping_definition.rendered,
  ])

}

module reward_mapping_reward_mapping_definition {
  source = "./struct"

  name = "reward Mapping Definition"
  attributes_content = join("\n", [
    module.reward_mapping_reward_mapping_definition_blacklisted.rendered,
    module.reward_mapping_reward_mapping_definition_hash.rendered,
    module.reward_mapping_reward_mapping_definition_index.rendered,
    module.reward_mapping_reward_mapping_definition_is_global.rendered,
    module.reward_mapping_reward_mapping_definition_mapping_hash.rendered,
    module.reward_mapping_reward_mapping_definition_mapping_index.rendered,
    module.reward_mapping_reward_mapping_definition_redacted.rendered,
  ])
}

module reward_mapping_reward_mapping_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module reward_mapping_reward_mapping_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module reward_mapping_reward_mapping_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module reward_mapping_reward_mapping_definition_is_global {
  source = "./attribute"

  name = "is Global"
  type = "bool"
}
module reward_mapping_reward_mapping_definition_mapping_hash {
  source = "./attribute"

  name = "mapping Hash"
  type = "int"
}
module reward_mapping_reward_mapping_definition_mapping_index {
  source = "./attribute"

  name = "mapping Index"
  type = "int"
}
module reward_mapping_reward_mapping_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


