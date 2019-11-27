// Code generated; DANGER ZONE FOR EDITS

module inventory_item_definition {
  source = "./definition"

  target = var.target

  name = "inventory Item"
  structs_content = join("\n", [
    module.inventory_item_inventory_item_definition.rendered,
    module.inventory_item_inventory_item_action.rendered,
    module.inventory_item_background_color.rendered,
    module.inventory_item_equipping_block.rendered,
    module.inventory_item_inventory.rendered,
    module.inventory_item_investment_stats.rendered,
    module.inventory_item_preview.rendered,
    module.inventory_item_quality.rendered,
    module.inventory_item_intrinsic_sockets.rendered,
    module.inventory_item_socket_categories.rendered,
    module.inventory_item_socket_entries.rendered,
    module.inventory_item_sockets.rendered,
    module.inventory_item_source_data.rendered,
    module.inventory_item_stats.rendered,
    module.inventory_item_talent_grid.rendered,
  ])

}

module inventory_item_inventory_item_definition {
  source = "./struct"

  name = "inventory Item Definition"
  attributes_content = join("\n", [
    module.inventory_item_inventory_item_definition_acquire_reward_site_hash.rendered,
    module.inventory_item_inventory_item_definition_acquire_unlock_hash.rendered,
    module.inventory_item_inventory_item_definition_action.rendered,
    module.inventory_item_inventory_item_definition_allow_actions.rendered,
    module.inventory_item_inventory_item_definition_background_color.rendered,
    module.inventory_item_inventory_item_definition_blacklisted.rendered,
    module.inventory_item_inventory_item_definition_breaker_type.rendered,
    module.inventory_item_inventory_item_definition_class_type.rendered,
    module.inventory_item_inventory_item_definition_default_damage_type.rendered,
    module.inventory_item_inventory_item_definition_display_properties.rendered,
    module.inventory_item_inventory_item_definition_display_source.rendered,
    module.inventory_item_inventory_item_definition_does_postmaster_pull_have_side_effects.rendered,
    module.inventory_item_inventory_item_definition_equippable.rendered,
    module.inventory_item_inventory_item_definition_equipping_block.rendered,
    module.inventory_item_inventory_item_definition_hash.rendered,
    module.inventory_item_inventory_item_definition_index.rendered,
    module.inventory_item_inventory_item_definition_inventory.rendered,
    module.inventory_item_inventory_item_definition_investment_stats.rendered,
    module.inventory_item_inventory_item_definition_is_wrapper.rendered,
    module.inventory_item_inventory_item_definition_item_category_hashes.rendered,
    module.inventory_item_inventory_item_definition_item_sub_type.rendered,
    module.inventory_item_inventory_item_definition_item_type.rendered,
    module.inventory_item_inventory_item_definition_item_type_and_tier_display_name.rendered,
    module.inventory_item_inventory_item_definition_item_type_display_name.rendered,
    module.inventory_item_inventory_item_definition_non_transferrable.rendered,
    module.inventory_item_inventory_item_definition_perks.rendered,
    module.inventory_item_inventory_item_definition_preview.rendered,
    module.inventory_item_inventory_item_definition_quality.rendered,
    module.inventory_item_inventory_item_definition_redacted.rendered,
    module.inventory_item_inventory_item_definition_secondary_icon.rendered,
    module.inventory_item_inventory_item_definition_secondary_overlay.rendered,
    module.inventory_item_inventory_item_definition_secondary_special.rendered,
    module.inventory_item_inventory_item_definition_screenshot.rendered,
    module.inventory_item_inventory_item_definition_sockets.rendered,
    module.inventory_item_inventory_item_definition_source_data.rendered,
    module.inventory_item_inventory_item_definition_special_item_type.rendered,
    module.inventory_item_inventory_item_definition_stats.rendered,
    module.inventory_item_inventory_item_definition_summary_item_hash.rendered,
    module.inventory_item_inventory_item_definition_talent_grid.rendered,
    module.inventory_item_inventory_item_definition_tooltip_notifications.rendered,
    module.inventory_item_inventory_item_definition_translation_block.rendered,
    module.inventory_item_inventory_item_definition_ui_item_display_style.rendered,
  ])
}

module inventory_item_inventory_item_definition_acquire_reward_site_hash {
  source = "./attribute"

  name = "acquire Reward Site Hash"
  type = "int"
}
module inventory_item_inventory_item_definition_acquire_unlock_hash {
  source = "./attribute"

  name = "acquire Unlock Hash"
  type = "int"
}
module inventory_item_inventory_item_definition_action {
  source = "./attribute"

  name = "action"
  type = "InventoryItemAction"
}
module inventory_item_inventory_item_definition_allow_actions {
  source = "./attribute"

  name = "allow Actions"
  type = "bool"
}
module inventory_item_inventory_item_definition_background_color {
  source = "./attribute"

  name = "background Color"
  type = "BackgroundColor"
}
module inventory_item_inventory_item_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module inventory_item_inventory_item_definition_breaker_type {
  source = "./attribute"

  name = "breaker Type"
  type = "int"
}
module inventory_item_inventory_item_definition_class_type {
  source = "./attribute"

  name = "class Type"
  type = "int"
}
module inventory_item_inventory_item_definition_default_damage_type {
  source = "./attribute"

  name = "default Damage Type"
  type = "int"
}
module inventory_item_inventory_item_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module inventory_item_inventory_item_definition_display_source {
  source = "./attribute"

  name = "display Source"
  type = "string"
}
module inventory_item_inventory_item_definition_does_postmaster_pull_have_side_effects {
  source = "./attribute"

  name = "does Postmaster Pull Have Side Effects"
  type = "bool"
}
module inventory_item_inventory_item_definition_equippable {
  source = "./attribute"

  name = "equippable"
  type = "bool"
}
module inventory_item_inventory_item_definition_equipping_block {
  source = "./attribute"

  name = "equipping Block"
  type = "EquippingBlock"
}
module inventory_item_inventory_item_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module inventory_item_inventory_item_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module inventory_item_inventory_item_definition_inventory {
  source = "./attribute"

  name = "inventory"
  type = "Inventory"
}
module inventory_item_inventory_item_definition_investment_stats {
  source = "./attribute"

  name = "investment Stats"
  type = "[]InvestmentStats"
}
module inventory_item_inventory_item_definition_is_wrapper {
  source = "./attribute"

  name = "is Wrapper"
  type = "bool"
}
module inventory_item_inventory_item_definition_item_category_hashes {
  source = "./attribute"

  name = "item Category Hashes"
  type = "[]int"
}
module inventory_item_inventory_item_definition_item_sub_type {
  source = "./attribute"

  name = "item Sub Type"
  type = "int"
}
module inventory_item_inventory_item_definition_item_type {
  source = "./attribute"

  name = "item Type"
  type = "int"
}
module inventory_item_inventory_item_definition_item_type_and_tier_display_name {
  source = "./attribute"

  name = "item Type And Tier Display Name"
  type = "string"
}
module inventory_item_inventory_item_definition_item_type_display_name {
  source = "./attribute"

  name = "item Type Display Name"
  type = "string"
}
module inventory_item_inventory_item_definition_non_transferrable {
  source = "./attribute"

  name = "non Transferrable"
  type = "bool"
}
module inventory_item_inventory_item_definition_perks {
  source = "./attribute"

  name = "perks"
  type = "[]interface{}"
}
module inventory_item_inventory_item_definition_preview {
  source = "./attribute"

  name = "preview"
  type = "Preview"
}
module inventory_item_inventory_item_definition_quality {
  source = "./attribute"

  name = "quality"
  type = "Quality"
}
module inventory_item_inventory_item_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module inventory_item_inventory_item_definition_secondary_icon {
  source = "./attribute"

  name = "secondary Icon"
  type = "string"
}
module inventory_item_inventory_item_definition_secondary_overlay {
  source = "./attribute"

  name = "secondary Overlay"
  type = "string"
}
module inventory_item_inventory_item_definition_secondary_special {
  source = "./attribute"

  name = "secondary Special"
  type = "string"
}
module inventory_item_inventory_item_definition_screenshot {
  source = "./attribute"

  name = "screenshot"
  type = "string"
}
module inventory_item_inventory_item_definition_sockets {
  source = "./attribute"

  name = "sockets"
  type = "Sockets"
}
module inventory_item_inventory_item_definition_source_data {
  source = "./attribute"

  name = "source Data"
  type = "SourceData"
}
module inventory_item_inventory_item_definition_special_item_type {
  source = "./attribute"

  name = "special Item Type"
  type = "int"
}
module inventory_item_inventory_item_definition_stats {
  source = "./attribute"

  name = "stats"
  type = "Stats"
}
module inventory_item_inventory_item_definition_summary_item_hash {
  source = "./attribute"

  name = "summary Item Hash"
  type = "int64"
}
module inventory_item_inventory_item_definition_talent_grid {
  source = "./attribute"

  name = "talent Grid"
  type = "TalentGrid"
}
module inventory_item_inventory_item_definition_tooltip_notifications {
  source = "./attribute"

  name = "tooltip Notifications"
  type = "[]interface{}"
}
module inventory_item_inventory_item_definition_translation_block {
  source = "./attribute"

  name = "translation Block"
  type = "TranslationBlock"
}
module inventory_item_inventory_item_definition_ui_item_display_style {
  source = "./attribute"

  name = "ui Item Display Style"
  type = "string"
}


module inventory_item_inventory_item_action {
  source = "./struct"

  name = "inventory Item Action"
  attributes_content = join("\n", [
    module.inventory_item_inventory_item_action_action_type_label.rendered,
    module.inventory_item_inventory_item_action_consume_entire_stack.rendered,
    module.inventory_item_inventory_item_action_delete_on_action.rendered,
    module.inventory_item_inventory_item_action_is_positive.rendered,
    module.inventory_item_inventory_item_action_progression_rewards.rendered,
    module.inventory_item_inventory_item_action_required_cooldown_hash.rendered,
    module.inventory_item_inventory_item_action_required_cooldown_seconds.rendered,
    module.inventory_item_inventory_item_action_required_items.rendered,
    module.inventory_item_inventory_item_action_reward_item_hash.rendered,
    module.inventory_item_inventory_item_action_reward_sheet_hash.rendered,
    module.inventory_item_inventory_item_action_reward_site_hash.rendered,
    module.inventory_item_inventory_item_action_use_on_acquire.rendered,
    module.inventory_item_inventory_item_action_verb_description.rendered,
    module.inventory_item_inventory_item_action_verb_name.rendered,
  ])
}

module inventory_item_inventory_item_action_action_type_label {
  source = "./attribute"

  name = "action Type Label"
  type = "string"
}
module inventory_item_inventory_item_action_consume_entire_stack {
  source = "./attribute"

  name = "consume Entire Stack"
  type = "bool"
}
module inventory_item_inventory_item_action_delete_on_action {
  source = "./attribute"

  name = "delete On Action"
  type = "bool"
}
module inventory_item_inventory_item_action_is_positive {
  source = "./attribute"

  name = "is Positive"
  type = "bool"
}
module inventory_item_inventory_item_action_progression_rewards {
  source = "./attribute"

  name = "progression Rewards"
  type = "[]interface{}"
}
module inventory_item_inventory_item_action_required_cooldown_hash {
  source = "./attribute"

  name = "required Cooldown Hash"
  type = "int"
}
module inventory_item_inventory_item_action_required_cooldown_seconds {
  source = "./attribute"

  name = "required Cooldown Seconds"
  type = "int"
}
module inventory_item_inventory_item_action_required_items {
  source = "./attribute"

  name = "required Items"
  type = "[]interface{}"
}
module inventory_item_inventory_item_action_reward_item_hash {
  source = "./attribute"

  name = "reward Item Hash"
  type = "int"
}
module inventory_item_inventory_item_action_reward_sheet_hash {
  source = "./attribute"

  name = "reward Sheet Hash"
  type = "int"
}
module inventory_item_inventory_item_action_reward_site_hash {
  source = "./attribute"

  name = "reward Site Hash"
  type = "int"
}
module inventory_item_inventory_item_action_use_on_acquire {
  source = "./attribute"

  name = "use On Acquire"
  type = "bool"
}
module inventory_item_inventory_item_action_verb_description {
  source = "./attribute"

  name = "verb Description"
  type = "string"
}
module inventory_item_inventory_item_action_verb_name {
  source = "./attribute"

  name = "verb Name"
  type = "string"
}


module inventory_item_background_color {
  source = "./struct"

  name = "background Color"
  attributes_content = join("\n", [
    module.inventory_item_background_color_alpha.rendered,
    module.inventory_item_background_color_blue.rendered,
    module.inventory_item_background_color_colorhash.rendered,
    module.inventory_item_background_color_green.rendered,
    module.inventory_item_background_color_red.rendered,
  ])
}

module inventory_item_background_color_alpha {
  source = "./attribute"

  name = "alpha"
  type = "int"
}
module inventory_item_background_color_blue {
  source = "./attribute"

  name = "blue"
  type = "int"
}
module inventory_item_background_color_colorhash {
  source = "./attribute"

  name = "colorHash"
  type = "int"
}
module inventory_item_background_color_green {
  source = "./attribute"

  name = "green"
  type = "int"
}
module inventory_item_background_color_red {
  source = "./attribute"

  name = "red"
  type = "int"
}


module inventory_item_equipping_block {
  source = "./struct"

  name = "equipping Block"
  attributes_content = join("\n", [
    module.inventory_item_equipping_block_ammo_type.rendered,
    module.inventory_item_equipping_block_attributes.rendered,
    module.inventory_item_equipping_block_display_strings.rendered,
    module.inventory_item_equipping_block_equipment_slot_type_hash.rendered,
    module.inventory_item_equipping_block_equipping_sound_hash.rendered,
    module.inventory_item_equipping_block_horn_sound_hash.rendered,
    module.inventory_item_equipping_block_unique_label_hash.rendered,
  ])
}

module inventory_item_equipping_block_ammo_type {
  source = "./attribute"

  name = "ammo Type"
  type = "int"
}
module inventory_item_equipping_block_attributes {
  source = "./attribute"

  name = "attributes"
  type = "int"
}
module inventory_item_equipping_block_display_strings {
  source = "./attribute"

  name = "display Strings"
  type = "[]string"
}
module inventory_item_equipping_block_equipment_slot_type_hash {
  source = "./attribute"

  name = "equipment Slot Type Hash"
  type = "int"
}
module inventory_item_equipping_block_equipping_sound_hash {
  source = "./attribute"

  name = "equipping Sound Hash"
  type = "int"
}
module inventory_item_equipping_block_horn_sound_hash {
  source = "./attribute"

  name = "horn Sound Hash"
  type = "int"
}
module inventory_item_equipping_block_unique_label_hash {
  source = "./attribute"

  name = "unique Label Hash"
  type = "int"
}


module inventory_item_inventory {
  source = "./struct"

  name = "inventory"
  attributes_content = join("\n", [
    module.inventory_item_inventory_bucket_type_hash.rendered,
    module.inventory_item_inventory_expiration_tooltip.rendered,
    module.inventory_item_inventory_expired_in_activity_message.rendered,
    module.inventory_item_inventory_expired_in_orbit_message.rendered,
    module.inventory_item_inventory_is_instance_item.rendered,
    module.inventory_item_inventory_max_stack_size.rendered,
    module.inventory_item_inventory_non_transferrable_original.rendered,
    module.inventory_item_inventory_recovery_bucket_type_hash.rendered,
    module.inventory_item_inventory_suppress_expiration_when_objectives_complete.rendered,
    module.inventory_item_inventory_tier_type.rendered,
    module.inventory_item_inventory_tier_type_hash.rendered,
    module.inventory_item_inventory_tier_type_name.rendered,
  ])
}

module inventory_item_inventory_bucket_type_hash {
  source = "./attribute"

  name = "bucket Type Hash"
  type = "int"
}
module inventory_item_inventory_expiration_tooltip {
  source = "./attribute"

  name = "expiration Tooltip"
  type = "string"
}
module inventory_item_inventory_expired_in_activity_message {
  source = "./attribute"

  name = "expired In Activity Message"
  type = "string"
}
module inventory_item_inventory_expired_in_orbit_message {
  source = "./attribute"

  name = "expired In Orbit Message"
  type = "string"
}
module inventory_item_inventory_is_instance_item {
  source = "./attribute"

  name = "is Instance Item"
  type = "bool"
}
module inventory_item_inventory_max_stack_size {
  source = "./attribute"

  name = "max Stack Size"
  type = "int"
}
module inventory_item_inventory_non_transferrable_original {
  source = "./attribute"

  name = "non Transferrable Original"
  type = "bool"
}
module inventory_item_inventory_recovery_bucket_type_hash {
  source = "./attribute"

  name = "recovery Bucket Type Hash"
  type = "int"
}
module inventory_item_inventory_suppress_expiration_when_objectives_complete {
  source = "./attribute"

  name = "suppress Expiration When Objectives Complete"
  type = "bool"
}
module inventory_item_inventory_tier_type {
  source = "./attribute"

  name = "tier Type"
  type = "int"
}
module inventory_item_inventory_tier_type_hash {
  source = "./attribute"

  name = "tier Type Hash"
  type = "int64"
}
module inventory_item_inventory_tier_type_name {
  source = "./attribute"

  name = "tier Type Name"
  type = "string"
}


module inventory_item_investment_stats {
  source = "./struct"

  name = "investment Stats"
  attributes_content = join("\n", [
    module.inventory_item_investment_stats_is_conditionally_active.rendered,
    module.inventory_item_investment_stats_stat_type_hash.rendered,
    module.inventory_item_investment_stats_value.rendered,
  ])
}

module inventory_item_investment_stats_is_conditionally_active {
  source = "./attribute"

  name = "is Conditionally Active"
  type = "bool"
}
module inventory_item_investment_stats_stat_type_hash {
  source = "./attribute"

  name = "stat Type Hash"
  type = "int64"
}
module inventory_item_investment_stats_value {
  source = "./attribute"

  name = "value"
  type = "int"
}


module inventory_item_preview {
  source = "./struct"

  name = "preview"
  attributes_content = join("\n", [
    module.inventory_item_preview_preview_action_string.rendered,
    module.inventory_item_preview_preview_vendor_hash.rendered,
    module.inventory_item_preview_screen_style.rendered,
  ])
}

module inventory_item_preview_preview_action_string {
  source = "./attribute"

  name = "preview Action String"
  type = "string"
}
module inventory_item_preview_preview_vendor_hash {
  source = "./attribute"

  name = "preview Vendor Hash"
  type = "int"
}
module inventory_item_preview_screen_style {
  source = "./attribute"

  name = "screen Style"
  type = "string"
}


module inventory_item_quality {
  source = "./struct"

  name = "quality"
  attributes_content = join("\n", [
    module.inventory_item_quality_infusion_category_hash.rendered,
    module.inventory_item_quality_infusion_category_hashes.rendered,
    module.inventory_item_quality_infusion_category_name.rendered,
    module.inventory_item_quality_item_levels.rendered,
    module.inventory_item_quality_progression_level_requirement_hash.rendered,
    module.inventory_item_quality_quality_level.rendered,
  ])
}

module inventory_item_quality_infusion_category_hash {
  source = "./attribute"

  name = "infusion Category Hash"
  type = "int64"
}
module inventory_item_quality_infusion_category_hashes {
  source = "./attribute"

  name = "infusion Category Hashes"
  type = "[]int64"
}
module inventory_item_quality_infusion_category_name {
  source = "./attribute"

  name = "infusion Category Name"
  type = "string"
}
module inventory_item_quality_item_levels {
  source = "./attribute"

  name = "item Levels"
  type = "[]interface{}"
}
module inventory_item_quality_progression_level_requirement_hash {
  source = "./attribute"

  name = "progression Level Requirement Hash"
  type = "int64"
}
module inventory_item_quality_quality_level {
  source = "./attribute"

  name = "quality Level"
  type = "int"
}


module inventory_item_intrinsic_sockets {
  source = "./struct"

  name = "intrinsic Sockets"
  attributes_content = join("\n", [
    module.inventory_item_intrinsic_sockets_default_visible.rendered,
    module.inventory_item_intrinsic_sockets_plug_item_hash.rendered,
    module.inventory_item_intrinsic_sockets_socket_type_hash.rendered,
  ])
}

module inventory_item_intrinsic_sockets_default_visible {
  source = "./attribute"

  name = "default Visible"
  type = "bool"
}
module inventory_item_intrinsic_sockets_plug_item_hash {
  source = "./attribute"

  name = "plug Item Hash"
  type = "int64"
}
module inventory_item_intrinsic_sockets_socket_type_hash {
  source = "./attribute"

  name = "socket Type Hash"
  type = "int"
}


module inventory_item_socket_categories {
  source = "./struct"

  name = "socket Categories"
  attributes_content = join("\n", [
    module.inventory_item_socket_categories_socket_category_hash.rendered,
    module.inventory_item_socket_categories_socket_indexes.rendered,
  ])
}

module inventory_item_socket_categories_socket_category_hash {
  source = "./attribute"

  name = "socket Category Hash"
  type = "int"
}
module inventory_item_socket_categories_socket_indexes {
  source = "./attribute"

  name = "socket Indexes"
  type = "[]int"
}


module inventory_item_socket_entries {
  source = "./struct"

  name = "socket Entries"
  attributes_content = join("\n", [
    module.inventory_item_socket_entries_default_visible.rendered,
    module.inventory_item_socket_entries_hide_perks_in_item_tooltip.rendered,
    module.inventory_item_socket_entries_overrides_ui_appearance.rendered,
    module.inventory_item_socket_entries_plug_sources.rendered,
    module.inventory_item_socket_entries_prevent_initialization_on_vendor_purchase.rendered,
    module.inventory_item_socket_entries_prevent_initialization_when_versioning.rendered,
    module.inventory_item_socket_entries_reusable_plug_items.rendered,
    module.inventory_item_socket_entries_reusable_plug_set_hash.rendered,
    module.inventory_item_socket_entries_single_initial_item_hash.rendered,
    module.inventory_item_socket_entries_socket_type_hash.rendered,
  ])
}

module inventory_item_socket_entries_default_visible {
  source = "./attribute"

  name = "default Visible"
  type = "bool"
}
module inventory_item_socket_entries_hide_perks_in_item_tooltip {
  source = "./attribute"

  name = "hide Perks In Item Tooltip"
  type = "bool"
}
module inventory_item_socket_entries_overrides_ui_appearance {
  source = "./attribute"

  name = "overrides Ui Appearance"
  type = "bool"
}
module inventory_item_socket_entries_plug_sources {
  source = "./attribute"

  name = "plug Sources"
  type = "int"
}
module inventory_item_socket_entries_prevent_initialization_on_vendor_purchase {
  source = "./attribute"

  name = "prevent Initialization On Vendor Purchase"
  type = "bool"
}
module inventory_item_socket_entries_prevent_initialization_when_versioning {
  source = "./attribute"

  name = "prevent Initialization When Versioning"
  type = "bool"
}
module inventory_item_socket_entries_reusable_plug_items {
  source = "./attribute"

  name = "reusable Plug Items"
  type = "[]interface{}"
}
module inventory_item_socket_entries_reusable_plug_set_hash {
  source = "./attribute"

  name = "reusable Plug Set Hash"
  type = "int64"
}
module inventory_item_socket_entries_single_initial_item_hash {
  source = "./attribute"

  name = "single Initial Item Hash"
  type = "int"
}
module inventory_item_socket_entries_socket_type_hash {
  source = "./attribute"

  name = "socket Type Hash"
  type = "int"
}


module inventory_item_sockets {
  source = "./struct"

  name = "sockets"
  attributes_content = join("\n", [
    module.inventory_item_sockets_detail.rendered,
    module.inventory_item_sockets_intrinsic_sockets.rendered,
    module.inventory_item_sockets_socket_categories.rendered,
    module.inventory_item_sockets_socket_entries.rendered,
  ])
}

module inventory_item_sockets_detail {
  source = "./attribute"

  name = "detail"
  type = "string"
}
module inventory_item_sockets_intrinsic_sockets {
  source = "./attribute"

  name = "intrinsic Sockets"
  type = "[]IntrinsicSockets"
}
module inventory_item_sockets_socket_categories {
  source = "./attribute"

  name = "socket Categories"
  type = "[]SocketCategories"
}
module inventory_item_sockets_socket_entries {
  source = "./attribute"

  name = "socket Entries"
  type = "[]SocketEntries"
}


module inventory_item_source_data {
  source = "./struct"

  name = "source Data"
  attributes_content = join("\n", [
    module.inventory_item_source_data_exclusive.rendered,
    module.inventory_item_source_data_vendor_sources.rendered,
  ])
}

module inventory_item_source_data_exclusive {
  source = "./attribute"

  name = "exclusive"
  type = "int"
}
module inventory_item_source_data_vendor_sources {
  source = "./attribute"

  name = "vendor Sources"
  type = "[]interface{}"
}


module inventory_item_stats {
  source = "./struct"

  name = "stats"
  attributes_content = join("\n", [
    module.inventory_item_stats_disable_primary_stat_display.rendered,
    module.inventory_item_stats_has_displayable_stats.rendered,
    module.inventory_item_stats_primary_base_stat_hash.rendered,
    module.inventory_item_stats_stat_group_hash.rendered,
    module.inventory_item_stats_stats.rendered,
  ])
}

module inventory_item_stats_disable_primary_stat_display {
  source = "./attribute"

  name = "disable Primary Stat Display"
  type = "bool"
}
module inventory_item_stats_has_displayable_stats {
  source = "./attribute"

  name = "has Displayable Stats"
  type = "bool"
}
module inventory_item_stats_primary_base_stat_hash {
  source = "./attribute"

  name = "primary Base Stat Hash"
  type = "int64"
}
module inventory_item_stats_stat_group_hash {
  source = "./attribute"

  name = "stat Group Hash"
  type = "int64"
}
module inventory_item_stats_stats {
  source = "./attribute"

  name = "stats"
  type = "map[string]interface{}"
}


module inventory_item_talent_grid {
  source = "./struct"

  name = "talent Grid"
  attributes_content = join("\n", [
    module.inventory_item_talent_grid_hud_damage_type.rendered,
    module.inventory_item_talent_grid_item_detail_string.rendered,
    module.inventory_item_talent_grid_talent_grid_hash.rendered,
  ])
}

module inventory_item_talent_grid_hud_damage_type {
  source = "./attribute"

  name = "hud Damage Type"
  type = "int"
}
module inventory_item_talent_grid_item_detail_string {
  source = "./attribute"

  name = "item Detail String"
  type = "string"
}
module inventory_item_talent_grid_talent_grid_hash {
  source = "./attribute"

  name = "talent Grid Hash"
  type = "int"
}


