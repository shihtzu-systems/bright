// Code generated; DANGER ZONE FOR EDITS

module objective_definition {
  source = "./definition"

  target = var.target

  name = "objective"
  structs_content = join("\n", [
    module.objective_objective_definition.rendered,
  ])

}

module objective_objective_definition {
  source = "./struct"

  name = "objective Definition"
  attributes_content = join("\n", [
    module.objective_objective_definition_allow_negative_value.rendered,
    module.objective_objective_definition_allow_overcompletion.rendered,
    module.objective_objective_definition_allow_value_change_when_completed.rendered,
    module.objective_objective_definition_blacklisted.rendered,
    module.objective_objective_definition_completed_value_style.rendered,
    module.objective_objective_definition_completion_value.rendered,
    module.objective_objective_definition_display_properties.rendered,
    module.objective_objective_definition_hash.rendered,
    module.objective_objective_definition_in_progress_value_style.rendered,
    module.objective_objective_definition_index.rendered,
    module.objective_objective_definition_is_counting_downward.rendered,
    module.objective_objective_definition_is_display_only_objective.rendered,
    module.objective_objective_definition_location_hash.rendered,
    module.objective_objective_definition_minimum_visibility_threshold.rendered,
    module.objective_objective_definition_progress_description.rendered,
    module.objective_objective_definition_redacted.rendered,
    module.objective_objective_definition_scope.rendered,
    module.objective_objective_definition_show_value_on_complete.rendered,
    module.objective_objective_definition_unlock_value_hash.rendered,
    module.objective_objective_definition_value_style.rendered,
  ])
}

module objective_objective_definition_allow_negative_value {
  source = "./attribute"

  name = "allow Negative Value"
  type = "bool"
}
module objective_objective_definition_allow_overcompletion {
  source = "./attribute"

  name = "allow Overcompletion"
  type = "bool"
}
module objective_objective_definition_allow_value_change_when_completed {
  source = "./attribute"

  name = "allow Value Change When Completed"
  type = "bool"
}
module objective_objective_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module objective_objective_definition_completed_value_style {
  source = "./attribute"

  name = "completed Value Style"
  type = "int"
}
module objective_objective_definition_completion_value {
  source = "./attribute"

  name = "completion Value"
  type = "int"
}
module objective_objective_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module objective_objective_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module objective_objective_definition_in_progress_value_style {
  source = "./attribute"

  name = "in Progress Value Style"
  type = "int"
}
module objective_objective_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module objective_objective_definition_is_counting_downward {
  source = "./attribute"

  name = "is Counting Downward"
  type = "bool"
}
module objective_objective_definition_is_display_only_objective {
  source = "./attribute"

  name = "is Display Only Objective"
  type = "bool"
}
module objective_objective_definition_location_hash {
  source = "./attribute"

  name = "location Hash"
  type = "int"
}
module objective_objective_definition_minimum_visibility_threshold {
  source = "./attribute"

  name = "minimum Visibility Threshold"
  type = "int"
}
module objective_objective_definition_progress_description {
  source = "./attribute"

  name = "progress Description"
  type = "string"
}
module objective_objective_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module objective_objective_definition_scope {
  source = "./attribute"

  name = "scope"
  type = "int"
}
module objective_objective_definition_show_value_on_complete {
  source = "./attribute"

  name = "show Value On Complete"
  type = "bool"
}
module objective_objective_definition_unlock_value_hash {
  source = "./attribute"

  name = "unlock Value Hash"
  type = "int"
}
module objective_objective_definition_value_style {
  source = "./attribute"

  name = "value Style"
  type = "int"
}


