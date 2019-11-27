// Code generated; DANGER ZONE FOR EDITS

module record_definition {
  source = "./definition"

  target = var.target

  name = "record"
  structs_content = join("\n", [
    module.record_record_definition.rendered,
    module.record_completion_info.rendered,
    module.record_expiration_info.rendered,
    module.record_interval_info.rendered,
    module.record_record_presentation_info.rendered,
    module.record_requirements.rendered,
    module.record_record_state_info.rendered,
    module.record_title_info.rendered,
  ])

}

module record_record_definition {
  source = "./struct"

  name = "record Definition"
  attributes_content = join("\n", [
    module.record_record_definition_blacklisted.rendered,
    module.record_record_definition_completion_info.rendered,
    module.record_record_definition_display_properties.rendered,
    module.record_record_definition_expiration_info.rendered,
    module.record_record_definition_hash.rendered,
    module.record_record_definition_index.rendered,
    module.record_record_definition_interval_info.rendered,
    module.record_record_definition_objective_hashes.rendered,
    module.record_record_definition_presentation_info.rendered,
    module.record_record_definition_record_value_style.rendered,
    module.record_record_definition_redacted.rendered,
    module.record_record_definition_requirements.rendered,
    module.record_record_definition_reward_items.rendered,
    module.record_record_definition_scope.rendered,
    module.record_record_definition_state_info.rendered,
    module.record_record_definition_title_info.rendered,
  ])
}

module record_record_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module record_record_definition_completion_info {
  source = "./attribute"

  name = "completion Info"
  type = "CompletionInfo"
}
module record_record_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module record_record_definition_expiration_info {
  source = "./attribute"

  name = "expiration Info"
  type = "ExpirationInfo"
}
module record_record_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module record_record_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module record_record_definition_interval_info {
  source = "./attribute"

  name = "interval Info"
  type = "IntervalInfo"
}
module record_record_definition_objective_hashes {
  source = "./attribute"

  name = "objective Hashes"
  type = "[]int64"
}
module record_record_definition_presentation_info {
  source = "./attribute"

  name = "presentation Info"
  type = "RecordPresentationInfo"
}
module record_record_definition_record_value_style {
  source = "./attribute"

  name = "record Value Style"
  type = "int"
}
module record_record_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module record_record_definition_requirements {
  source = "./attribute"

  name = "requirements"
  type = "Requirements"
}
module record_record_definition_reward_items {
  source = "./attribute"

  name = "reward Items"
  type = "[]interface{}"
}
module record_record_definition_scope {
  source = "./attribute"

  name = "scope"
  type = "int"
}
module record_record_definition_state_info {
  source = "./attribute"

  name = "state Info"
  type = "RecordStateInfo"
}
module record_record_definition_title_info {
  source = "./attribute"

  name = "title Info"
  type = "TitleInfo"
}


module record_completion_info {
  source = "./struct"

  name = "completion Info"
  attributes_content = join("\n", [
    module.record_completion_info_score_value.rendered,
    module.record_completion_info_partial_completion_objective_count_threshold.rendered,
    module.record_completion_info_should_fire_toast.rendered,
    module.record_completion_info_toast_style.rendered,
  ])
}

module record_completion_info_score_value {
  source = "./attribute"

  name = "Score Value"
  type = "int"
}
module record_completion_info_partial_completion_objective_count_threshold {
  source = "./attribute"

  name = "partial Completion Objective Count Threshold"
  type = "int"
}
module record_completion_info_should_fire_toast {
  source = "./attribute"

  name = "should Fire Toast"
  type = "bool"
}
module record_completion_info_toast_style {
  source = "./attribute"

  name = "toast Style"
  type = "int"
}


module record_expiration_info {
  source = "./struct"

  name = "expiration Info"
  attributes_content = join("\n", [
    module.record_expiration_info_description.rendered,
    module.record_expiration_info_has_expiration.rendered,
  ])
}

module record_expiration_info_description {
  source = "./attribute"

  name = "description"
  type = "string"
}
module record_expiration_info_has_expiration {
  source = "./attribute"

  name = "has Expiration"
  type = "bool"
}


module record_interval_info {
  source = "./struct"

  name = "interval Info"
  attributes_content = join("\n", [
    module.record_interval_info_interval_objectives.rendered,
    module.record_interval_info_is_interval_versioned_from_normal_record.rendered,
    module.record_interval_info_original_objective_array_insertion_index.rendered,
  ])
}

module record_interval_info_interval_objectives {
  source = "./attribute"

  name = "interval Objectives"
  type = "[]interface{}"
}
module record_interval_info_is_interval_versioned_from_normal_record {
  source = "./attribute"

  name = "is Interval Versioned From Normal Record"
  type = "bool"
}
module record_interval_info_original_objective_array_insertion_index {
  source = "./attribute"

  name = "original Objective Array Insertion Index"
  type = "int"
}


module record_record_presentation_info {
  source = "./struct"

  name = "record Presentation Info"
  attributes_content = join("\n", [
    module.record_record_presentation_info_display_style.rendered,
    module.record_record_presentation_info_parent_presentation_node_hashes.rendered,
    module.record_record_presentation_info_presentation_node_type.rendered,
  ])
}

module record_record_presentation_info_display_style {
  source = "./attribute"

  name = "display Style"
  type = "int"
}
module record_record_presentation_info_parent_presentation_node_hashes {
  source = "./attribute"

  name = "parent Presentation Node Hashes"
  type = "[]int64"
}
module record_record_presentation_info_presentation_node_type {
  source = "./attribute"

  name = "presentation Node Type"
  type = "int"
}


module record_requirements {
  source = "./struct"

  name = "requirements"
  attributes_content = join("\n", [
    module.record_requirements_entitlement_unavailable_message.rendered,
  ])
}

module record_requirements_entitlement_unavailable_message {
  source = "./attribute"

  name = "entitlement Unavailable Message"
  type = "string"
}


module record_record_state_info {
  source = "./struct"

  name = "record State Info"
  attributes_content = join("\n", [
    module.record_record_state_info_claimed_unlock_hash.rendered,
    module.record_record_state_info_complete_unlock_hash.rendered,
    module.record_record_state_info_featured_priority.rendered,
  ])
}

module record_record_state_info_claimed_unlock_hash {
  source = "./attribute"

  name = "claimed Unlock Hash"
  type = "int"
}
module record_record_state_info_complete_unlock_hash {
  source = "./attribute"

  name = "complete Unlock Hash"
  type = "int"
}
module record_record_state_info_featured_priority {
  source = "./attribute"

  name = "featured Priority"
  type = "int64"
}


module record_title_info {
  source = "./struct"

  name = "title Info"
  attributes_content = join("\n", [
    module.record_title_info_has_title.rendered,
  ])
}

module record_title_info_has_title {
  source = "./attribute"

  name = "has Title"
  type = "bool"
}


