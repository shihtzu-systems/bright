// Code generated; DANGER ZONE FOR EDITS

module destination_definition {
  source = "./definition"

  target = var.target

  name = "destination"
  structs_content = join("\n", [
    module.destination_destination_definition.rendered,
  ])

}

module destination_destination_definition {
  source = "./struct"

  name = "destination Definition"
  attributes_content = join("\n", [
    module.destination_destination_definition_activity_graph_entries.rendered,
    module.destination_destination_definition_blacklisted.rendered,
    module.destination_destination_definition_bubble_settings.rendered,
    module.destination_destination_definition_bubbles.rendered,
    module.destination_destination_definition_default_freeroam_activity_hash.rendered,
    module.destination_destination_definition_display_properties.rendered,
    module.destination_destination_definition_hash.rendered,
    module.destination_destination_definition_index.rendered,
    module.destination_destination_definition_place_hash.rendered,
    module.destination_destination_definition_redacted.rendered,
  ])
}

module destination_destination_definition_activity_graph_entries {
  source = "./attribute"

  name = "activity Graph Entries"
  type = "[]interface{} "
}
module destination_destination_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module destination_destination_definition_bubble_settings {
  source = "./attribute"

  name = "bubble Settings"
  type = "[]interface{}"
}
module destination_destination_definition_bubbles {
  source = "./attribute"

  name = "bubbles"
  type = "[]interface{}"
}
module destination_destination_definition_default_freeroam_activity_hash {
  source = "./attribute"

  name = "default Freeroam Activity Hash"
  type = "int"
}
module destination_destination_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module destination_destination_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module destination_destination_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module destination_destination_definition_place_hash {
  source = "./attribute"

  name = "place Hash"
  type = "int64"
}
module destination_destination_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


