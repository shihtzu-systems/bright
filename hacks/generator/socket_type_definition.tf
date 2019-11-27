// Code generated; DANGER ZONE FOR EDITS

module socket_type_definition {
  source = "./definition"

  target = var.target

  name = "socket Type"
  structs_content = join("\n", [
    module.socket_type_socket_type_definition.rendered,
    module.socket_type_insert_action.rendered,
    module.socket_type_plug_whitelist.rendered,
  ])

}

module socket_type_socket_type_definition {
  source = "./struct"

  name = "socket Type Definition"
  attributes_content = join("\n", [
    module.socket_type_socket_type_definition_always_randomize_sockets.rendered,
    module.socket_type_socket_type_definition_avoid_duplicates_on_initialization.rendered,
    module.socket_type_socket_type_definition_blacklisted.rendered,
    module.socket_type_socket_type_definition_currency_scalars.rendered,
    module.socket_type_socket_type_definition_display_properties.rendered,
    module.socket_type_socket_type_definition_hash.rendered,
    module.socket_type_socket_type_definition_hide_duplicate_reusable_plugs.rendered,
    module.socket_type_socket_type_definition_index.rendered,
    module.socket_type_socket_type_definition_insert_action.rendered,
    module.socket_type_socket_type_definition_is_preview_enabled.rendered,
    module.socket_type_socket_type_definition_overrides_ui_appearance.rendered,
    module.socket_type_socket_type_definition_plug_whitelist.rendered,
    module.socket_type_socket_type_definition_redacted.rendered,
    module.socket_type_socket_type_definition_socket_category_hash.rendered,
    module.socket_type_socket_type_definition_visibility.rendered,
  ])
}

module socket_type_socket_type_definition_always_randomize_sockets {
  source = "./attribute"

  name = "always Randomize Sockets"
  type = "bool"
}
module socket_type_socket_type_definition_avoid_duplicates_on_initialization {
  source = "./attribute"

  name = "avoid Duplicates On Initialization"
  type = "bool"
}
module socket_type_socket_type_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module socket_type_socket_type_definition_currency_scalars {
  source = "./attribute"

  name = "currency Scalars"
  type = "[]interface{} "
}
module socket_type_socket_type_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module socket_type_socket_type_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module socket_type_socket_type_definition_hide_duplicate_reusable_plugs {
  source = "./attribute"

  name = "hide Duplicate Reusable Plugs"
  type = "bool"
}
module socket_type_socket_type_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module socket_type_socket_type_definition_insert_action {
  source = "./attribute"

  name = "insert Action"
  type = "InsertAction"
}
module socket_type_socket_type_definition_is_preview_enabled {
  source = "./attribute"

  name = "is Preview Enabled"
  type = "bool"
}
module socket_type_socket_type_definition_overrides_ui_appearance {
  source = "./attribute"

  name = "overrides Ui Appearance"
  type = "bool"
}
module socket_type_socket_type_definition_plug_whitelist {
  source = "./attribute"

  name = "plug Whitelist"
  type = "[]PlugWhitelist"
}
module socket_type_socket_type_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module socket_type_socket_type_definition_socket_category_hash {
  source = "./attribute"

  name = "socket Category Hash"
  type = "int"
}
module socket_type_socket_type_definition_visibility {
  source = "./attribute"

  name = "visibility"
  type = "int"
}


module socket_type_insert_action {
  source = "./struct"

  name = "insert Action"
  attributes_content = join("\n", [
    module.socket_type_insert_action_action_execute_seconds.rendered,
    module.socket_type_insert_action_action_sound_hash.rendered,
    module.socket_type_insert_action_action_type.rendered,
    module.socket_type_insert_action_is_positive_action.rendered,
  ])
}

module socket_type_insert_action_action_execute_seconds {
  source = "./attribute"

  name = "action Execute Seconds"
  type = "int"
}
module socket_type_insert_action_action_sound_hash {
  source = "./attribute"

  name = "action Sound Hash"
  type = "int"
}
module socket_type_insert_action_action_type {
  source = "./attribute"

  name = "action Type"
  type = "int"
}
module socket_type_insert_action_is_positive_action {
  source = "./attribute"

  name = "is Positive Action"
  type = "bool"
}


module socket_type_plug_whitelist {
  source = "./struct"

  name = "plug Whitelist"
  attributes_content = join("\n", [
    module.socket_type_plug_whitelist_category_hash.rendered,
    module.socket_type_plug_whitelist_category_identifier.rendered,
  ])
}

module socket_type_plug_whitelist_category_hash {
  source = "./attribute"

  name = "category Hash"
  type = "int"
}
module socket_type_plug_whitelist_category_identifier {
  source = "./attribute"

  name = "category Identifier"
  type = "string"
}


