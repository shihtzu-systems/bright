// Code generated; DANGER ZONE FOR EDITS

module art_dye_reference_definition {
  source = "./definition"

  target = var.target

  name = "art Dye Reference"
  structs_content = join("\n", [
    module.art_dye_reference_art_dye_reference_definition.rendered,
  ])

}

module art_dye_reference_art_dye_reference_definition {
  source = "./struct"

  name = "art Dye Reference Definition"
  attributes_content = join("\n", [
    module.art_dye_reference_art_dye_reference_definition_art_dye_hash.rendered,
    module.art_dye_reference_art_dye_reference_definition_blacklisted.rendered,
    module.art_dye_reference_art_dye_reference_definition_dye_manifest_hash.rendered,
    module.art_dye_reference_art_dye_reference_definition_hash.rendered,
    module.art_dye_reference_art_dye_reference_definition_index.rendered,
    module.art_dye_reference_art_dye_reference_definition_redacted.rendered,
  ])
}

module art_dye_reference_art_dye_reference_definition_art_dye_hash {
  source = "./attribute"

  name = "art Dye Hash"
  type = "int"
}
module art_dye_reference_art_dye_reference_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module art_dye_reference_art_dye_reference_definition_dye_manifest_hash {
  source = "./attribute"

  name = "dye Manifest Hash"
  type = "int64"
}
module art_dye_reference_art_dye_reference_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module art_dye_reference_art_dye_reference_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module art_dye_reference_art_dye_reference_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


