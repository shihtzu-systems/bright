// Code generated; DANGER ZONE FOR EDITS

module vendor_group_definition {
  source = "./definition"

  target = var.target

  name = "vendor Group"
  structs_content = join("\n", [
    module.vendor_group_vendor_group_definition.rendered,
  ])

}

module vendor_group_vendor_group_definition {
  source = "./struct"

  name = "vendor Group Definition"
  attributes_content = join("\n", [
    module.vendor_group_vendor_group_definition_blacklisted.rendered,
    module.vendor_group_vendor_group_definition_category_name.rendered,
    module.vendor_group_vendor_group_definition_hash.rendered,
    module.vendor_group_vendor_group_definition_index.rendered,
    module.vendor_group_vendor_group_definition_order.rendered,
    module.vendor_group_vendor_group_definition_redacted.rendered,
  ])
}

module vendor_group_vendor_group_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module vendor_group_vendor_group_definition_category_name {
  source = "./attribute"

  name = "category Name"
  type = "string"
}
module vendor_group_vendor_group_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module vendor_group_vendor_group_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module vendor_group_vendor_group_definition_order {
  source = "./attribute"

  name = "order"
  type = "int"
}
module vendor_group_vendor_group_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


