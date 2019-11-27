// Code generated; DANGER ZONE FOR EDITS

module artifact_definition {
  source = "./definition"

  target = var.target

  name = "artifact"
  structs_content = join("\n", [
    module.artifact_artifact_definition.rendered,
    module.artifact_artifact_items.rendered,
    module.artifact_tiers.rendered,
  ])

}

module artifact_artifact_definition {
  source = "./struct"

  name = "artifact Definition"
  attributes_content = join("\n", [
    module.artifact_artifact_definition_blacklisted.rendered,
    module.artifact_artifact_definition_display_properties.rendered,
    module.artifact_artifact_definition_hash.rendered,
    module.artifact_artifact_definition_index.rendered,
    module.artifact_artifact_definition_redacted.rendered,
    module.artifact_artifact_definition_tiers.rendered,
    module.artifact_artifact_definition_translation_block.rendered,
  ])
}

module artifact_artifact_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module artifact_artifact_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module artifact_artifact_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module artifact_artifact_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module artifact_artifact_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module artifact_artifact_definition_tiers {
  source = "./attribute"

  name = "tiers"
  type = "[]Tiers"
}
module artifact_artifact_definition_translation_block {
  source = "./attribute"

  name = "translation Block"
  type = "TranslationBlock"
}


module artifact_artifact_items {
  source = "./struct"

  name = "artifact Items"
  attributes_content = join("\n", [
    module.artifact_artifact_items_active_unlock_hash.rendered,
    module.artifact_artifact_items_item_hash.rendered,
  ])
}

module artifact_artifact_items_active_unlock_hash {
  source = "./attribute"

  name = "active Unlock Hash"
  type = "int"
}
module artifact_artifact_items_item_hash {
  source = "./attribute"

  name = "item Hash"
  type = "int64"
}


module artifact_tiers {
  source = "./struct"

  name = "tiers"
  attributes_content = join("\n", [
    module.artifact_tiers_display_title.rendered,
    module.artifact_tiers_items.rendered,
    module.artifact_tiers_minimum_unlock_points_used_requirement.rendered,
    module.artifact_tiers_progress_requirement_message.rendered,
    module.artifact_tiers_tier_hash.rendered,
  ])
}

module artifact_tiers_display_title {
  source = "./attribute"

  name = "display Title"
  type = "string"
}
module artifact_tiers_items {
  source = "./attribute"

  name = "items"
  type = "[]ArtifactItems"
}
module artifact_tiers_minimum_unlock_points_used_requirement {
  source = "./attribute"

  name = "minimum Unlock Points Used Requirement"
  type = "int"
}
module artifact_tiers_progress_requirement_message {
  source = "./attribute"

  name = "progress Requirement Message"
  type = "string"
}
module artifact_tiers_tier_hash {
  source = "./attribute"

  name = "tier Hash"
  type = "int64"
}


