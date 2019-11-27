// Code generated; DANGER ZONE FOR EDITS

module progression_level_requirement_definition {
  source = "./definition"

  target = var.target

  name = "progression Level Requirement"
  structs_content = join("\n", [
    module.progression_level_requirement_progression_level_requirement_definition.rendered,
    module.progression_level_requirement_requirement_curve.rendered,
  ])

}

module progression_level_requirement_progression_level_requirement_definition {
  source = "./struct"

  name = "progression Level Requirement Definition"
  attributes_content = join("\n", [
    module.progression_level_requirement_progression_level_requirement_definition_blacklisted.rendered,
    module.progression_level_requirement_progression_level_requirement_definition_hash.rendered,
    module.progression_level_requirement_progression_level_requirement_definition_index.rendered,
    module.progression_level_requirement_progression_level_requirement_definition_progression_hash.rendered,
    module.progression_level_requirement_progression_level_requirement_definition_redacted.rendered,
    module.progression_level_requirement_progression_level_requirement_definition_requirement_curve.rendered,
  ])
}

module progression_level_requirement_progression_level_requirement_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module progression_level_requirement_progression_level_requirement_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module progression_level_requirement_progression_level_requirement_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module progression_level_requirement_progression_level_requirement_definition_progression_hash {
  source = "./attribute"

  name = "progression Hash"
  type = "int"
}
module progression_level_requirement_progression_level_requirement_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module progression_level_requirement_progression_level_requirement_definition_requirement_curve {
  source = "./attribute"

  name = "requirement Curve"
  type = "[]RequirementCurve"
}


module progression_level_requirement_requirement_curve {
  source = "./struct"

  name = "requirement Curve"
  attributes_content = join("\n", [
    module.progression_level_requirement_requirement_curve_value.rendered,
    module.progression_level_requirement_requirement_curve_weight.rendered,
  ])
}

module progression_level_requirement_requirement_curve_value {
  source = "./attribute"

  name = "value"
  type = "float64"
}
module progression_level_requirement_requirement_curve_weight {
  source = "./attribute"

  name = "weight"
  type = "float64"
}


