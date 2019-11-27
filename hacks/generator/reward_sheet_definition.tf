// Code generated; DANGER ZONE FOR EDITS

module reward_sheet_definition {
  source = "./definition"

  target = var.target

  name = "reward Sheet"
  structs_content = join("\n", [
    module.reward_sheet_reward_sheet_definition.rendered,
  ])

}

module reward_sheet_reward_sheet_definition {
  source = "./struct"

  name = "reward Sheet Definition"
  attributes_content = join("\n", [
    module.reward_sheet_reward_sheet_definition_blacklisted.rendered,
    module.reward_sheet_reward_sheet_definition_hash.rendered,
    module.reward_sheet_reward_sheet_definition_index.rendered,
    module.reward_sheet_reward_sheet_definition_redacted.rendered,
    module.reward_sheet_reward_sheet_definition_sheet_hash.rendered,
    module.reward_sheet_reward_sheet_definition_sheet_index.rendered,
  ])
}

module reward_sheet_reward_sheet_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module reward_sheet_reward_sheet_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module reward_sheet_reward_sheet_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module reward_sheet_reward_sheet_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module reward_sheet_reward_sheet_definition_sheet_hash {
  source = "./attribute"

  name = "sheet Hash"
  type = "int"
}
module reward_sheet_reward_sheet_definition_sheet_index {
  source = "./attribute"

  name = "sheet Index"
  type = "int"
}


