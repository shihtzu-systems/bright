// Code generated; DANGER ZONE FOR EDITS

module item_tier_type_definition {
  source = "./definition"

  target = var.target

  name = "item Tier Type"
  structs_content = join("\n", [
    module.item_tier_type_item_tier_type_definition.rendered,
    module.item_tier_type_infusion_process.rendered,
  ])

}

module item_tier_type_item_tier_type_definition {
  source = "./struct"

  name = "item Tier Type Definition"
  attributes_content = join("\n", [
    module.item_tier_type_item_tier_type_definition_blacklisted.rendered,
    module.item_tier_type_item_tier_type_definition_display_properties.rendered,
    module.item_tier_type_item_tier_type_definition_hash.rendered,
    module.item_tier_type_item_tier_type_definition_index.rendered,
    module.item_tier_type_item_tier_type_definition_infusionprocess.rendered,
    module.item_tier_type_item_tier_type_definition_redacted.rendered,
  ])
}

module item_tier_type_item_tier_type_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module item_tier_type_item_tier_type_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module item_tier_type_item_tier_type_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module item_tier_type_item_tier_type_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module item_tier_type_item_tier_type_definition_infusionprocess {
  source = "./attribute"

  name = "infusionProcess"
  type = "InfusionProcess"
}
module item_tier_type_item_tier_type_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


module item_tier_type_infusion_process {
  source = "./struct"

  name = "infusion Process"
  attributes_content = join("\n", [
    module.item_tier_type_infusion_process_base_quality_transfer_ratio.rendered,
    module.item_tier_type_infusion_process_minimum_quality_increment.rendered,
  ])
}

module item_tier_type_infusion_process_base_quality_transfer_ratio {
  source = "./attribute"

  name = "base Quality Transfer Ratio"
  type = "float64"
}
module item_tier_type_infusion_process_minimum_quality_increment {
  source = "./attribute"

  name = "minimum Quality Increment"
  type = "int"
}


