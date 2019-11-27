// Code generated; DANGER ZONE FOR EDITS

module activity_interactable_definition {
  source = "./definition"

  target = var.target

  name = "activity Interactable"
  structs_content = join("\n", [
    module.activity_interactable_activity_interactable_definition.rendered,
    module.activity_interactable_activity_interactable_entries.rendered,
  ])

}

module activity_interactable_activity_interactable_definition {
  source = "./struct"

  name = "activity Interactable Definition"
  attributes_content = join("\n", [
    module.activity_interactable_activity_interactable_definition_blacklisted.rendered,
    module.activity_interactable_activity_interactable_definition_entries.rendered,
    module.activity_interactable_activity_interactable_definition_hash.rendered,
    module.activity_interactable_activity_interactable_definition_index.rendered,
    module.activity_interactable_activity_interactable_definition_redacted.rendered,
  ])
}

module activity_interactable_activity_interactable_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module activity_interactable_activity_interactable_definition_entries {
  source = "./attribute"

  name = "entries"
  type = "[]ActivityInteractableEntries"
}
module activity_interactable_activity_interactable_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module activity_interactable_activity_interactable_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module activity_interactable_activity_interactable_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


module activity_interactable_activity_interactable_entries {
  source = "./struct"

  name = "activity Interactable Entries"
  attributes_content = join("\n", [
    module.activity_interactable_activity_interactable_entries_activity_hash.rendered,
  ])
}

module activity_interactable_activity_interactable_entries_activity_hash {
  source = "./attribute"

  name = "activity Hash"
  type = "int"
}


