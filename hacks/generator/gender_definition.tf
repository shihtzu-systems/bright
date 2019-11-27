// Code generated; DANGER ZONE FOR EDITS

module gender_definition {
  source = "./definition"

  target = var.target

  name = "gender"
  structs_content = join("\n", [
    module.gender_gender_definition.rendered,
  ])

}

module gender_gender_definition {
  source = "./struct"

  name = "gender Definition"
  attributes_content = join("\n", [
    module.gender_gender_definition_blacklisted.rendered,
    module.gender_gender_definition_display_properties.rendered,
    module.gender_gender_definition_gender_type.rendered,
    module.gender_gender_definition_hash.rendered,
    module.gender_gender_definition_index.rendered,
    module.gender_gender_definition_redacted.rendered,
  ])
}

module gender_gender_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module gender_gender_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module gender_gender_definition_gender_type {
  source = "./attribute"

  name = "gender Type"
  type = "int"
}
module gender_gender_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module gender_gender_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module gender_gender_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


