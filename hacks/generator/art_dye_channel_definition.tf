// Code generated; DANGER ZONE FOR EDITS

module art_dye_channel_definition {
  source = "./definition"

  target = var.target

  name = "art Dye Channel"
  structs_content = join("\n", [
    module.art_dye_channel_art_dye_channel_definition.rendered,
  ])

}

module art_dye_channel_art_dye_channel_definition {
  source = "./struct"

  name = "art Dye Channel Definition"
  attributes_content = join("\n", [
    module.art_dye_channel_art_dye_channel_definition_blacklisted.rendered,
    module.art_dye_channel_art_dye_channel_definition_channel_hash.rendered,
    module.art_dye_channel_art_dye_channel_definition_hash.rendered,
    module.art_dye_channel_art_dye_channel_definition_index.rendered,
    module.art_dye_channel_art_dye_channel_definition_redacted.rendered,
  ])
}

module art_dye_channel_art_dye_channel_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module art_dye_channel_art_dye_channel_definition_channel_hash {
  source = "./attribute"

  name = "channel Hash"
  type = "int64"
}
module art_dye_channel_art_dye_channel_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module art_dye_channel_art_dye_channel_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module art_dye_channel_art_dye_channel_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


