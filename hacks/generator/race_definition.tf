// Code generated; DANGER ZONE FOR EDITS

module race_definition {
  source = "./definition"

  target = var.target

  name = "race"
  structs_content = join("\n", [
    module.race_race_definition.rendered,
    module.race_gendered_race_names.rendered,
  ])

}

module race_race_definition {
  source = "./struct"

  name = "race Definition"
  attributes_content = join("\n", [
    module.race_race_definition_blacklisted.rendered,
    module.race_race_definition_display_properties.rendered,
    module.race_race_definition_gendered_race_names.rendered,
    module.race_race_definition_gendered_race_names_by_gender_hash.rendered,
    module.race_race_definition_hash.rendered,
    module.race_race_definition_index.rendered,
    module.race_race_definition_race_type.rendered,
    module.race_race_definition_redacted.rendered,
  ])
}

module race_race_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module race_race_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module race_race_definition_gendered_race_names {
  source = "./attribute"

  name = "gendered Race Names"
  type = "GenderedRaceNames"
}
module race_race_definition_gendered_race_names_by_gender_hash {
  source = "./attribute"

  name = "gendered Race Names By Gender Hash"
  type = "map[string]string"
}
module race_race_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module race_race_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module race_race_definition_race_type {
  source = "./attribute"

  name = "race Type"
  type = "int"
}
module race_race_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


module race_gendered_race_names {
  source = "./struct"

  name = "gendered Race Names"
  attributes_content = join("\n", [
    module.race_gendered_race_names_female.rendered,
    module.race_gendered_race_names_male.rendered,
  ])
}

module race_gendered_race_names_female {
  source = "./attribute"

  name = "Female"
  type = "string"
}
module race_gendered_race_names_male {
  source = "./attribute"

  name = "Male"
  type = "string"
}


