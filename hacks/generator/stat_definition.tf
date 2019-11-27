// Code generated; DANGER ZONE FOR EDITS

module stat_definition {
  source = "./definition"

  target = var.target

  name = "stat"
  structs_content = join("\n", [
    module.stat_stat_definition.rendered,
  ])

}

module stat_stat_definition {
  source = "./struct"

  name = "stat Definition"
  attributes_content = join("\n", [
    module.stat_stat_definition_aggregation_type.rendered,
    module.stat_stat_definition_blacklisted.rendered,
    module.stat_stat_definition_display_properties.rendered,
    module.stat_stat_definition_has_computed_block.rendered,
    module.stat_stat_definition_hash.rendered,
    module.stat_stat_definition_index.rendered,
    module.stat_stat_definition_interpolate.rendered,
    module.stat_stat_definition_redacted.rendered,
    module.stat_stat_definition_stat_category.rendered,
  ])
}

module stat_stat_definition_aggregation_type {
  source = "./attribute"

  name = "aggregation Type"
  type = "int"
}
module stat_stat_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module stat_stat_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module stat_stat_definition_has_computed_block {
  source = "./attribute"

  name = "has Computed Block"
  type = "bool"
}
module stat_stat_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module stat_stat_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module stat_stat_definition_interpolate {
  source = "./attribute"

  name = "interpolate"
  type = "bool"
}
module stat_stat_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module stat_stat_definition_stat_category {
  source = "./attribute"

  name = "stat Category"
  type = "int"
}


