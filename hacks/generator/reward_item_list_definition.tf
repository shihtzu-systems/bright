// Code generated; DANGER ZONE FOR EDITS

module reward_item_list_definition {
  source = "./definition"

  target = var.target

  name = "reward Item List"
  structs_content = join("\n", [
    module.reward_item_list_reward_item_list_definition.rendered,
  ])

}

module reward_item_list_reward_item_list_definition {
  source = "./struct"

  name = "reward Item List Definition"
  attributes_content = join("\n", [
    module.reward_item_list_reward_item_list_definition_blacklisted.rendered,
    module.reward_item_list_reward_item_list_definition_hash.rendered,
    module.reward_item_list_reward_item_list_definition_index.rendered,
    module.reward_item_list_reward_item_list_definition_redacted.rendered,
  ])
}

module reward_item_list_reward_item_list_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module reward_item_list_reward_item_list_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module reward_item_list_reward_item_list_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module reward_item_list_reward_item_list_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


