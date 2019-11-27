// Code generated; DANGER ZONE FOR EDITS

module plug_set_definition {
  source = "./definition"

  target = var.target

  name = "plug Set"
  structs_content = join("\n", [
    module.plug_set_plug_set_definition.rendered,
    module.plug_set_reusable_plug_items.rendered,
  ])

}

module plug_set_plug_set_definition {
  source = "./struct"

  name = "plug Set Definition"
  attributes_content = join("\n", [
    module.plug_set_plug_set_definition_blacklisted.rendered,
    module.plug_set_plug_set_definition_display_properties.rendered,
    module.plug_set_plug_set_definition_hash.rendered,
    module.plug_set_plug_set_definition_index.rendered,
    module.plug_set_plug_set_definition_redacted.rendered,
    module.plug_set_plug_set_definition_reusable_plug_items.rendered,
  ])
}

module plug_set_plug_set_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module plug_set_plug_set_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module plug_set_plug_set_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module plug_set_plug_set_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module plug_set_plug_set_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module plug_set_plug_set_definition_reusable_plug_items {
  source = "./attribute"

  name = "reusable Plug Items"
  type = "[]ReusablePlugItems"
}


module plug_set_reusable_plug_items {
  source = "./struct"

  name = "reusable Plug Items"
  attributes_content = join("\n", [
    module.plug_set_reusable_plug_items_alternate_weight.rendered,
    module.plug_set_reusable_plug_items_plug_item_hash.rendered,
    module.plug_set_reusable_plug_items_weight.rendered,
  ])
}

module plug_set_reusable_plug_items_alternate_weight {
  source = "./attribute"

  name = "alternate Weight"
  type = "float64"
}
module plug_set_reusable_plug_items_plug_item_hash {
  source = "./attribute"

  name = "plug Item Hash"
  type = "float64"
}
module plug_set_reusable_plug_items_weight {
  source = "./attribute"

  name = "weight"
  type = "float64"
}


