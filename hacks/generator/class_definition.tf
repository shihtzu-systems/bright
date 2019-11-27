// Code generated; DANGER ZONE FOR EDITS

module class_definition {
  source = "./definition"

  target = var.target

  name = "class"
  structs_content = join("\n", [
    module.class_class_definition.rendered,
    module.class_gendered_class_names.rendered,
  ])

}

module class_class_definition {
  source = "./struct"

  name = "class Definition"
  attributes_content = join("\n", [
    module.class_class_definition_blacklisted.rendered,
    module.class_class_definition_class_type.rendered,
    module.class_class_definition_display_properties.rendered,
    module.class_class_definition_gendered_class_names.rendered,
    module.class_class_definition_gendered_class_names_by_gender_hash.rendered,
    module.class_class_definition_hash.rendered,
    module.class_class_definition_index.rendered,
    module.class_class_definition_redacted.rendered,
  ])
}

module class_class_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module class_class_definition_class_type {
  source = "./attribute"

  name = "class Type"
  type = "int"
}
module class_class_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module class_class_definition_gendered_class_names {
  source = "./attribute"

  name = "gendered Class Names"
  type = "GenderedClassNames"
}
module class_class_definition_gendered_class_names_by_gender_hash {
  source = "./attribute"

  name = "gendered Class Names By Gender Hash"
  type = "map[string]interface{}"
}
module class_class_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module class_class_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module class_class_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


module class_gendered_class_names {
  source = "./struct"

  name = "gendered class Names"
  attributes_content = join("\n", [
    module.class_gendered_class_names_female.rendered,
    module.class_gendered_class_names_male.rendered,
  ])
}

module class_gendered_class_names_female {
  source = "./attribute"

  name = "female"
  type = "string"
}
module class_gendered_class_names_male {
  source = "./attribute"

  name = "male"
  type = "string"
}


