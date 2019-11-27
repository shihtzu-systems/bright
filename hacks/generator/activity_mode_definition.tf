// Code generated; DANGER ZONE FOR EDITS

module activity_mode_definition {
  source = "./definition"

  target = var.target

  name = "activity Mode"
  structs_content = join("\n", [
    module.activity_mode_activity_mode_definition.rendered,
  ])

}

module activity_mode_activity_mode_definition {
  source = "./struct"

  name = "activity Mode Definition"
  attributes_content = join("\n", [
    module.activity_mode_activity_mode_definition_activity_mode_category.rendered,
    module.activity_mode_activity_mode_definition_blacklisted.rendered,
    module.activity_mode_activity_mode_definition_display.rendered,
    module.activity_mode_activity_mode_definition_display_properties.rendered,
    module.activity_mode_activity_mode_definition_friendly_name.rendered,
    module.activity_mode_activity_mode_definition_hash.rendered,
    module.activity_mode_activity_mode_definition_index.rendered,
    module.activity_mode_activity_mode_definition_is_aggregate_mode.rendered,
    module.activity_mode_activity_mode_definition_is_team_based.rendered,
    module.activity_mode_activity_mode_definition_mode_type.rendered,
    module.activity_mode_activity_mode_definition_order.rendered,
    module.activity_mode_activity_mode_definition_parent_hashes.rendered,
    module.activity_mode_activity_mode_definition_pgcr_image.rendered,
    module.activity_mode_activity_mode_definition_redacted.rendered,
    module.activity_mode_activity_mode_definition_supports_feed_filtering.rendered,
    module.activity_mode_activity_mode_definition_tier.rendered,
  ])
}

module activity_mode_activity_mode_definition_activity_mode_category {
  source = "./attribute"

  name = "activity Mode Category"
  type = "int"
}
module activity_mode_activity_mode_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module activity_mode_activity_mode_definition_display {
  source = "./attribute"

  name = "display"
  type = "bool"
}
module activity_mode_activity_mode_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module activity_mode_activity_mode_definition_friendly_name {
  source = "./attribute"

  name = "friendly Name"
  type = "string"
}
module activity_mode_activity_mode_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module activity_mode_activity_mode_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module activity_mode_activity_mode_definition_is_aggregate_mode {
  source = "./attribute"

  name = "is Aggregate Mode"
  type = "bool"
}
module activity_mode_activity_mode_definition_is_team_based {
  source = "./attribute"

  name = "is Team Based"
  type = "bool"
}
module activity_mode_activity_mode_definition_mode_type {
  source = "./attribute"

  name = "mode Type"
  type = "int"
}
module activity_mode_activity_mode_definition_order {
  source = "./attribute"

  name = "order"
  type = "int"
}
module activity_mode_activity_mode_definition_parent_hashes {
  source = "./attribute"

  name = "parent Hashes"
  type = "[]int"
}
module activity_mode_activity_mode_definition_pgcr_image {
  source = "./attribute"

  name = "pgcr Image"
  type = "string"
}
module activity_mode_activity_mode_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module activity_mode_activity_mode_definition_supports_feed_filtering {
  source = "./attribute"

  name = "supports Feed Filtering"
  type = "bool"
}
module activity_mode_activity_mode_definition_tier {
  source = "./attribute"

  name = "tier"
  type = "int"
}


