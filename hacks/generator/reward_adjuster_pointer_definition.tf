// Code generated; DANGER ZONE FOR EDITS

module reward_adjuster_pointer_definition {
  source = "./definition"

  target = var.target

  name = "reward Adjuster Pointer"
  structs_content = join("\n", [
    module.reward_adjuster_pointer_reward_adjuster_pointer_definition.rendered,
  ])

}

module reward_adjuster_pointer_reward_adjuster_pointer_definition {
  source = "./struct"

  name = "reward Adjuster Pointer Definition"
  attributes_content = join("\n", [
    module.reward_adjuster_pointer_reward_adjuster_pointer_definition_adjuster_type.rendered,
    module.reward_adjuster_pointer_reward_adjuster_pointer_definition_blacklisted.rendered,
    module.reward_adjuster_pointer_reward_adjuster_pointer_definition_hash.rendered,
    module.reward_adjuster_pointer_reward_adjuster_pointer_definition_index.rendered,
    module.reward_adjuster_pointer_reward_adjuster_pointer_definition_redacted.rendered,
  ])
}

module reward_adjuster_pointer_reward_adjuster_pointer_definition_adjuster_type {
  source = "./attribute"

  name = "adjuster Type"
  type = "int"
}
module reward_adjuster_pointer_reward_adjuster_pointer_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module reward_adjuster_pointer_reward_adjuster_pointer_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module reward_adjuster_pointer_reward_adjuster_pointer_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module reward_adjuster_pointer_reward_adjuster_pointer_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


