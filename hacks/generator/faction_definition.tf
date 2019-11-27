// Code generated; DANGER ZONE FOR EDITS

module faction_definition {
  source = "./definition"

  target = var.target

  name = "faction"
  structs_content = join("\n", [
    module.faction_faction_definition.rendered,
    module.faction_vendors.rendered,
  ])

}

module faction_faction_definition {
  source = "./struct"

  name = "faction Definition"
  attributes_content = join("\n", [
    module.faction_faction_definition_blacklisted.rendered,
    module.faction_faction_definition_display_properties.rendered,
    module.faction_faction_definition_hash.rendered,
    module.faction_faction_definition_index.rendered,
    module.faction_faction_definition_progression_hash.rendered,
    module.faction_faction_definition_redacted.rendered,
    module.faction_faction_definition_reward_item_hash.rendered,
    module.faction_faction_definition_reward_vendor_hash.rendered,
    module.faction_faction_definition_token_values.rendered,
    module.faction_faction_definition_vendors.rendered,
  ])
}

module faction_faction_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module faction_faction_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module faction_faction_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module faction_faction_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module faction_faction_definition_progression_hash {
  source = "./attribute"

  name = "progression Hash"
  type = "int"
}
module faction_faction_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module faction_faction_definition_reward_item_hash {
  source = "./attribute"

  name = "reward Item Hash"
  type = "int64"
}
module faction_faction_definition_reward_vendor_hash {
  source = "./attribute"

  name = "reward Vendor Hash"
  type = "int64"
}
module faction_faction_definition_token_values {
  source = "./attribute"

  name = "token Values"
  type = "map[string]interface{}"
}
module faction_faction_definition_vendors {
  source = "./attribute"

  name = "vendors"
  type = "[]Vendors"
}


module faction_vendors {
  source = "./struct"

  name = "vendors"
  attributes_content = join("\n", [
    module.faction_vendors_background_image_path.rendered,
    module.faction_vendors_destination_hash.rendered,
    module.faction_vendors_vendor_hash.rendered,
  ])
}

module faction_vendors_background_image_path {
  source = "./attribute"

  name = "background Image Path"
  type = "string"
}
module faction_vendors_destination_hash {
  source = "./attribute"

  name = "destination Hash"
  type = "int64"
}
module faction_vendors_vendor_hash {
  source = "./attribute"

  name = "vendor Hash"
  type = "int64"
}


