// Code generated; DANGER ZONE FOR EDITS

module node_step_summary_definition {
  source = "./definition"

  target = var.target

  name = "node Step Summary"
  structs_content = join("\n", [
    module.node_step_summary_node_step_summary_definition.rendered,
  ])

}

module node_step_summary_node_step_summary_definition {
  source = "./struct"

  name = "node Step Summary Definition"
  attributes_content = join("\n", [
  ])
}



