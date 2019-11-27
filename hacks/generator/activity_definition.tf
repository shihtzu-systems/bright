// Code generated; DANGER ZONE FOR EDITS

module activity_definition {
  source = "./definition"

  target = var.target

  name = "activity"
  structs_content = join("\n", [
    module.activity_activity_definition.rendered,
    module.activity_matchmaking.rendered,
    module.activity_original_display_properties.rendered,
  ])

}

module activity_activity_definition {
  source = "./struct"

  name = "activity Definition"
  attributes_content = join("\n", [
    module.activity_activity_definition_activity_level.rendered,
    module.activity_activity_definition_activity_light_level.rendered,
    module.activity_activity_definition_activity_location_mappings.rendered,
    module.activity_activity_definition_activity_mode_hashes.rendered,
    module.activity_activity_definition_activity_mode_types.rendered,
    module.activity_activity_definition_activity_type_hash.rendered,
    module.activity_activity_definition_blacklisted.rendered,
    module.activity_activity_definition_challenges.rendered,
    module.activity_activity_definition_completion_unlock_hash.rendered,
    module.activity_activity_definition_destination_hash.rendered,
    module.activity_activity_definition_direct_activity_mode_hash.rendered,
    module.activity_activity_definition_direct_activity_mode_type.rendered,
    module.activity_activity_definition_display_properties.rendered,
    module.activity_activity_definition_hash.rendered,
    module.activity_activity_definition_index.rendered,
    module.activity_activity_definition_inherit_from_free_roam.rendered,
    module.activity_activity_definition_insertion_points.rendered,
    module.activity_activity_definition_is_playlist.rendered,
    module.activity_activity_definition_is_pvp.rendered,
    module.activity_activity_definition_matchmaking.rendered,
    module.activity_activity_definition_modifiers.rendered,
    module.activity_activity_definition_optional_unlock_strings.rendered,
    module.activity_activity_definition_original_display_properties.rendered,
    module.activity_activity_definition_pgcr_image.rendered,
    module.activity_activity_definition_place_hash.rendered,
    module.activity_activity_definition_playlist_items.rendered,
    module.activity_activity_definition_redacted.rendered,
    module.activity_activity_definition_release_icon.rendered,
    module.activity_activity_definition_release_time.rendered,
    module.activity_activity_definition_rewards.rendered,
    module.activity_activity_definition_suppress_other_rewards.rendered,
    module.activity_activity_definition_tier.rendered,
  ])
}

module activity_activity_definition_activity_level {
  source = "./attribute"

  name = "activity Level"
  type = "int"
}
module activity_activity_definition_activity_light_level {
  source = "./attribute"

  name = "activity Light Level"
  type = "int"
}
module activity_activity_definition_activity_location_mappings {
  source = "./attribute"

  name = "activity Location Mappings"
  type = "[]interface{}"
}
module activity_activity_definition_activity_mode_hashes {
  source = "./attribute"

  name = "activity Mode Hashes"
  type = "[]int"
}
module activity_activity_definition_activity_mode_types {
  source = "./attribute"

  name = "activity Mode Types"
  type = "[]int"
}
module activity_activity_definition_activity_type_hash {
  source = "./attribute"

  name = "activity Type Hash"
  type = "int"
}
module activity_activity_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module activity_activity_definition_challenges {
  source = "./attribute"

  name = "challenges"
  type = "[]interface{}"
}
module activity_activity_definition_completion_unlock_hash {
  source = "./attribute"

  name = "completion Unlock Hash"
  type = "int"
}
module activity_activity_definition_destination_hash {
  source = "./attribute"

  name = "destination Hash"
  type = "int"
}
module activity_activity_definition_direct_activity_mode_hash {
  source = "./attribute"

  name = "direct Activity Mode Hash"
  type = "int"
}
module activity_activity_definition_direct_activity_mode_type {
  source = "./attribute"

  name = "direct Activity Mode Type"
  type = "int"
}
module activity_activity_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module activity_activity_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int64"
}
module activity_activity_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module activity_activity_definition_inherit_from_free_roam {
  source = "./attribute"

  name = "inherit From Free Roam"
  type = "bool"
}
module activity_activity_definition_insertion_points {
  source = "./attribute"

  name = "insertion Points"
  type = "[]interface{}"
}
module activity_activity_definition_is_playlist {
  source = "./attribute"

  name = "is Playlist"
  type = "bool"
}
module activity_activity_definition_is_pvp {
  source = "./attribute"

  name = "is PvP"
  type = "bool"
}
module activity_activity_definition_matchmaking {
  source = "./attribute"

  name = "matchmaking"
  type = "Matchmaking"
}
module activity_activity_definition_modifiers {
  source = "./attribute"

  name = "modifiers"
  type = "[]interface{}"
}
module activity_activity_definition_optional_unlock_strings {
  source = "./attribute"

  name = "optional Unlock Strings"
  type = "[]interface{}"
}
module activity_activity_definition_original_display_properties {
  source = "./attribute"

  name = "original Display Properties"
  type = "OriginalDisplayProperties"
}
module activity_activity_definition_pgcr_image {
  source = "./attribute"

  name = "pgcr Image"
  type = "string"
}
module activity_activity_definition_place_hash {
  source = "./attribute"

  name = "place Hash"
  type = "int64"
}
module activity_activity_definition_playlist_items {
  source = "./attribute"

  name = "playlist Items"
  type = "[]interface{}"
}
module activity_activity_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module activity_activity_definition_release_icon {
  source = "./attribute"

  name = "release Icon"
  type = "string"
}
module activity_activity_definition_release_time {
  source = "./attribute"

  name = "release Time"
  type = "int"
}
module activity_activity_definition_rewards {
  source = "./attribute"

  name = "rewards"
  type = "[]interface{}"
}
module activity_activity_definition_suppress_other_rewards {
  source = "./attribute"

  name = "suppress Other Rewards"
  type = "bool"
}
module activity_activity_definition_tier {
  source = "./attribute"

  name = "tier"
  type = "int"
}


module activity_matchmaking {
  source = "./struct"

  name = "matchmaking"
  attributes_content = join("\n", [
    module.activity_matchmaking_is_matchmade.rendered,
    module.activity_matchmaking_max_party.rendered,
    module.activity_matchmaking_max_players.rendered,
    module.activity_matchmaking_min_party.rendered,
    module.activity_matchmaking_requires_guardian_oath.rendered,
  ])
}

module activity_matchmaking_is_matchmade {
  source = "./attribute"

  name = "is Matchmade"
  type = "bool"
}
module activity_matchmaking_max_party {
  source = "./attribute"

  name = "max Party"
  type = "int"
}
module activity_matchmaking_max_players {
  source = "./attribute"

  name = "max Players"
  type = "int"
}
module activity_matchmaking_min_party {
  source = "./attribute"

  name = "min Party"
  type = "int"
}
module activity_matchmaking_requires_guardian_oath {
  source = "./attribute"

  name = "requires Guardian Oath"
  type = "bool"
}


module activity_original_display_properties {
  source = "./struct"

  name = "original Display Properties"
  attributes_content = join("\n", [
    module.activity_original_display_properties_description.rendered,
    module.activity_original_display_properties_has_icon.rendered,
    module.activity_original_display_properties_icon.rendered,
    module.activity_original_display_properties_name.rendered,
  ])
}

module activity_original_display_properties_description {
  source = "./attribute"

  name = "description"
  type = "string"
}
module activity_original_display_properties_has_icon {
  source = "./attribute"

  name = "has Icon"
  type = "bool"
}
module activity_original_display_properties_icon {
  source = "./attribute"

  name = "icon"
  type = "string"
}
module activity_original_display_properties_name {
  source = "./attribute"

  name = "name"
  type = "string"
}


