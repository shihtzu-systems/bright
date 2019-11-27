// Code generated; DANGER ZONE FOR EDITS

module collectible_definition {
  source = "./definition"

  target = var.target

  name = "collectible"
  structs_content = join("\n", [
    module.collectible_collectible_definition.rendered,
    module.collectible_acquisition_info.rendered,
    module.collectible_collectible_presentation_info.rendered,
    module.collectible_collectible_requirements.rendered,
    module.collectible_collectible_state_info.rendered,
  ])

}

module collectible_collectible_definition {
  source = "./struct"

  name = "collectible Definition"
  attributes_content = join("\n", [
    module.collectible_collectible_definition_acquisition_info.rendered,
    module.collectible_collectible_definition_blacklisted.rendered,
    module.collectible_collectible_definition_display_properties.rendered,
    module.collectible_collectible_definition_hash.rendered,
    module.collectible_collectible_definition_index.rendered,
    module.collectible_collectible_definition_item_hash.rendered,
    module.collectible_collectible_definition_presentation_info.rendered,
    module.collectible_collectible_definition_redacted.rendered,
    module.collectible_collectible_definition_scope.rendered,
    module.collectible_collectible_definition_source_hash.rendered,
    module.collectible_collectible_definition_source_string.rendered,
    module.collectible_collectible_definition_state_info.rendered,
  ])
}

module collectible_collectible_definition_acquisition_info {
  source = "./attribute"

  name = "acquisition Info"
  type = "AcquisitionInfo"
}
module collectible_collectible_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module collectible_collectible_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module collectible_collectible_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module collectible_collectible_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module collectible_collectible_definition_item_hash {
  source = "./attribute"

  name = "item Hash"
  type = "int"
}
module collectible_collectible_definition_presentation_info {
  source = "./attribute"

  name = "presentation Info"
  type = "CollectiblePresentationInfo"
}
module collectible_collectible_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module collectible_collectible_definition_scope {
  source = "./attribute"

  name = "scope"
  type = "int"
}
module collectible_collectible_definition_source_hash {
  source = "./attribute"

  name = "source Hash"
  type = "int"
}
module collectible_collectible_definition_source_string {
  source = "./attribute"

  name = "source String"
  type = "string"
}
module collectible_collectible_definition_state_info {
  source = "./attribute"

  name = "state Info"
  type = "CollectibleStateInfo"
}


module collectible_acquisition_info {
  source = "./struct"

  name = "acquisition Info"
  attributes_content = join("\n", [
    module.collectible_acquisition_info_acquire_material_requirement_hash.rendered,
    module.collectible_acquisition_info_run_only_acquisition_reward_site.rendered,
  ])
}

module collectible_acquisition_info_acquire_material_requirement_hash {
  source = "./attribute"

  name = "acquire Material Requirement Hash"
  type = "int"
}
module collectible_acquisition_info_run_only_acquisition_reward_site {
  source = "./attribute"

  name = "run Only Acquisition Reward Site"
  type = "bool"
}


module collectible_collectible_presentation_info {
  source = "./struct"

  name = "collectible Presentation Info"
  attributes_content = join("\n", [
    module.collectible_collectible_presentation_info_display_style.rendered,
    module.collectible_collectible_presentation_info_parent_presentation_node_hashes.rendered,
    module.collectible_collectible_presentation_info_presentation_node_type.rendered,
  ])
}

module collectible_collectible_presentation_info_display_style {
  source = "./attribute"

  name = "display Style"
  type = "int"
}
module collectible_collectible_presentation_info_parent_presentation_node_hashes {
  source = "./attribute"

  name = "parent Presentation Node Hashes"
  type = "[]int"
}
module collectible_collectible_presentation_info_presentation_node_type {
  source = "./attribute"

  name = "presentation Node Type"
  type = "int"
}


module collectible_collectible_requirements {
  source = "./struct"

  name = "collectible Requirements"
  attributes_content = join("\n", [
    module.collectible_collectible_requirements_entitlement_unavailable_message.rendered,
  ])
}

module collectible_collectible_requirements_entitlement_unavailable_message {
  source = "./attribute"

  name = "entitlement Unavailable Message"
  type = "string"
}


module collectible_collectible_state_info {
  source = "./struct"

  name = "collectible State Info"
  attributes_content = join("\n", [
    module.collectible_collectible_state_info_requirements.rendered,
  ])
}

module collectible_collectible_state_info_requirements {
  source = "./attribute"

  name = "requirements"
  type = "CollectibleRequirements"
}


