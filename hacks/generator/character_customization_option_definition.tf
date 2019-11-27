// Code generated; DANGER ZONE FOR EDITS

module character_customization_option_definition {
  source = "./definition"

  target = var.target

  name = "character Customization Option"
  structs_content = join("\n", [
    module.character_customization_option_character_customization_option_definition.rendered,
    module.character_customization_option_float_slice_options.rendered,
    module.character_customization_option_int_options.rendered,
    module.character_customization_option_decal_color_options.rendered,
    module.character_customization_option_decal_options.rendered,
    module.character_customization_option_eye_color_options.rendered,
    module.character_customization_option_face_options.rendered,
    module.character_customization_option_feature_color_options.rendered,
    module.character_customization_option_feature_options.rendered,
    module.character_customization_option_hair_color_options.rendered,
    module.character_customization_option_hair_options.rendered,
    module.character_customization_option_helmet_preferences.rendered,
    module.character_customization_option_lip_color_options.rendered,
    module.character_customization_option_personality_options.rendered,
    module.character_customization_option_skin_color_options.rendered,
  ])

}

module character_customization_option_character_customization_option_definition {
  source = "./struct"

  name = "character Customization Option Definition"
  attributes_content = join("\n", [
    module.character_customization_option_character_customization_option_definition_blacklisted.rendered,
    module.character_customization_option_character_customization_option_definition_decal_color_options.rendered,
    module.character_customization_option_character_customization_option_definition_decal_options.rendered,
    module.character_customization_option_character_customization_option_definition_display_properties.rendered,
    module.character_customization_option_character_customization_option_definition_eye_color_options.rendered,
    module.character_customization_option_character_customization_option_definition_face_option_cinematic_host_pattern_ids.rendered,
    module.character_customization_option_character_customization_option_definition_face_options.rendered,
    module.character_customization_option_character_customization_option_definition_feature_color_options.rendered,
    module.character_customization_option_character_customization_option_definition_feature_options.rendered,
    module.character_customization_option_character_customization_option_definition_gender_hash.rendered,
    module.character_customization_option_character_customization_option_definition_hair_color_options.rendered,
    module.character_customization_option_character_customization_option_definition_hair_options.rendered,
    module.character_customization_option_character_customization_option_definition_hash.rendered,
    module.character_customization_option_character_customization_option_definition_helmet_preferences.rendered,
    module.character_customization_option_character_customization_option_definition_index.rendered,
    module.character_customization_option_character_customization_option_definition_lip_color_options.rendered,
    module.character_customization_option_character_customization_option_definition_personality_options.rendered,
    module.character_customization_option_character_customization_option_definition_race_hash.rendered,
    module.character_customization_option_character_customization_option_definition_redacted.rendered,
    module.character_customization_option_character_customization_option_definition_skin_color_options.rendered,
  ])
}

module character_customization_option_character_customization_option_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module character_customization_option_character_customization_option_definition_decal_color_options {
  source = "./attribute"

  name = "decal Color Options"
  type = "DecalColorOptions"
}
module character_customization_option_character_customization_option_definition_decal_options {
  source = "./attribute"

  name = "decal Options"
  type = "DecalOptions"
}
module character_customization_option_character_customization_option_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module character_customization_option_character_customization_option_definition_eye_color_options {
  source = "./attribute"

  name = "eye Color Options"
  type = "EyeColorOptions"
}
module character_customization_option_character_customization_option_definition_face_option_cinematic_host_pattern_ids {
  source = "./attribute"

  name = "face Option Cinematic Host Pattern Ids"
  type = "[]interface{}"
}
module character_customization_option_character_customization_option_definition_face_options {
  source = "./attribute"

  name = "face Options"
  type = "FaceOptions"
}
module character_customization_option_character_customization_option_definition_feature_color_options {
  source = "./attribute"

  name = "feature Color Options"
  type = "FeatureColorOptions"
}
module character_customization_option_character_customization_option_definition_feature_options {
  source = "./attribute"

  name = "feature Options"
  type = "FeatureOptions"
}
module character_customization_option_character_customization_option_definition_gender_hash {
  source = "./attribute"

  name = "gender Hash"
  type = "int64"
}
module character_customization_option_character_customization_option_definition_hair_color_options {
  source = "./attribute"

  name = "hair Color Options"
  type = "HairColorOptions"
}
module character_customization_option_character_customization_option_definition_hair_options {
  source = "./attribute"

  name = "hair Options"
  type = "HairOptions"
}
module character_customization_option_character_customization_option_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module character_customization_option_character_customization_option_definition_helmet_preferences {
  source = "./attribute"

  name = "helmet Preferences"
  type = "HelmetPreferences"
}
module character_customization_option_character_customization_option_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module character_customization_option_character_customization_option_definition_lip_color_options {
  source = "./attribute"

  name = "lip Color Options"
  type = "LipColorOptions"
}
module character_customization_option_character_customization_option_definition_personality_options {
  source = "./attribute"

  name = "personality Options"
  type = "PersonalityOptions"
}
module character_customization_option_character_customization_option_definition_race_hash {
  source = "./attribute"

  name = "race Hash"
  type = "int"
}
module character_customization_option_character_customization_option_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module character_customization_option_character_customization_option_definition_skin_color_options {
  source = "./attribute"

  name = "skin Color Options"
  type = "SkinColorOptions"
}


module character_customization_option_float_slice_options {
  source = "./struct"

  name = "float Slice Options"
  attributes_content = join("\n", [
    module.character_customization_option_float_slice_options_display_properties.rendered,
    module.character_customization_option_float_slice_options_value.rendered,
  ])
}

module character_customization_option_float_slice_options_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module character_customization_option_float_slice_options_value {
  source = "./attribute"

  name = "value"
  type = "[]float64"
}


module character_customization_option_int_options {
  source = "./struct"

  name = "int Options"
  attributes_content = join("\n", [
    module.character_customization_option_int_options_display_properties.rendered,
    module.character_customization_option_int_options_value.rendered,
  ])
}

module character_customization_option_int_options_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module character_customization_option_int_options_value {
  source = "./attribute"

  name = "value"
  type = "int64"
}


module character_customization_option_decal_color_options {
  source = "./struct"

  name = "decal Color Options"
  attributes_content = join("\n", [
    module.character_customization_option_decal_color_options_customization_category_hash.rendered,
    module.character_customization_option_decal_color_options_display_name.rendered,
    module.character_customization_option_decal_color_options_options.rendered,
  ])
}

module character_customization_option_decal_color_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int"
}
module character_customization_option_decal_color_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_decal_color_options_options {
  source = "./attribute"

  name = "options"
  type = "[]IntOptions"
}


module character_customization_option_decal_options {
  source = "./struct"

  name = "decal Options"
  attributes_content = join("\n", [
    module.character_customization_option_decal_options_customization_category_hash.rendered,
    module.character_customization_option_decal_options_display_name.rendered,
    module.character_customization_option_decal_options_options.rendered,
  ])
}

module character_customization_option_decal_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int"
}
module character_customization_option_decal_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_decal_options_options {
  source = "./attribute"

  name = "options"
  type = "[]IntOptions"
}


module character_customization_option_eye_color_options {
  source = "./struct"

  name = "eye Color Options"
  attributes_content = join("\n", [
    module.character_customization_option_eye_color_options_customization_category_hash.rendered,
    module.character_customization_option_eye_color_options_display_name.rendered,
    module.character_customization_option_eye_color_options_options.rendered,
  ])
}

module character_customization_option_eye_color_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int64"
}
module character_customization_option_eye_color_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_eye_color_options_options {
  source = "./attribute"

  name = "options"
  type = "[]IntOptions"
}


module character_customization_option_face_options {
  source = "./struct"

  name = "face Options"
  attributes_content = join("\n", [
    module.character_customization_option_face_options_customization_category_hash.rendered,
    module.character_customization_option_face_options_display_name.rendered,
    module.character_customization_option_face_options_options.rendered,
  ])
}

module character_customization_option_face_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int64"
}
module character_customization_option_face_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_face_options_options {
  source = "./attribute"

  name = "options"
  type = "[]IntOptions"
}


module character_customization_option_feature_color_options {
  source = "./struct"

  name = "feature Color Options"
  attributes_content = join("\n", [
    module.character_customization_option_feature_color_options_customization_category_hash.rendered,
    module.character_customization_option_feature_color_options_display_name.rendered,
    module.character_customization_option_feature_color_options_options.rendered,
  ])
}

module character_customization_option_feature_color_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int"
}
module character_customization_option_feature_color_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_feature_color_options_options {
  source = "./attribute"

  name = "options"
  type = "[]interface{}"
}


module character_customization_option_feature_options {
  source = "./struct"

  name = "feature Options"
  attributes_content = join("\n", [
    module.character_customization_option_feature_options_customization_category_hash.rendered,
    module.character_customization_option_feature_options_display_name.rendered,
    module.character_customization_option_feature_options_options.rendered,
  ])
}

module character_customization_option_feature_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int"
}
module character_customization_option_feature_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_feature_options_options {
  source = "./attribute"

  name = "options"
  type = "[]IntOptions"
}


module character_customization_option_hair_color_options {
  source = "./struct"

  name = "hair Color Options"
  attributes_content = join("\n", [
    module.character_customization_option_hair_color_options_customization_category_hash.rendered,
    module.character_customization_option_hair_color_options_display_name.rendered,
    module.character_customization_option_hair_color_options_options.rendered,
  ])
}

module character_customization_option_hair_color_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "float64"
}
module character_customization_option_hair_color_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_hair_color_options_options {
  source = "./attribute"

  name = "options"
  type = "[]FloatSliceOptions"
}


module character_customization_option_hair_options {
  source = "./struct"

  name = "hair Options"
  attributes_content = join("\n", [
    module.character_customization_option_hair_options_customization_category_hash.rendered,
    module.character_customization_option_hair_options_display_name.rendered,
    module.character_customization_option_hair_options_options.rendered,
  ])
}

module character_customization_option_hair_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int"
}
module character_customization_option_hair_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_hair_options_options {
  source = "./attribute"

  name = "options"
  type = "[]IntOptions"
}


module character_customization_option_helmet_preferences {
  source = "./struct"

  name = "helmet Preferences"
  attributes_content = join("\n", [
    module.character_customization_option_helmet_preferences_customization_category_hash.rendered,
    module.character_customization_option_helmet_preferences_display_name.rendered,
    module.character_customization_option_helmet_preferences_options.rendered,
  ])
}

module character_customization_option_helmet_preferences_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int"
}
module character_customization_option_helmet_preferences_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_helmet_preferences_options {
  source = "./attribute"

  name = "options"
  type = "[]interface{}"
}


module character_customization_option_lip_color_options {
  source = "./struct"

  name = "lip Color Options"
  attributes_content = join("\n", [
    module.character_customization_option_lip_color_options_customization_category_hash.rendered,
    module.character_customization_option_lip_color_options_display_name.rendered,
    module.character_customization_option_lip_color_options_options.rendered,
  ])
}

module character_customization_option_lip_color_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int64"
}
module character_customization_option_lip_color_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_lip_color_options_options {
  source = "./attribute"

  name = "options"
  type = "[]IntOptions"
}


module character_customization_option_personality_options {
  source = "./struct"

  name = "personality Options"
  attributes_content = join("\n", [
    module.character_customization_option_personality_options_customization_category_hash.rendered,
    module.character_customization_option_personality_options_display_name.rendered,
    module.character_customization_option_personality_options_options.rendered,
  ])
}

module character_customization_option_personality_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int"
}
module character_customization_option_personality_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_personality_options_options {
  source = "./attribute"

  name = "options"
  type = "[]interface{}"
}


module character_customization_option_skin_color_options {
  source = "./struct"

  name = "skin Color Options"
  attributes_content = join("\n", [
    module.character_customization_option_skin_color_options_customization_category_hash.rendered,
    module.character_customization_option_skin_color_options_display_name.rendered,
    module.character_customization_option_skin_color_options_options.rendered,
  ])
}

module character_customization_option_skin_color_options_customization_category_hash {
  source = "./attribute"

  name = "customization Category Hash"
  type = "int64"
}
module character_customization_option_skin_color_options_display_name {
  source = "./attribute"

  name = "display Name"
  type = "string"
}
module character_customization_option_skin_color_options_options {
  source = "./attribute"

  name = "options"
  type = "[]IntOptions"
}


