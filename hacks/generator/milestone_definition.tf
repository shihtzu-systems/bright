// Code generated; DANGER ZONE FOR EDITS

module milestone_definition {
  source = "./definition"

  target = var.target

  name = "milestone"
  structs_content = join("\n", [
    module.milestone_milestone_definition.rendered,
    module.milestone_milestone_items.rendered,
    module.milestone_quest_rewards.rendered,
    module.milestone_quest.rendered,
  ])

}

module milestone_milestone_definition {
  source = "./struct"

  name = "milestone Definition"
  attributes_content = join("\n", [
    module.milestone_milestone_definition_blacklisted.rendered,
    module.milestone_milestone_definition_default_order.rendered,
    module.milestone_milestone_definition_display_properties.rendered,
    module.milestone_milestone_definition_explore_prioritizes_activity_image.rendered,
    module.milestone_milestone_definition_has_predictable_dates.rendered,
    module.milestone_milestone_definition_hash.rendered,
    module.milestone_milestone_definition_index.rendered,
    module.milestone_milestone_definition_is_in_game_milestone.rendered,
    module.milestone_milestone_definition_milestone_type.rendered,
    module.milestone_milestone_definition_quests.rendered,
    module.milestone_milestone_definition_recruitable.rendered,
    module.milestone_milestone_definition_redacted.rendered,
    module.milestone_milestone_definition_show_in_explorer.rendered,
    module.milestone_milestone_definition_show_in_milestones.rendered,
  ])
}

module milestone_milestone_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module milestone_milestone_definition_default_order {
  source = "./attribute"

  name = "default Order"
  type = "int"
}
module milestone_milestone_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module milestone_milestone_definition_explore_prioritizes_activity_image {
  source = "./attribute"

  name = "explore Prioritizes Activity Image"
  type = "bool"
}
module milestone_milestone_definition_has_predictable_dates {
  source = "./attribute"

  name = "has Predictable Dates"
  type = "bool"
}
module milestone_milestone_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module milestone_milestone_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module milestone_milestone_definition_is_in_game_milestone {
  source = "./attribute"

  name = "is In Game Milestone"
  type = "bool"
}
module milestone_milestone_definition_milestone_type {
  source = "./attribute"

  name = "milestone Type"
  type = "int"
}
module milestone_milestone_definition_quests {
  source = "./attribute"

  name = "quests"
  type = "map[string]Quest"
}
module milestone_milestone_definition_recruitable {
  source = "./attribute"

  name = "recruitable"
  type = "bool"
}
module milestone_milestone_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module milestone_milestone_definition_show_in_explorer {
  source = "./attribute"

  name = "show In Explorer"
  type = "bool"
}
module milestone_milestone_definition_show_in_milestones {
  source = "./attribute"

  name = "show In Milestones"
  type = "bool"
}


module milestone_milestone_items {
  source = "./struct"

  name = "milestone Items"
  attributes_content = join("\n", [
    module.milestone_milestone_items_item_hash.rendered,
    module.milestone_milestone_items_quantity.rendered,
  ])
}

module milestone_milestone_items_item_hash {
  source = "./attribute"

  name = "item Hash"
  type = "int"
}
module milestone_milestone_items_quantity {
  source = "./attribute"

  name = "quantity"
  type = "int"
}


module milestone_quest_rewards {
  source = "./struct"

  name = "quest Rewards"
  attributes_content = join("\n", [
    module.milestone_quest_rewards_items.rendered,
  ])
}

module milestone_quest_rewards_items {
  source = "./attribute"

  name = "items"
  type = "[]MilestoneItems"
}


module milestone_quest {
  source = "./struct"

  name = "quest"
  attributes_content = join("\n", [
    module.milestone_quest_display_properties.rendered,
    module.milestone_quest_quest_item_hash.rendered,
    module.milestone_quest_quest_rewards.rendered,
    module.milestone_quest_tracking_unlock_value_hash.rendered,
  ])
}

module milestone_quest_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module milestone_quest_quest_item_hash {
  source = "./attribute"

  name = "quest Item Hash"
  type = "int"
}
module milestone_quest_quest_rewards {
  source = "./attribute"

  name = "quest Rewards"
  type = "QuestRewards"
}
module milestone_quest_tracking_unlock_value_hash {
  source = "./attribute"

  name = "tracking Unlock Value Hash"
  type = "int"
}


