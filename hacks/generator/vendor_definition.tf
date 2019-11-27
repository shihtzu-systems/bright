// Code generated; DANGER ZONE FOR EDITS

module vendor_definition {
  source = "./definition"

  target = var.target

  name = "vendor"
  structs_content = join("\n", [
    module.vendor_vendor_definition.rendered,
    module.vendor_categories.rendered,
    module.vendor_display_categories.rendered,
    module.vendor_vendor_action.rendered,
    module.vendor_creation_levels.rendered,
    module.vendor_item_list.rendered,
    module.vendor_original_categories.rendered,
  ])

}

module vendor_vendor_definition {
  source = "./struct"

  name = "vendor Definition"
  attributes_content = join("\n", [
    module.vendor_vendor_definition_accepted_items.rendered,
    module.vendor_vendor_definition_actions.rendered,
    module.vendor_vendor_definition_blacklisted.rendered,
    module.vendor_vendor_definition_categories.rendered,
    module.vendor_vendor_definition_consolidate_categories.rendered,
    module.vendor_vendor_definition_display_categories.rendered,
    module.vendor_vendor_definition_display_item_hash.rendered,
    module.vendor_vendor_definition_display_properties.rendered,
    module.vendor_vendor_definition_enabled.rendered,
    module.vendor_vendor_definition_faction_hash.rendered,
    module.vendor_vendor_definition_failure_strings.rendered,
    module.vendor_vendor_definition_hash.rendered,
    module.vendor_vendor_definition_ignore_sale_item_hashes.rendered,
    module.vendor_vendor_definition_index.rendered,
    module.vendor_vendor_definition_inhibit_buying.rendered,
    module.vendor_vendor_definition_inhibit_selling.rendered,
    module.vendor_vendor_definition_interactions.rendered,
    module.vendor_vendor_definition_inventory_flyouts.rendered,
    module.vendor_vendor_definition_item_list.rendered,
    module.vendor_vendor_definition_original_categories.rendered,
    module.vendor_vendor_definition_redacted.rendered,
    module.vendor_vendor_definition_reset_interval_minutes.rendered,
    module.vendor_vendor_definition_reset_offset_minutes.rendered,
    module.vendor_vendor_definition_return_with_vendor_request.rendered,
    module.vendor_vendor_definition_services.rendered,
    module.vendor_vendor_definition_unlock_ranges.rendered,
    module.vendor_vendor_definition_unlock_value_hash.rendered,
    module.vendor_vendor_definition_visible.rendered,
  ])
}

module vendor_vendor_definition_accepted_items {
  source = "./attribute"

  name = "accepted Items"
  type = "[]interface{}"
}
module vendor_vendor_definition_actions {
  source = "./attribute"

  name = "actions"
  type = "[]interface{}"
}
module vendor_vendor_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module vendor_vendor_definition_categories {
  source = "./attribute"

  name = "categories"
  type = "[]Categories"
}
module vendor_vendor_definition_consolidate_categories {
  source = "./attribute"

  name = "consolidate Categories"
  type = "bool"
}
module vendor_vendor_definition_display_categories {
  source = "./attribute"

  name = "display Categories"
  type = "[]DisplayCategories"
}
module vendor_vendor_definition_display_item_hash {
  source = "./attribute"

  name = "display Item Hash"
  type = "int"
}
module vendor_vendor_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module vendor_vendor_definition_enabled {
  source = "./attribute"

  name = "enabled"
  type = "bool"
}
module vendor_vendor_definition_faction_hash {
  source = "./attribute"

  name = "faction Hash"
  type = "int"
}
module vendor_vendor_definition_failure_strings {
  source = "./attribute"

  name = "failure Strings"
  type = "[]interface{}"
}
module vendor_vendor_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module vendor_vendor_definition_ignore_sale_item_hashes {
  source = "./attribute"

  name = "ignore Sale Item Hashes"
  type = "[]interface{}"
}
module vendor_vendor_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module vendor_vendor_definition_inhibit_buying {
  source = "./attribute"

  name = "inhibit Buying"
  type = "bool"
}
module vendor_vendor_definition_inhibit_selling {
  source = "./attribute"

  name = "inhibit Selling"
  type = "bool"
}
module vendor_vendor_definition_interactions {
  source = "./attribute"

  name = "interactions"
  type = "[]interface{}"
}
module vendor_vendor_definition_inventory_flyouts {
  source = "./attribute"

  name = "inventory Flyouts"
  type = "[]interface{}"
}
module vendor_vendor_definition_item_list {
  source = "./attribute"

  name = "item List"
  type = "[]ItemList"
}
module vendor_vendor_definition_original_categories {
  source = "./attribute"

  name = "original Categories"
  type = "[]OriginalCategories"
}
module vendor_vendor_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module vendor_vendor_definition_reset_interval_minutes {
  source = "./attribute"

  name = "reset Interval Minutes"
  type = "int"
}
module vendor_vendor_definition_reset_offset_minutes {
  source = "./attribute"

  name = "reset Offset Minutes"
  type = "int"
}
module vendor_vendor_definition_return_with_vendor_request {
  source = "./attribute"

  name = "return With Vendor Request"
  type = "bool"
}
module vendor_vendor_definition_services {
  source = "./attribute"

  name = "services"
  type = "[]interface{}"
}
module vendor_vendor_definition_unlock_ranges {
  source = "./attribute"

  name = "unlock Ranges"
  type = "[]interface{}"
}
module vendor_vendor_definition_unlock_value_hash {
  source = "./attribute"

  name = "unlock Value Hash"
  type = "int"
}
module vendor_vendor_definition_visible {
  source = "./attribute"

  name = "visible"
  type = "bool"
}


module vendor_categories {
  source = "./struct"

  name = "categories"
  attributes_content = join("\n", [
    module.vendor_categories_buy_string_override.rendered,
    module.vendor_categories_category_hash.rendered,
    module.vendor_categories_category_index.rendered,
    module.vendor_categories_disabled_description.rendered,
    module.vendor_categories_hide_from_regular_purchase.rendered,
    module.vendor_categories_hide_if_no_currency.rendered,
    module.vendor_categories_is_display_only.rendered,
    module.vendor_categories_is_preview.rendered,
    module.vendor_categories_quantity_available.rendered,
    module.vendor_categories_reset_interval_minutes_override.rendered,
    module.vendor_categories_reset_offset_minutes_override.rendered,
    module.vendor_categories_show_unavailable_items.rendered,
    module.vendor_categories_sort_value.rendered,
    module.vendor_categories_vendor_item_indexes.rendered,
  ])
}

module vendor_categories_buy_string_override {
  source = "./attribute"

  name = "buy String Override"
  type = "string"
}
module vendor_categories_category_hash {
  source = "./attribute"

  name = "category Hash"
  type = "int"
}
module vendor_categories_category_index {
  source = "./attribute"

  name = "category Index"
  type = "int"
}
module vendor_categories_disabled_description {
  source = "./attribute"

  name = "disabled Description"
  type = "string"
}
module vendor_categories_hide_from_regular_purchase {
  source = "./attribute"

  name = "hide From Regular Purchase"
  type = "bool"
}
module vendor_categories_hide_if_no_currency {
  source = "./attribute"

  name = "hide If No Currency"
  type = "bool"
}
module vendor_categories_is_display_only {
  source = "./attribute"

  name = "is Display Only"
  type = "bool"
}
module vendor_categories_is_preview {
  source = "./attribute"

  name = "is Preview"
  type = "bool"
}
module vendor_categories_quantity_available {
  source = "./attribute"

  name = "quantity Available"
  type = "int"
}
module vendor_categories_reset_interval_minutes_override {
  source = "./attribute"

  name = "reset Interval Minutes Override"
  type = "int"
}
module vendor_categories_reset_offset_minutes_override {
  source = "./attribute"

  name = "reset Offset Minutes Override"
  type = "int"
}
module vendor_categories_show_unavailable_items {
  source = "./attribute"

  name = "show Unavailable Items"
  type = "bool"
}
module vendor_categories_sort_value {
  source = "./attribute"

  name = "sort Value"
  type = "int"
}
module vendor_categories_vendor_item_indexes {
  source = "./attribute"

  name = "vendor Item Indexes"
  type = "[]int"
}


module vendor_display_categories {
  source = "./struct"

  name = "display Categories"
  attributes_content = join("\n", [
    module.vendor_display_categories_display_category_hash.rendered,
    module.vendor_display_categories_display_in_banner.rendered,
    module.vendor_display_categories_display_properties.rendered,
    module.vendor_display_categories_identifier.rendered,
    module.vendor_display_categories_index.rendered,
    module.vendor_display_categories_sort_order.rendered,
  ])
}

module vendor_display_categories_display_category_hash {
  source = "./attribute"

  name = "display Category Hash"
  type = "int"
}
module vendor_display_categories_display_in_banner {
  source = "./attribute"

  name = "display In Banner"
  type = "bool"
}
module vendor_display_categories_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module vendor_display_categories_identifier {
  source = "./attribute"

  name = "identifier"
  type = "string"
}
module vendor_display_categories_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module vendor_display_categories_sort_order {
  source = "./attribute"

  name = "sort Order"
  type = "int"
}


module vendor_vendor_action {
  source = "./struct"

  name = "vendor Action"
  attributes_content = join("\n", [
    module.vendor_vendor_action_execute_seconds.rendered,
    module.vendor_vendor_action_is_positive.rendered,
  ])
}

module vendor_vendor_action_execute_seconds {
  source = "./attribute"

  name = "execute Seconds"
  type = "float64"
}
module vendor_vendor_action_is_positive {
  source = "./attribute"

  name = "is Positive"
  type = "bool"
}


module vendor_creation_levels {
  source = "./struct"

  name = "creation Levels"
  attributes_content = join("\n", [
    module.vendor_creation_levels_level.rendered,
  ])
}

module vendor_creation_levels_level {
  source = "./attribute"

  name = "level"
  type = "int"
}


module vendor_item_list {
  source = "./struct"

  name = "item List"
  attributes_content = join("\n", [
    module.vendor_item_list_action.rendered,
    module.vendor_item_list_category_index.rendered,
    module.vendor_item_list_creation_levels.rendered,
    module.vendor_item_list_currencies.rendered,
    module.vendor_item_list_display_category.rendered,
    module.vendor_item_list_display_category_index.rendered,
    module.vendor_item_list_exclusivity.rendered,
    module.vendor_item_list_expiration_tooltip.rendered,
    module.vendor_item_list_failure_indexes.rendered,
    module.vendor_item_list_inventory_bucket_hash.rendered,
    module.vendor_item_list_item_hash.rendered,
    module.vendor_item_list_license_unlock_hash.rendered,
    module.vendor_item_list_maximum_level.rendered,
    module.vendor_item_list_minimum_level.rendered,
    module.vendor_item_list_original_category_index.rendered,
    module.vendor_item_list_price_override_enabled.rendered,
    module.vendor_item_list_purchasable_scope.rendered,
    module.vendor_item_list_quantity.rendered,
    module.vendor_item_list_redirect_to_sale_indexes.rendered,
    module.vendor_item_list_refund_policy.rendered,
    module.vendor_item_list_refund_time_limit.rendered,
    module.vendor_item_list_reward_adjustor_pointer_hash.rendered,
    module.vendor_item_list_seed_override.rendered,
    module.vendor_item_list_socket_overrides.rendered,
    module.vendor_item_list_sort_value.rendered,
    module.vendor_item_list_vendor_item_index.rendered,
    module.vendor_item_list_visibility_scope.rendered,
    module.vendor_item_list_weight.rendered,
  ])
}

module vendor_item_list_action {
  source = "./attribute"

  name = "action"
  type = "VendorAction"
}
module vendor_item_list_category_index {
  source = "./attribute"

  name = "category Index"
  type = "int"
}
module vendor_item_list_creation_levels {
  source = "./attribute"

  name = "creation Levels"
  type = "[]CreationLevels"
}
module vendor_item_list_currencies {
  source = "./attribute"

  name = "currencies"
  type = "[]interface{}"
}
module vendor_item_list_display_category {
  source = "./attribute"

  name = "display Category"
  type = "string"
}
module vendor_item_list_display_category_index {
  source = "./attribute"

  name = "display Category Index"
  type = "int"
}
module vendor_item_list_exclusivity {
  source = "./attribute"

  name = "exclusivity"
  type = "int"
}
module vendor_item_list_expiration_tooltip {
  source = "./attribute"

  name = "expiration Tooltip"
  type = "string"
}
module vendor_item_list_failure_indexes {
  source = "./attribute"

  name = "failure Indexes"
  type = "[]interface{}"
}
module vendor_item_list_inventory_bucket_hash {
  source = "./attribute"

  name = "inventory Bucket Hash"
  type = "int64"
}
module vendor_item_list_item_hash {
  source = "./attribute"

  name = "item Hash"
  type = "int"
}
module vendor_item_list_license_unlock_hash {
  source = "./attribute"

  name = "license Unlock Hash"
  type = "int"
}
module vendor_item_list_maximum_level {
  source = "./attribute"

  name = "maximum Level"
  type = "int"
}
module vendor_item_list_minimum_level {
  source = "./attribute"

  name = "minimum Level"
  type = "int"
}
module vendor_item_list_original_category_index {
  source = "./attribute"

  name = "original Category Index"
  type = "int"
}
module vendor_item_list_price_override_enabled {
  source = "./attribute"

  name = "price Override Enabled"
  type = "bool"
}
module vendor_item_list_purchasable_scope {
  source = "./attribute"

  name = "purchasable Scope"
  type = "int"
}
module vendor_item_list_quantity {
  source = "./attribute"

  name = "quantity"
  type = "int"
}
module vendor_item_list_redirect_to_sale_indexes {
  source = "./attribute"

  name = "redirect To Sale Indexes"
  type = "[]interface{}"
}
module vendor_item_list_refund_policy {
  source = "./attribute"

  name = "refund Policy"
  type = "int"
}
module vendor_item_list_refund_time_limit {
  source = "./attribute"

  name = "refund Time Limit"
  type = "int"
}
module vendor_item_list_reward_adjustor_pointer_hash {
  source = "./attribute"

  name = "reward Adjustor Pointer Hash"
  type = "int"
}
module vendor_item_list_seed_override {
  source = "./attribute"

  name = "seed Override"
  type = "int"
}
module vendor_item_list_socket_overrides {
  source = "./attribute"

  name = "socket Overrides"
  type = "[]interface{}"
}
module vendor_item_list_sort_value {
  source = "./attribute"

  name = "sort Value"
  type = "int"
}
module vendor_item_list_vendor_item_index {
  source = "./attribute"

  name = "vendor Item Index"
  type = "int"
}
module vendor_item_list_visibility_scope {
  source = "./attribute"

  name = "visibility Scope"
  type = "int"
}
module vendor_item_list_weight {
  source = "./attribute"

  name = "weight"
  type = "float64"
}


module vendor_original_categories {
  source = "./struct"

  name = "original Categories"
  attributes_content = join("\n", [
    module.vendor_original_categories_buy_string_override.rendered,
    module.vendor_original_categories_category_hash.rendered,
    module.vendor_original_categories_category_index.rendered,
    module.vendor_original_categories_disabled_description.rendered,
    module.vendor_original_categories_hide_from_regular_purchase.rendered,
    module.vendor_original_categories_hide_if_no_currency.rendered,
    module.vendor_original_categories_is_display_only.rendered,
    module.vendor_original_categories_is_preview.rendered,
    module.vendor_original_categories_quantity_available.rendered,
    module.vendor_original_categories_reset_interval_minutes_override.rendered,
    module.vendor_original_categories_reset_offset_minutes_override.rendered,
    module.vendor_original_categories_show_unavailable_items.rendered,
    module.vendor_original_categories_sort_value.rendered,
    module.vendor_original_categories_vendor_item_indexes.rendered,
  ])
}

module vendor_original_categories_buy_string_override {
  source = "./attribute"

  name = "buy String Override"
  type = "string"
}
module vendor_original_categories_category_hash {
  source = "./attribute"

  name = "category Hash"
  type = "int"
}
module vendor_original_categories_category_index {
  source = "./attribute"

  name = "category Index"
  type = "int"
}
module vendor_original_categories_disabled_description {
  source = "./attribute"

  name = "disabled Description"
  type = "string"
}
module vendor_original_categories_hide_from_regular_purchase {
  source = "./attribute"

  name = "hide From Regular Purchase"
  type = "bool"
}
module vendor_original_categories_hide_if_no_currency {
  source = "./attribute"

  name = "hide If No Currency"
  type = "bool"
}
module vendor_original_categories_is_display_only {
  source = "./attribute"

  name = "is Display Only"
  type = "bool"
}
module vendor_original_categories_is_preview {
  source = "./attribute"

  name = "is Preview"
  type = "bool"
}
module vendor_original_categories_quantity_available {
  source = "./attribute"

  name = "quantity Available"
  type = "int"
}
module vendor_original_categories_reset_interval_minutes_override {
  source = "./attribute"

  name = "reset Interval Minutes Override"
  type = "int"
}
module vendor_original_categories_reset_offset_minutes_override {
  source = "./attribute"

  name = "reset Offset Minutes Override"
  type = "int"
}
module vendor_original_categories_show_unavailable_items {
  source = "./attribute"

  name = "show Unavailable Items"
  type = "bool"
}
module vendor_original_categories_sort_value {
  source = "./attribute"

  name = "sort Value"
  type = "int"
}
module vendor_original_categories_vendor_item_indexes {
  source = "./attribute"

  name = "vendor Item Indexes"
  type = "[]int"
}


