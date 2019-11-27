// Code generated; DANGER ZONE FOR EDITS

module item_category_definition {
  source = "./definition"

  target = var.target

  name = "item Category"
  structs_content = join("\n", [
    module.item_category_item_category_definition.rendered,
  ])

}

module item_category_item_category_definition {
  source = "./struct"

  name = "item Category Definition"
  attributes_content = join("\n", [
    module.item_category_item_category_definition_blacklisted.rendered,
    module.item_category_item_category_definition_deprecated.rendered,
    module.item_category_item_category_definition_display_properties.rendered,
    module.item_category_item_category_definition_grant_destiny_breaker_type.rendered,
    module.item_category_item_category_definition_grant_destiny_class.rendered,
    module.item_category_item_category_definition_grant_destiny_item_type.rendered,
    module.item_category_item_category_definition_grant_destiny_sub_type.rendered,
    module.item_category_item_category_definition_group_category_only.rendered,
    module.item_category_item_category_definition_grouped_category_hashes.rendered,
    module.item_category_item_category_definition_hash.rendered,
    module.item_category_item_category_definition_index.rendered,
    module.item_category_item_category_definition_is_plug.rendered,
    module.item_category_item_category_definition_parent_category_hashes.rendered,
    module.item_category_item_category_definition_redacted.rendered,
    module.item_category_item_category_definition_short_title.rendered,
    module.item_category_item_category_definition_visible.rendered,
  ])
}

module item_category_item_category_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module item_category_item_category_definition_deprecated {
  source = "./attribute"

  name = "deprecated"
  type = "bool"
}
module item_category_item_category_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module item_category_item_category_definition_grant_destiny_breaker_type {
  source = "./attribute"

  name = "grant Destiny Breaker Type"
  type = "int"
}
module item_category_item_category_definition_grant_destiny_class {
  source = "./attribute"

  name = "grant Destiny Class"
  type = "int"
}
module item_category_item_category_definition_grant_destiny_item_type {
  source = "./attribute"

  name = "grant Destiny Item Type"
  type = "int"
}
module item_category_item_category_definition_grant_destiny_sub_type {
  source = "./attribute"

  name = "grant Destiny Sub Type"
  type = "int"
}
module item_category_item_category_definition_group_category_only {
  source = "./attribute"

  name = "group Category Only"
  type = "bool"
}
module item_category_item_category_definition_grouped_category_hashes {
  source = "./attribute"

  name = "grouped Category Hashes"
  type = "[]interface{}"
}
module item_category_item_category_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module item_category_item_category_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module item_category_item_category_definition_is_plug {
  source = "./attribute"

  name = "is Plug"
  type = "bool"
}
module item_category_item_category_definition_parent_category_hashes {
  source = "./attribute"

  name = "parent Category Hashes"
  type = "[]int"
}
module item_category_item_category_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module item_category_item_category_definition_short_title {
  source = "./attribute"

  name = "short Title"
  type = "string"
}
module item_category_item_category_definition_visible {
  source = "./attribute"

  name = "visible"
  type = "bool"
}


