// Code generated; DANGER ZONE FOR EDITS

module equipment_slot_definition {
  source = "./definition"

  target = var.target

  name = "equipment Slot"
  structs_content = join("\n", [
    module.equipment_slot_equipment_slot_definition.rendered,
    module.equipment_slot_art_dye_channels.rendered,
  ])

}

module equipment_slot_equipment_slot_definition {
  source = "./struct"

  name = "equipment Slot Definition"
  attributes_content = join("\n", [
    module.equipment_slot_equipment_slot_definition_apply_custom_art_dyes.rendered,
    module.equipment_slot_equipment_slot_definition_art_dye_channels.rendered,
    module.equipment_slot_equipment_slot_definition_blacklisted.rendered,
    module.equipment_slot_equipment_slot_definition_bucket_type_hash.rendered,
    module.equipment_slot_equipment_slot_definition_display_properties.rendered,
    module.equipment_slot_equipment_slot_definition_equipment_category_hash.rendered,
    module.equipment_slot_equipment_slot_definition_hash.rendered,
    module.equipment_slot_equipment_slot_definition_index.rendered,
    module.equipment_slot_equipment_slot_definition_redacted.rendered,
  ])
}

module equipment_slot_equipment_slot_definition_apply_custom_art_dyes {
  source = "./attribute"

  name = "apply Custom Art Dyes"
  type = "bool"
}
module equipment_slot_equipment_slot_definition_art_dye_channels {
  source = "./attribute"

  name = "art Dye Channels"
  type = "[]ArtDyeChannels"
}
module equipment_slot_equipment_slot_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module equipment_slot_equipment_slot_definition_bucket_type_hash {
  source = "./attribute"

  name = "bucket Type Hash"
  type = "int64"
}
module equipment_slot_equipment_slot_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module equipment_slot_equipment_slot_definition_equipment_category_hash {
  source = "./attribute"

  name = "equipment Category Hash"
  type = "int"
}
module equipment_slot_equipment_slot_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module equipment_slot_equipment_slot_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module equipment_slot_equipment_slot_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


module equipment_slot_art_dye_channels {
  source = "./struct"

  name = "art Dye Channels"
  attributes_content = join("\n", [
    module.equipment_slot_art_dye_channels_art_dye_channel_hash.rendered,
  ])
}

module equipment_slot_art_dye_channels_art_dye_channel_hash {
  source = "./attribute"

  name = "art Dye Channel Hash"
  type = "int"
}


