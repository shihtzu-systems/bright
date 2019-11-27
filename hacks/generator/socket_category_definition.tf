// Code generated; DANGER ZONE FOR EDITS

module socket_category_definition {
  source = "./definition"

  target = var.target

  name = "socket Category"
  structs_content = join("\n", [
    module.socket_category_socket_category_definition.rendered,
  ])

}

module socket_category_socket_category_definition {
  source = "./struct"

  name = "socket Category Definition"
  attributes_content = join("\n", [
    module.socket_category_socket_category_definition_blacklisted.rendered,
    module.socket_category_socket_category_definition_category_style.rendered,
    module.socket_category_socket_category_definition_display_properties.rendered,
    module.socket_category_socket_category_definition_hash.rendered,
    module.socket_category_socket_category_definition_index.rendered,
    module.socket_category_socket_category_definition_redacted.rendered,
    module.socket_category_socket_category_definition_ui_category_style.rendered,
  ])
}

module socket_category_socket_category_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module socket_category_socket_category_definition_category_style {
  source = "./attribute"

  name = "category Style"
  type = "int"
}
module socket_category_socket_category_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module socket_category_socket_category_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module socket_category_socket_category_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module socket_category_socket_category_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module socket_category_socket_category_definition_ui_category_style {
  source = "./attribute"

  name = "ui Category Style"
  type = "int"
}


