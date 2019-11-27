// Code generated; DANGER ZONE FOR EDITS

module sack_reward_item_list_definition {
  source = "./definition"

  target = var.target

  name = "sack Reward Item List"
  structs_content = join("\n", [
    module.sack_reward_item_list_sack_reward_item_list_definition.rendered,
  ])

}

module sack_reward_item_list_sack_reward_item_list_definition {
  source = "./struct"

  name = "sack Reward Item List Definition"
  attributes_content = join("\n", [
    module.sack_reward_item_list_sack_reward_item_list_definition_blacklisted.rendered,
    module.sack_reward_item_list_sack_reward_item_list_definition_hash.rendered,
    module.sack_reward_item_list_sack_reward_item_list_definition_index.rendered,
    module.sack_reward_item_list_sack_reward_item_list_definition_redacted.rendered,
  ])
}

module sack_reward_item_list_sack_reward_item_list_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module sack_reward_item_list_sack_reward_item_list_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module sack_reward_item_list_sack_reward_item_list_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module sack_reward_item_list_sack_reward_item_list_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


