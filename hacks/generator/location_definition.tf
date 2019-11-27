// Code generated; DANGER ZONE FOR EDITS

module location_definition {
  source = "./definition"

  target = var.target

  name = "location"
  structs_content = join("\n", [
    module.location_location_definition.rendered,
    module.location_location_releases.rendered,
  ])

}

module location_location_definition {
  source = "./struct"

  name = "location Definition"
  attributes_content = join("\n", [
    module.location_location_definition_blacklisted.rendered,
    module.location_location_definition_hash.rendered,
    module.location_location_definition_index.rendered,
    module.location_location_definition_location_releases.rendered,
    module.location_location_definition_redacted.rendered,
    module.location_location_definition_vendor_hash.rendered,
  ])
}

module location_location_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module location_location_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module location_location_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module location_location_definition_location_releases {
  source = "./attribute"

  name = "location Releases"
  type = "[]LocationReleases"
}
module location_location_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module location_location_definition_vendor_hash {
  source = "./attribute"

  name = "vendor Hash"
  type = "int"
}


module location_location_releases {
  source = "./struct"

  name = "location Releases"
  attributes_content = join("\n", [
    module.location_location_releases_activity_bubble_name.rendered,
    module.location_location_releases_activity_graph_hash.rendered,
    module.location_location_releases_activity_graph_node_hash.rendered,
    module.location_location_releases_activity_hash.rendered,
    module.location_location_releases_activity_path_bundle.rendered,
    module.location_location_releases_activity_path_destination.rendered,
    module.location_location_releases_destination_hash.rendered,
    module.location_location_releases_display_properties.rendered,
    module.location_location_releases_nav_point_type.rendered,
    module.location_location_releases_spawn_point.rendered,
    module.location_location_releases_world_position.rendered,
  ])
}

module location_location_releases_activity_bubble_name {
  source = "./attribute"

  name = "activity Bubble Name"
  type = "int64"
}
module location_location_releases_activity_graph_hash {
  source = "./attribute"

  name = "activity Graph Hash"
  type = "int64"
}
module location_location_releases_activity_graph_node_hash {
  source = "./attribute"

  name = "activity Graph Node Hash"
  type = "int"
}
module location_location_releases_activity_hash {
  source = "./attribute"

  name = "activity Hash"
  type = "int"
}
module location_location_releases_activity_path_bundle {
  source = "./attribute"

  name = "activity Path Bundle"
  type = "int"
}
module location_location_releases_activity_path_destination {
  source = "./attribute"

  name = "activity Path Destination"
  type = "int"
}
module location_location_releases_destination_hash {
  source = "./attribute"

  name = "destination Hash"
  type = "int"
}
module location_location_releases_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module location_location_releases_nav_point_type {
  source = "./attribute"

  name = "nav Point Type"
  type = "int"
}
module location_location_releases_spawn_point {
  source = "./attribute"

  name = "spawn Point"
  type = "int"
}
module location_location_releases_world_position {
  source = "./attribute"

  name = "world Position"
  type = "[]int"
}


