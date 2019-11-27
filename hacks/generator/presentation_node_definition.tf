// Code generated; DANGER ZONE FOR EDITS

module presentation_node_definition {
  source = "./definition"

  target = var.target

  name = "presentation Node"
  structs_content = join("\n", [
    module.presentation_node_presentation_node_definition.rendered,
    module.presentation_node_presentation_nodes.rendered,
    module.presentation_node_children.rendered,
    module.presentation_node_presentation_requirements.rendered,
  ])

}

module presentation_node_presentation_node_definition {
  source = "./struct"

  name = "presentation Node Definition"
  attributes_content = join("\n", [
    module.presentation_node_presentation_node_definition_blacklisted.rendered,
    module.presentation_node_presentation_node_definition_children.rendered,
    module.presentation_node_presentation_node_definition_disable_child_subscreen_navigation.rendered,
    module.presentation_node_presentation_node_definition_display_properties.rendered,
    module.presentation_node_presentation_node_definition_display_style.rendered,
    module.presentation_node_presentation_node_definition_hash.rendered,
    module.presentation_node_presentation_node_definition_index.rendered,
    module.presentation_node_presentation_node_definition_node_type.rendered,
    module.presentation_node_presentation_node_definition_objective_hash.rendered,
    module.presentation_node_presentation_node_definition_original_icon.rendered,
    module.presentation_node_presentation_node_definition_parent_node_hashes.rendered,
    module.presentation_node_presentation_node_definition_redacted.rendered,
    module.presentation_node_presentation_node_definition_requirements.rendered,
    module.presentation_node_presentation_node_definition_root_view_icon.rendered,
    module.presentation_node_presentation_node_definition_scope.rendered,
    module.presentation_node_presentation_node_definition_screen_style.rendered,
  ])
}

module presentation_node_presentation_node_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module presentation_node_presentation_node_definition_children {
  source = "./attribute"

  name = "children"
  type = "Children"
}
module presentation_node_presentation_node_definition_disable_child_subscreen_navigation {
  source = "./attribute"

  name = "disable Child Subscreen Navigation"
  type = "bool"
}
module presentation_node_presentation_node_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module presentation_node_presentation_node_definition_display_style {
  source = "./attribute"

  name = "display Style"
  type = "int"
}
module presentation_node_presentation_node_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module presentation_node_presentation_node_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module presentation_node_presentation_node_definition_node_type {
  source = "./attribute"

  name = "node Type"
  type = "int"
}
module presentation_node_presentation_node_definition_objective_hash {
  source = "./attribute"

  name = "objective Hash"
  type = "int"
}
module presentation_node_presentation_node_definition_original_icon {
  source = "./attribute"

  name = "original Icon"
  type = "string"
}
module presentation_node_presentation_node_definition_parent_node_hashes {
  source = "./attribute"

  name = "parent Node Hashes"
  type = "[]int64"
}
module presentation_node_presentation_node_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module presentation_node_presentation_node_definition_requirements {
  source = "./attribute"

  name = "requirements"
  type = "PresentationRequirements"
}
module presentation_node_presentation_node_definition_root_view_icon {
  source = "./attribute"

  name = "root View Icon"
  type = "string"
}
module presentation_node_presentation_node_definition_scope {
  source = "./attribute"

  name = "scope"
  type = "int"
}
module presentation_node_presentation_node_definition_screen_style {
  source = "./attribute"

  name = "screen Style"
  type = "int"
}


module presentation_node_presentation_nodes {
  source = "./struct"

  name = "presentation Nodes"
  attributes_content = join("\n", [
    module.presentation_node_presentation_nodes_presentation_node_hash.rendered,
  ])
}

module presentation_node_presentation_nodes_presentation_node_hash {
  source = "./attribute"

  name = "presentation Node Hash"
  type = "int64"
}


module presentation_node_children {
  source = "./struct"

  name = "children"
  attributes_content = join("\n", [
    module.presentation_node_children_collectibles.rendered,
    module.presentation_node_children_presentation_nodes.rendered,
    module.presentation_node_children_records.rendered,
  ])
}

module presentation_node_children_collectibles {
  source = "./attribute"

  name = "collectibles"
  type = "[]interface{}"
}
module presentation_node_children_presentation_nodes {
  source = "./attribute"

  name = "presentation Nodes"
  type = "[]PresentationNodes"
}
module presentation_node_children_records {
  source = "./attribute"

  name = "records"
  type = "[]interface{}"
}


module presentation_node_presentation_requirements {
  source = "./struct"

  name = "presentation Requirements"
  attributes_content = join("\n", [
    module.presentation_node_presentation_requirements_entitlement_unavailable_message.rendered,
  ])
}

module presentation_node_presentation_requirements_entitlement_unavailable_message {
  source = "./attribute"

  name = "entitlement Unavailable Message"
  type = "string"
}


