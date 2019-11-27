// Code generated; DANGER ZONE FOR EDITS

module stat_group_definition {
  source = "./definition"

  target = var.target

  name = "stat Group"
  structs_content = join("\n", [
    module.stat_group_stat_group_definition.rendered,
    module.stat_group_display_interpolation.rendered,
    module.stat_group_scaled_stats.rendered,
  ])

}

module stat_group_stat_group_definition {
  source = "./struct"

  name = "stat Group Definition"
  attributes_content = join("\n", [
    module.stat_group_stat_group_definition_blacklisted.rendered,
    module.stat_group_stat_group_definition_hash.rendered,
    module.stat_group_stat_group_definition_index.rendered,
    module.stat_group_stat_group_definition_maximum_value.rendered,
    module.stat_group_stat_group_definition_overrides.rendered,
    module.stat_group_stat_group_definition_redacted.rendered,
    module.stat_group_stat_group_definition_scaled_stats.rendered,
    module.stat_group_stat_group_definition_ui_position.rendered,
  ])
}

module stat_group_stat_group_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module stat_group_stat_group_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module stat_group_stat_group_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module stat_group_stat_group_definition_maximum_value {
  source = "./attribute"

  name = "maximum Value"
  type = "int"
}
module stat_group_stat_group_definition_overrides {
  source = "./attribute"

  name = "overrides"
  type = "interface{}"
}
module stat_group_stat_group_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module stat_group_stat_group_definition_scaled_stats {
  source = "./attribute"

  name = "scaled Stats"
  type = "[]ScaledStats"
}
module stat_group_stat_group_definition_ui_position {
  source = "./attribute"

  name = "ui Position"
  type = "int"
}


module stat_group_display_interpolation {
  source = "./struct"

  name = "display Interpolation"
  attributes_content = join("\n", [
    module.stat_group_display_interpolation_value.rendered,
    module.stat_group_display_interpolation_weight.rendered,
  ])
}

module stat_group_display_interpolation_value {
  source = "./attribute"

  name = "value"
  type = "int"
}
module stat_group_display_interpolation_weight {
  source = "./attribute"

  name = "weight"
  type = "int"
}


module stat_group_scaled_stats {
  source = "./struct"

  name = "scaled Stats"
  attributes_content = join("\n", [
    module.stat_group_scaled_stats_display_as_numeric.rendered,
    module.stat_group_scaled_stats_display_interpolation.rendered,
    module.stat_group_scaled_stats_maximum_value.rendered,
    module.stat_group_scaled_stats_stat_hash.rendered,
  ])
}

module stat_group_scaled_stats_display_as_numeric {
  source = "./attribute"

  name = "display As Numeric"
  type = "bool"
}
module stat_group_scaled_stats_display_interpolation {
  source = "./attribute"

  name = "display Interpolation"
  type = "[]DisplayInterpolation"
}
module stat_group_scaled_stats_maximum_value {
  source = "./attribute"

  name = "maximum Value"
  type = "int"
}
module stat_group_scaled_stats_stat_hash {
  source = "./attribute"

  name = "stat Hash"
  type = "int64"
}


