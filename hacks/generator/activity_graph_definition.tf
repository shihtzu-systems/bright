// Code generated; DANGER ZONE FOR EDITS

module activity_graph_definition {
  source = "./definition"

  target = var.target

  name = "activity Graph"
  structs_content = join("\n", [
    module.activity_graph_activity_graph_definition.rendered,
  ])

}

module activity_graph_activity_graph_definition {
  source = "./struct"

  name = "activity Graph Definition"
  attributes_content = join("\n", [
    module.activity_graph_activity_graph_definition_art_elements.rendered,
    module.activity_graph_activity_graph_definition_blacklisted.rendered,
    module.activity_graph_activity_graph_definition_connections.rendered,
    module.activity_graph_activity_graph_definition_display_objectives.rendered,
    module.activity_graph_activity_graph_definition_display_progressions.rendered,
    module.activity_graph_activity_graph_definition_hash.rendered,
    module.activity_graph_activity_graph_definition_ignore_for_milestones.rendered,
    module.activity_graph_activity_graph_definition_index.rendered,
    module.activity_graph_activity_graph_definition_linked_graphs.rendered,
    module.activity_graph_activity_graph_definition_nodes.rendered,
    module.activity_graph_activity_graph_definition_redacted.rendered,
    module.activity_graph_activity_graph_definition_ui_screen.rendered,
  ])
}

module activity_graph_activity_graph_definition_art_elements {
  source = "./attribute"

  name = "art Elements"
  type = "[]interface{}"
}
module activity_graph_activity_graph_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module activity_graph_activity_graph_definition_connections {
  source = "./attribute"

  name = "connections"
  type = "[]interface{}"
}
module activity_graph_activity_graph_definition_display_objectives {
  source = "./attribute"

  name = "display Objectives"
  type = "[]interface{}"
}
module activity_graph_activity_graph_definition_display_progressions {
  source = "./attribute"

  name = "display Progressions"
  type = "[]interface{}"
}
module activity_graph_activity_graph_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module activity_graph_activity_graph_definition_ignore_for_milestones {
  source = "./attribute"

  name = "ignore For Milestones"
  type = "bool"
}
module activity_graph_activity_graph_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module activity_graph_activity_graph_definition_linked_graphs {
  source = "./attribute"

  name = "linked Graphs"
  type = "[]interface{}"
}
module activity_graph_activity_graph_definition_nodes {
  source = "./attribute"

  name = "nodes"
  type = "[]interface{}"
}
module activity_graph_activity_graph_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module activity_graph_activity_graph_definition_ui_screen {
  source = "./attribute"

  name = "ui Screen"
  type = "int"
}


