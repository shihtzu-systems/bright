// Code generated; DANGER ZONE FOR EDITS

module progression_definition {
  source = "./definition"

  target = var.target

  name = "progression"
  structs_content = join("\n", [
    module.progression_progression_definition.rendered,
    module.progression_steps.rendered,
  ])

}

module progression_progression_definition {
  source = "./struct"

  name = "progression Definition"
  attributes_content = join("\n", [
    module.progression_progression_definition_blacklisted.rendered,
    module.progression_progression_definition_current_reset_count_unlock_value_hash.rendered,
    module.progression_progression_definition_display_properties.rendered,
    module.progression_progression_definition_hash.rendered,
    module.progression_progression_definition_index.rendered,
    module.progression_progression_definition_progress_to_next_step_scaling.rendered,
    module.progression_progression_definition_redacted.rendered,
    module.progression_progression_definition_repeat_last_step.rendered,
    module.progression_progression_definition_reward_items.rendered,
    module.progression_progression_definition_scope.rendered,
    module.progression_progression_definition_steps.rendered,
    module.progression_progression_definition_storage_mapping_index.rendered,
    module.progression_progression_definition_visible.rendered,
  ])
}

module progression_progression_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module progression_progression_definition_current_reset_count_unlock_value_hash {
  source = "./attribute"

  name = "current Reset Count Unlock Value Hash"
  type = "int"
}
module progression_progression_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module progression_progression_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module progression_progression_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module progression_progression_definition_progress_to_next_step_scaling {
  source = "./attribute"

  name = "progress To Next Step Scaling"
  type = "int"
}
module progression_progression_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module progression_progression_definition_repeat_last_step {
  source = "./attribute"

  name = "repeat Last Step"
  type = "bool"
}
module progression_progression_definition_reward_items {
  source = "./attribute"

  name = "reward Items"
  type = "[]interface{}"
}
module progression_progression_definition_scope {
  source = "./attribute"

  name = "scope"
  type = "int"
}
module progression_progression_definition_steps {
  source = "./attribute"

  name = "steps"
  type = "[]Steps"
}
module progression_progression_definition_storage_mapping_index {
  source = "./attribute"

  name = "storage Mapping Index"
  type = "int"
}
module progression_progression_definition_visible {
  source = "./attribute"

  name = "visible"
  type = "bool"
}


module progression_steps {
  source = "./struct"

  name = "steps"
  attributes_content = join("\n", [
    module.progression_steps_display_effect_type.rendered,
    module.progression_steps_progress_total.rendered,
    module.progression_steps_step_name.rendered,
  ])
}

module progression_steps_display_effect_type {
  source = "./attribute"

  name = "display Effect Type"
  type = "int"
}
module progression_steps_progress_total {
  source = "./attribute"

  name = "progress Total"
  type = "int"
}
module progression_steps_step_name {
  source = "./attribute"

  name = "step Name"
  type = "string"
}


