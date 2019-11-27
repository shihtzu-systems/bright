// Code generated; DANGER ZONE FOR EDITS

module unlock_event_definition {
  source = "./definition"

  target = var.target

  name = "unlock Event"
  structs_content = join("\n", [
    module.unlock_event_unlock_event_definition.rendered,
    module.unlock_event_unlock_entries.rendered,
  ])

}

module unlock_event_unlock_event_definition {
  source = "./struct"

  name = "unlock Event Definition"
  attributes_content = join("\n", [
    module.unlock_event_unlock_event_definition_blacklisted.rendered,
    module.unlock_event_unlock_event_definition_hash.rendered,
    module.unlock_event_unlock_event_definition_index.rendered,
    module.unlock_event_unlock_event_definition_new_sequence_reward_site_hash.rendered,
    module.unlock_event_unlock_event_definition_redacted.rendered,
    module.unlock_event_unlock_event_definition_sequence_last_updated_unlock_value_hash.rendered,
    module.unlock_event_unlock_event_definition_sequence_unlock_value_hash.rendered,
    module.unlock_event_unlock_event_definition_unlock_entries.rendered,
  ])
}

module unlock_event_unlock_event_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module unlock_event_unlock_event_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module unlock_event_unlock_event_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module unlock_event_unlock_event_definition_new_sequence_reward_site_hash {
  source = "./attribute"

  name = "new Sequence Reward Site Hash"
  type = "int"
}
module unlock_event_unlock_event_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module unlock_event_unlock_event_definition_sequence_last_updated_unlock_value_hash {
  source = "./attribute"

  name = "sequence Last Updated Unlock Value Hash"
  type = "int"
}
module unlock_event_unlock_event_definition_sequence_unlock_value_hash {
  source = "./attribute"

  name = "sequence Unlock Value Hash"
  type = "int"
}
module unlock_event_unlock_event_definition_unlock_entries {
  source = "./attribute"

  name = "unlock Entries"
  type = "[]UnlockEntries"
}


module unlock_event_unlock_entries {
  source = "./struct"

  name = "unlock Entries"
  attributes_content = join("\n", [
    module.unlock_event_unlock_entries_cleared_value.rendered,
    module.unlock_event_unlock_entries_selected_value.rendered,
    module.unlock_event_unlock_entries_unlock_hash.rendered,
    module.unlock_event_unlock_entries_unlock_value_hash.rendered,
  ])
}

module unlock_event_unlock_entries_cleared_value {
  source = "./attribute"

  name = "cleared Value"
  type = "float64"
}
module unlock_event_unlock_entries_selected_value {
  source = "./attribute"

  name = "selected Value"
  type = "float64"
}
module unlock_event_unlock_entries_unlock_hash {
  source = "./attribute"

  name = "unlock Hash"
  type = "float64"
}
module unlock_event_unlock_entries_unlock_value_hash {
  source = "./attribute"

  name = "unlock Value Hash"
  type = "float64"
}


