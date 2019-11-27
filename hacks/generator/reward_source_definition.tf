// Code generated; DANGER ZONE FOR EDITS

module reward_source_definition {
  source = "./definition"

  target = var.target

  name = "reward Source"
  structs_content = join("\n", [
    module.reward_source_reward_source_definition.rendered,
  ])

}

module reward_source_reward_source_definition {
  source = "./struct"

  name = "reward Source Definition"
  attributes_content = join("\n", [
  ])
}



