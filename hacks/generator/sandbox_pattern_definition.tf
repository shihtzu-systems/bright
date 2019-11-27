// Code generated; DANGER ZONE FOR EDITS

module sandbox_pattern_definition {
  source = "./definition"

  target = var.target

  name = "sandbox Pattern"
  structs_content = join("\n", [
    module.sandbox_pattern_sandbox_pattern_definition.rendered,
    module.sandbox_pattern_filter.rendered,
  ])

}

module sandbox_pattern_sandbox_pattern_definition {
  source = "./struct"

  name = "sandbox Pattern Definition"
  attributes_content = join("\n", [
    module.sandbox_pattern_sandbox_pattern_definition_blacklisted.rendered,
    module.sandbox_pattern_sandbox_pattern_definition_filters.rendered,
    module.sandbox_pattern_sandbox_pattern_definition_hash.rendered,
    module.sandbox_pattern_sandbox_pattern_definition_index.rendered,
    module.sandbox_pattern_sandbox_pattern_definition_pattern_global_tag_id_hash.rendered,
    module.sandbox_pattern_sandbox_pattern_definition_pattern_hash.rendered,
    module.sandbox_pattern_sandbox_pattern_definition_redacted.rendered,
    module.sandbox_pattern_sandbox_pattern_definition_weapon_content_group_hash.rendered,
    module.sandbox_pattern_sandbox_pattern_definition_weapon_translation_group_hash.rendered,
    module.sandbox_pattern_sandbox_pattern_definition_weapon_type.rendered,
    module.sandbox_pattern_sandbox_pattern_definition_weapon_type_hash.rendered,
  ])
}

module sandbox_pattern_sandbox_pattern_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module sandbox_pattern_sandbox_pattern_definition_filters {
  source = "./attribute"

  name = "filters"
  type = "[]Filter"
}
module sandbox_pattern_sandbox_pattern_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module sandbox_pattern_sandbox_pattern_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module sandbox_pattern_sandbox_pattern_definition_pattern_global_tag_id_hash {
  source = "./attribute"

  name = "pattern Global Tag Id Hash"
  type = "int64"
}
module sandbox_pattern_sandbox_pattern_definition_pattern_hash {
  source = "./attribute"

  name = "pattern Hash"
  type = "int64"
}
module sandbox_pattern_sandbox_pattern_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module sandbox_pattern_sandbox_pattern_definition_weapon_content_group_hash {
  source = "./attribute"

  name = "weapon Content Group Hash"
  type = "int64"
}
module sandbox_pattern_sandbox_pattern_definition_weapon_translation_group_hash {
  source = "./attribute"

  name = "weapon Translation Group Hash"
  type = "int64"
}
module sandbox_pattern_sandbox_pattern_definition_weapon_type {
  source = "./attribute"

  name = "weapon Type"
  type = "int"
}
module sandbox_pattern_sandbox_pattern_definition_weapon_type_hash {
  source = "./attribute"

  name = "weapon Type Hash"
  type = "int64"
}


module sandbox_pattern_filter {
  source = "./struct"

  name = "filter"
  attributes_content = join("\n", [
    module.sandbox_pattern_filter_arrangement_index_by_stat_value.rendered,
    module.sandbox_pattern_filter_art_arrangement_region_hash.rendered,
    module.sandbox_pattern_filter_art_arrangement_region_index.rendered,
    module.sandbox_pattern_filter_stat_hash.rendered,
  ])
}

module sandbox_pattern_filter_arrangement_index_by_stat_value {
  source = "./attribute"

  name = "arrangement Index By Stat Value"
  type = "map[string]int"
}
module sandbox_pattern_filter_art_arrangement_region_hash {
  source = "./attribute"

  name = "art Arrangement Region Hash"
  type = "int64"
}
module sandbox_pattern_filter_art_arrangement_region_index {
  source = "./attribute"

  name = "art Arrangement Region Index"
  type = "int"
}
module sandbox_pattern_filter_stat_hash {
  source = "./attribute"

  name = "stat Hash"
  type = "int64"
}


