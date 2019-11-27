// Code generated; DANGER ZONE FOR EDITS

module material_requirement_set_definition {
  source = "./definition"

  target = var.target

  name = "material Requirement Set"
  structs_content = join("\n", [
    module.material_requirement_set_material_requirement_set_definition.rendered,
    module.material_requirement_set_materials.rendered,
  ])

}

module material_requirement_set_material_requirement_set_definition {
  source = "./struct"

  name = "material Requirement Set Definition"
  attributes_content = join("\n", [
    module.material_requirement_set_material_requirement_set_definition_blacklisted.rendered,
    module.material_requirement_set_material_requirement_set_definition_hash.rendered,
    module.material_requirement_set_material_requirement_set_definition_index.rendered,
    module.material_requirement_set_material_requirement_set_definition_materials.rendered,
    module.material_requirement_set_material_requirement_set_definition_redacted.rendered,
  ])
}

module material_requirement_set_material_requirement_set_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module material_requirement_set_material_requirement_set_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module material_requirement_set_material_requirement_set_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module material_requirement_set_material_requirement_set_definition_materials {
  source = "./attribute"

  name = "materials"
  type = "[]Materials"
}
module material_requirement_set_material_requirement_set_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


module material_requirement_set_materials {
  source = "./struct"

  name = "materials"
  attributes_content = join("\n", [
    module.material_requirement_set_materials_count.rendered,
    module.material_requirement_set_materials_delete_on_action.rendered,
    module.material_requirement_set_materials_item_hash.rendered,
    module.material_requirement_set_materials_omit_from_requirements.rendered,
  ])
}

module material_requirement_set_materials_count {
  source = "./attribute"

  name = "count"
  type = "int"
}
module material_requirement_set_materials_delete_on_action {
  source = "./attribute"

  name = "delete On Action"
  type = "bool"
}
module material_requirement_set_materials_item_hash {
  source = "./attribute"

  name = "item Hash"
  type = "int64"
}
module material_requirement_set_materials_omit_from_requirements {
  source = "./attribute"

  name = "omit From Requirements"
  type = "bool"
}


