// Code generated; DANGER ZONE FOR EDITS

module checklist_definition {
  source = "./definition"

  target = var.target

  name = "checklist"
  structs_content = join("\n", [
    module.checklist_checklist_definition.rendered,
    module.checklist_checklist_entries.rendered,
  ])

}

module checklist_checklist_definition {
  source = "./struct"

  name = "checklist Definition"
  attributes_content = join("\n", [
    module.checklist_checklist_definition_blacklisted.rendered,
    module.checklist_checklist_definition_display_properties.rendered,
    module.checklist_checklist_definition_entries.rendered,
    module.checklist_checklist_definition_hash.rendered,
    module.checklist_checklist_definition_index.rendered,
    module.checklist_checklist_definition_redacted.rendered,
    module.checklist_checklist_definition_scope.rendered,
    module.checklist_checklist_definition_view_action_string.rendered,
  ])
}

module checklist_checklist_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module checklist_checklist_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module checklist_checklist_definition_entries {
  source = "./attribute"

  name = "entries"
  type = "[]ChecklistEntries"
}
module checklist_checklist_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module checklist_checklist_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module checklist_checklist_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module checklist_checklist_definition_scope {
  source = "./attribute"

  name = "scope"
  type = "int"
}
module checklist_checklist_definition_view_action_string {
  source = "./attribute"

  name = "view Action String"
  type = "string"
}


module checklist_checklist_entries {
  source = "./struct"

  name = "checklist Entries"
  attributes_content = join("\n", [
    module.checklist_checklist_entries_bubble_hash.rendered,
    module.checklist_checklist_entries_destination_hash.rendered,
    module.checklist_checklist_entries_display_properties.rendered,
    module.checklist_checklist_entries_hash.rendered,
    module.checklist_checklist_entries_location_hash.rendered,
    module.checklist_checklist_entries_scope.rendered,
  ])
}

module checklist_checklist_entries_bubble_hash {
  source = "./attribute"

  name = "bubble Hash"
  type = "int64"
}
module checklist_checklist_entries_destination_hash {
  source = "./attribute"

  name = "destination Hash"
  type = "int64"
}
module checklist_checklist_entries_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module checklist_checklist_entries_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module checklist_checklist_entries_location_hash {
  source = "./attribute"

  name = "location Hash"
  type = "int"
}
module checklist_checklist_entries_scope {
  source = "./attribute"

  name = "scope"
  type = "int"
}


