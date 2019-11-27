// Code generated; DANGER ZONE FOR EDITS

module talent_grid_definition {
  source = "./definition"

  target = var.target

  name = "talent Grid"
  structs_content = join("\n", [
    module.talent_grid_talent_grid_definition.rendered,
    module.talent_grid_group.rendered,
    module.talent_grid_exclusive_set.rendered,
    module.talent_grid_node_category.rendered,
    module.talent_grid_activation_requirement.rendered,
    module.talent_grid_node.rendered,
    module.talent_grid_step.rendered,
  ])

}

module talent_grid_talent_grid_definition {
  source = "./struct"

  name = "talent Grid Definition"
  attributes_content = join("\n", [
    module.talent_grid_talent_grid_definition_blacklisted.rendered,
    module.talent_grid_talent_grid_definition_calc_max_grid_level.rendered,
    module.talent_grid_talent_grid_definition_calc_progress_to_max_level.rendered,
    module.talent_grid_talent_grid_definition_exclusive_sets.rendered,
    module.talent_grid_talent_grid_definition_grid_level_per_column.rendered,
    module.talent_grid_talent_grid_definition_groups.rendered,
    module.talent_grid_talent_grid_definition_hash.rendered,
    module.talent_grid_talent_grid_definition_independent_node_indexes.rendered,
    module.talent_grid_talent_grid_definition_index.rendered,
    module.talent_grid_talent_grid_definition_max_grid_level.rendered,
    module.talent_grid_talent_grid_definition_maximum_random_material_requirements.rendered,
    module.talent_grid_talent_grid_definition_node_categories.rendered,
    module.talent_grid_talent_grid_definition_nodes.rendered,
    module.talent_grid_talent_grid_definition_progression_hash.rendered,
    module.talent_grid_talent_grid_definition_redacted.rendered,
  ])
}

module talent_grid_talent_grid_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module talent_grid_talent_grid_definition_calc_max_grid_level {
  source = "./attribute"

  name = "calc Max Grid Level"
  type = "int"
}
module talent_grid_talent_grid_definition_calc_progress_to_max_level {
  source = "./attribute"

  name = "calc Progress To Max Level"
  type = "int"
}
module talent_grid_talent_grid_definition_exclusive_sets {
  source = "./attribute"

  name = "exclusive Sets"
  type = "[]ExclusiveSet"
}
module talent_grid_talent_grid_definition_grid_level_per_column {
  source = "./attribute"

  name = "grid Level Per Column"
  type = "int"
}
module talent_grid_talent_grid_definition_groups {
  source = "./attribute"

  name = "groups"
  type = "map[string]Group"
}
module talent_grid_talent_grid_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module talent_grid_talent_grid_definition_independent_node_indexes {
  source = "./attribute"

  name = "independent Node Indexes"
  type = "[]float64"
}
module talent_grid_talent_grid_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module talent_grid_talent_grid_definition_max_grid_level {
  source = "./attribute"

  name = "max Grid Level"
  type = "int"
}
module talent_grid_talent_grid_definition_maximum_random_material_requirements {
  source = "./attribute"

  name = "maximum Random Material Requirements"
  type = "int"
}
module talent_grid_talent_grid_definition_node_categories {
  source = "./attribute"

  name = "node Categories"
  type = "[]NodeCategory"
}
module talent_grid_talent_grid_definition_nodes {
  source = "./attribute"

  name = "nodes"
  type = "[]Node"
}
module talent_grid_talent_grid_definition_progression_hash {
  source = "./attribute"

  name = "progression Hash"
  type = "int"
}
module talent_grid_talent_grid_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


module talent_grid_group {
  source = "./struct"

  name = "group"
  attributes_content = join("\n", [
    module.talent_grid_group_group_hash.rendered,
    module.talent_grid_group_node_hashes.rendered,
    module.talent_grid_group_opposing_group_hashes.rendered,
    module.talent_grid_group_opposing_node_hashes.rendered,
  ])
}

module talent_grid_group_group_hash {
  source = "./attribute"

  name = "group Hash"
  type = "int64"
}
module talent_grid_group_node_hashes {
  source = "./attribute"

  name = "node Hashes"
  type = "[]int64"
}
module talent_grid_group_opposing_group_hashes {
  source = "./attribute"

  name = "opposing Group Hashes"
  type = "[]int64"
}
module talent_grid_group_opposing_node_hashes {
  source = "./attribute"

  name = "opposing Node Hashes"
  type = "[]int64"
}


module talent_grid_exclusive_set {
  source = "./struct"

  name = "exclusive Set"
  attributes_content = join("\n", [
  ])
}



module talent_grid_node_category {
  source = "./struct"

  name = "node Category"
  attributes_content = join("\n", [
    module.talent_grid_node_category_identifier.rendered,
    module.talent_grid_node_category_is_lore_driven.rendered,
    module.talent_grid_node_category_display_properties.rendered,
    module.talent_grid_node_category_node_hashes.rendered,
  ])
}

module talent_grid_node_category_identifier {
  source = "./attribute"

  name = "identifier"
  type = "string"
}
module talent_grid_node_category_is_lore_driven {
  source = "./attribute"

  name = "is Lore Driven"
  type = "bool"
}
module talent_grid_node_category_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module talent_grid_node_category_node_hashes {
  source = "./attribute"

  name = "node Hashes"
  type = "[]int64"
}


module talent_grid_activation_requirement {
  source = "./struct"

  name = "activation Requirement"
  attributes_content = join("\n", [
    module.talent_grid_activation_requirement_exclusive_set_required_hash.rendered,
    module.talent_grid_activation_requirement_grid_level.rendered,
    module.talent_grid_activation_requirement_material_requirement_hashes.rendered,
  ])
}

module talent_grid_activation_requirement_exclusive_set_required_hash {
  source = "./attribute"

  name = "exclusive Set Required Hash"
  type = "int"
}
module talent_grid_activation_requirement_grid_level {
  source = "./attribute"

  name = "grid Level"
  type = "int"
}
module talent_grid_activation_requirement_material_requirement_hashes {
  source = "./attribute"

  name = "material Requirement Hashes"
  type = "[]int64"
}


module talent_grid_node {
  source = "./struct"

  name = "node"
  attributes_content = join("\n", [
    module.talent_grid_node_auto_unlocks.rendered,
    module.talent_grid_node_binary_pair_node_index.rendered,
    module.talent_grid_node_column.rendered,
    module.talent_grid_node_exclusive_set_hash.rendered,
    module.talent_grid_node_exclusive_with_node_hashes.rendered,
    module.talent_grid_node_group_scope_index.rendered,
    module.talent_grid_node_ignore_for_completion.rendered,
    module.talent_grid_node_is_random.rendered,
    module.talent_grid_node_is_random_repurchasable.rendered,
    module.talent_grid_node_is_real_step_selection_random.rendered,
    module.talent_grid_node_last_step_repeats.rendered,
    module.talent_grid_node_layout_identifier.rendered,
    module.talent_grid_node_node_hash.rendered,
    module.talent_grid_node_node_index.rendered,
    module.talent_grid_node_node_style_identifier.rendered,
    module.talent_grid_node_original_node_hash.rendered,
    module.talent_grid_node_prerequisite_node_indexes.rendered,
    module.talent_grid_node_random_start_progression_bar_at_progression.rendered,
    module.talent_grid_node_row.rendered,
    module.talent_grid_node_steps.rendered,
    module.talent_grid_node_talent_node_types.rendered,
    module.talent_grid_node_group_hash.rendered,
  ])
}

module talent_grid_node_auto_unlocks {
  source = "./attribute"

  name = "auto Unlocks"
  type = "bool"
}
module talent_grid_node_binary_pair_node_index {
  source = "./attribute"

  name = "binary Pair Node Index"
  type = "int"
}
module talent_grid_node_column {
  source = "./attribute"

  name = "column"
  type = "int"
}
module talent_grid_node_exclusive_set_hash {
  source = "./attribute"

  name = "exclusive Set Hash"
  type = "int"
}
module talent_grid_node_exclusive_with_node_hashes {
  source = "./attribute"

  name = "exclusive With Node Hashes"
  type = "[]int64"
}
module talent_grid_node_group_scope_index {
  source = "./attribute"

  name = "group Scope Index"
  type = "int"
}
module talent_grid_node_ignore_for_completion {
  source = "./attribute"

  name = "ignore For Completion"
  type = "bool"
}
module talent_grid_node_is_random {
  source = "./attribute"

  name = "is Random"
  type = "bool"
}
module talent_grid_node_is_random_repurchasable {
  source = "./attribute"

  name = "is Random Repurchasable"
  type = "bool"
}
module talent_grid_node_is_real_step_selection_random {
  source = "./attribute"

  name = "is Real Step Selection Random"
  type = "bool"
}
module talent_grid_node_last_step_repeats {
  source = "./attribute"

  name = "last Step Repeats"
  type = "bool"
}
module talent_grid_node_layout_identifier {
  source = "./attribute"

  name = "layout Identifier"
  type = "string"
}
module talent_grid_node_node_hash {
  source = "./attribute"

  name = "node Hash"
  type = "int"
}
module talent_grid_node_node_index {
  source = "./attribute"

  name = "node Index"
  type = "int"
}
module talent_grid_node_node_style_identifier {
  source = "./attribute"

  name = "node Style Identifier"
  type = "string"
}
module talent_grid_node_original_node_hash {
  source = "./attribute"

  name = "original Node Hash"
  type = "int"
}
module talent_grid_node_prerequisite_node_indexes {
  source = "./attribute"

  name = "prerequisite Node Indexes"
  type = "[]int32"
}
module talent_grid_node_random_start_progression_bar_at_progression {
  source = "./attribute"

  name = "random Start Progression Bar At Progression"
  type = "int"
}
module talent_grid_node_row {
  source = "./attribute"

  name = "row"
  type = "int"
}
module talent_grid_node_steps {
  source = "./attribute"

  name = "steps"
  type = "[]Step"
}
module talent_grid_node_talent_node_types {
  source = "./attribute"

  name = "talent Node Types"
  type = "int"
}
module talent_grid_node_group_hash {
  source = "./attribute"

  name = "group Hash"
  type = "int64"
}


module talent_grid_step {
  source = "./struct"

  name = "step"
  attributes_content = join("\n", [
    module.talent_grid_step_activation_requirement.rendered,
    module.talent_grid_step_affects_level.rendered,
    module.talent_grid_step_affects_quality.rendered,
    module.talent_grid_step_can_activate_next_step.rendered,
    module.talent_grid_step_damage_type.rendered,
    module.talent_grid_step_display_properties.rendered,
    module.talent_grid_step_interaction_description.rendered,
    module.talent_grid_step_is_next_step_random.rendered,
    module.talent_grid_step_next_step_index.rendered,
    module.talent_grid_step_node_step_hash.rendered,
    module.talent_grid_step_perk_hashes.rendered,
    module.talent_grid_step_start_progression_bar_at_progress.rendered,
    module.talent_grid_step_stat_hashes.rendered,
    module.talent_grid_step_step_index.rendered,
    module.talent_grid_step_true_property_index.rendered,
    module.talent_grid_step_true_step_index.rendered,
  ])
}

module talent_grid_step_activation_requirement {
  source = "./attribute"

  name = "activation Requirement"
  type = "ActivationRequirement"
}
module talent_grid_step_affects_level {
  source = "./attribute"

  name = "affects Level"
  type = "bool"
}
module talent_grid_step_affects_quality {
  source = "./attribute"

  name = "affects Quality"
  type = "bool"
}
module talent_grid_step_can_activate_next_step {
  source = "./attribute"

  name = "can Activate Next Step"
  type = "bool"
}
module talent_grid_step_damage_type {
  source = "./attribute"

  name = "damage Type"
  type = "int"
}
module talent_grid_step_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module talent_grid_step_interaction_description {
  source = "./attribute"

  name = "interaction Description"
  type = "string"
}
module talent_grid_step_is_next_step_random {
  source = "./attribute"

  name = "is Next Step Random"
  type = "bool"
}
module talent_grid_step_next_step_index {
  source = "./attribute"

  name = "next Step Index"
  type = "int"
}
module talent_grid_step_node_step_hash {
  source = "./attribute"

  name = "node Step Hash"
  type = "int"
}
module talent_grid_step_perk_hashes {
  source = "./attribute"

  name = "perk Hashes"
  type = "[]int64"
}
module talent_grid_step_start_progression_bar_at_progress {
  source = "./attribute"

  name = "start Progression Bar At Progress"
  type = "int"
}
module talent_grid_step_stat_hashes {
  source = "./attribute"

  name = "stat Hashes"
  type = "[]int64"
}
module talent_grid_step_step_index {
  source = "./attribute"

  name = "step Index"
  type = "int"
}
module talent_grid_step_true_property_index {
  source = "./attribute"

  name = "true Property Index"
  type = "int"
}
module talent_grid_step_true_step_index {
  source = "./attribute"

  name = "true Step Index"
  type = "int"
}


