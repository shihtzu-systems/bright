// Code generated; DANGER ZONE FOR EDITS

module inventory_bucket_definition {
  source = "./definition"

  target = var.target

  name = "inventory Bucket"
  structs_content = join("\n", [
    module.inventory_bucket_inventory_bucket_definition.rendered,
  ])

}

module inventory_bucket_inventory_bucket_definition {
  source = "./struct"

  name = "inventory Bucket Definition"
  attributes_content = join("\n", [
    module.inventory_bucket_inventory_bucket_definition_blacklisted.rendered,
    module.inventory_bucket_inventory_bucket_definition_bucketorder.rendered,
    module.inventory_bucket_inventory_bucket_definition_category.rendered,
    module.inventory_bucket_inventory_bucket_definition_display_properties.rendered,
    module.inventory_bucket_inventory_bucket_definition_enabled.rendered,
    module.inventory_bucket_inventory_bucket_definition_fifo.rendered,
    module.inventory_bucket_inventory_bucket_definition_has_transfer_destination.rendered,
    module.inventory_bucket_inventory_bucket_definition_hash.rendered,
    module.inventory_bucket_inventory_bucket_definition_index.rendered,
    module.inventory_bucket_inventory_bucket_definition_item_count.rendered,
    module.inventory_bucket_inventory_bucket_definition_location.rendered,
    module.inventory_bucket_inventory_bucket_definition_redacted.rendered,
    module.inventory_bucket_inventory_bucket_definition_scope.rendered,
  ])
}

module inventory_bucket_inventory_bucket_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module inventory_bucket_inventory_bucket_definition_bucketorder {
  source = "./attribute"

  name = "bucketOrder"
  type = "int"
}
module inventory_bucket_inventory_bucket_definition_category {
  source = "./attribute"

  name = "category"
  type = "int"
}
module inventory_bucket_inventory_bucket_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module inventory_bucket_inventory_bucket_definition_enabled {
  source = "./attribute"

  name = "enabled"
  type = "bool"
}
module inventory_bucket_inventory_bucket_definition_fifo {
  source = "./attribute"

  name = "fifo"
  type = "bool"
}
module inventory_bucket_inventory_bucket_definition_has_transfer_destination {
  source = "./attribute"

  name = "has Transfer Destination"
  type = "bool"
}
module inventory_bucket_inventory_bucket_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module inventory_bucket_inventory_bucket_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module inventory_bucket_inventory_bucket_definition_item_count {
  source = "./attribute"

  name = "item Count"
  type = "int"
}
module inventory_bucket_inventory_bucket_definition_location {
  source = "./attribute"

  name = "location"
  type = "int"
}
module inventory_bucket_inventory_bucket_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module inventory_bucket_inventory_bucket_definition_scope {
  source = "./attribute"

  name = "scope"
  type = "int"
}


