resource local_file this {
  for_each = {

    achievement = {
      name = "achievement"
      structs = [
        {
          name = "achievement Definition"
          attrs = [
            {
              name = "acccumulator Threshold"
              type = "int"

              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"

            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "platform Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    activity = {
      name = "activity"
      structs = [
        {
          name = "activity Definition"
          attrs = [
            {
              name    = "activity Level"
              type    = "int"
              display = "none"
            },
            {
              name    = "activity Light Level"
              type    = "int"
              display = "int"
            },
            {
              name    = "activity Location Mappings"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "activity Mode Hashes"
              type    = "[]int"
              display = "none"
            },
            {
              name    = "activity Mode Types"
              type    = "[]int"
              display = "none"
            },
            {
              name    = "activity Type Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "challenges"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "completion Unlock Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "destination Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "direct Activity Mode Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "direct Activity Mode Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "inherit From Free Roam"
              type    = "bool"
              display = "none"
            },
            {
              name    = "insertion Points"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "is Playlist"
              type    = "bool"
              display = "default"
            },
            {
              name    = "is PvP"
              type    = "bool"
              display = "default"
            },
            {
              name    = "matchmaking"
              type    = "Matchmaking"
              display = "nested"
            },
            {
              name    = "modifiers"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "optional Unlock Strings"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "original Display Properties"
              type    = "OriginalDisplayProperties"
              display = "none"
            },
            {
              name    = "pgcr Image"
              type    = "string"
              display = "image"
            },
            {
              name    = "place Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "playlist Items"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "release Icon"
              type    = "string"
              display = "none"
            },
            {
              name    = "release Time"
              type    = "int"
              display = "none"
            },
            {
              name    = "rewards"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "suppress Other Rewards"
              type    = "bool"
              display = "none"
            },
            {
              name    = "tier"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "matchmaking"
          attrs = [
            {
              name    = "is Matchmade"
              type    = "bool"
              display = "default"
            },
            {
              name    = "max Party"
              type    = "int"
              display = "default"
            },
            {
              name    = "max Players"
              type    = "int"
              display = "default"
            },
            {
              name    = "min Party"
              type    = "int"
              display = "default"
            },
            {
              name    = "requires Guardian Oath"
              type    = "bool"
              display = "default"
            },
          ]
        },
        {
          name = "original Display Properties"
          attrs = [
            {
              name    = "description"
              type    = "string"
              display = "none"
            },
            {
              name    = "has Icon"
              type    = "bool"
              display = "none"
            },
            {
              name    = "icon"
              type    = "string"
              display = "none"
            },
            {
              name    = "name"
              type    = "string"
              display = "none"
            },
          ]
        }
      ]
    }

    activity_graph = {
      name = "activity Graph"
      structs = [
        {
          name = "activity Graph Definition"
          attrs = [

            {
              name    = "art Elements"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "connections"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "display Objectives"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "display Progressions"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "ignore For Milestones"
              type    = "bool"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "linked Graphs"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "nodes"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "ui Screen"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    activity_interactable = {
      name = "activity Interactable"
      structs = [
        {
          name = "activity Interactable Definition"
          attrs = [

            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "entries"
              type    = "[]ActivityInteractableEntries"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "activity Interactable Entries"
          attrs = [
            {
              name    = "activity Hash"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    activity_mode = {
      name = "activity Mode"
      structs = [
        {
          name = "activity Mode Definition"
          attrs = [
            {
              name    = "activity Mode Category"
              type    = "int"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display"
              type    = "bool"
              display = "default"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "friendly Name"
              type    = "string"
              display = "default"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "is Aggregate Mode"
              type    = "bool"
              display = "default"
            },
            {
              name    = "is Team Based"
              type    = "bool"
              display = "default"
            },
            {
              name    = "mode Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "order"
              type    = "int"
              display = "none"
            },
            {
              name    = "parent Hashes"
              type    = "[]int"
              display = "none"
            },
            {
              name    = "pgcr Image"
              type    = "string"
              display = "image"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "supports Feed Filtering"
              type    = "bool"
              display = "none"
            },
            {
              name    = "tier"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    activity_modifier = {
      name = "activity Modifier"
      structs = [
        {
          name = "activity Modifier Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    activity_type = {
      name = "activity Type"
      structs = [
        {
          name = "activity Type Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    art_dye_channel = {
      name = "art Dye Channel"
      structs = [
        {
          name = "art Dye Channel Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "channel Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    art_dye_reference = {
      name = "art Dye Reference"
      structs = [
        {
          name = "art Dye Reference Definition"
          attrs = [
            {
              name    = "art Dye Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "dye Manifest Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    artifact = {
      name = "artifact"
      structs = [
        {
          name = "artifact Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "tiers"
              type    = "[]Tiers"
              display = "none"
            },
            {
              name    = "translation Block"
              type    = "TranslationBlock"
              display = "none"
            },
          ]
        },
        {
          name = "artifact Items"
          attrs = [
            {
              name    = "active Unlock Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "item Hash"
              type    = "int64"
              display = "none"
            },
          ]
        },
        {
          name = "tiers"
          attrs = [
            {
              name    = "display Title"
              type    = "string"
              display = "none"
            },
            {
              name    = "items"
              type    = "[]ArtifactItems"
              display = "none"
            },
            {
              name    = "minimum Unlock Points Used Requirement"
              type    = "int"
              display = "none"
            },
            {
              name    = "progress Requirement Message"
              type    = "string"
              display = "none"
            },
            {
              name    = "tier Hash"
              type    = "int64"
              display = "none"
            },
          ]
        }
      ]
    }

    bond = {
      name = "bond"
      structs = [
        {
          name = "bond Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "provided Unlock Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "provided Unlock Value Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    breaker_type = {
      name = "breaker Type"
      structs = [
        {
          name = "breaker Type Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "enum Value"
              type    = "int"
              display = "default"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "unlock Hash"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    character_customization_category = {
      name = "character Customization Category"
      structs = [
        {
          name = "character Customization Category Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    character_customization_option = {
      name = "character Customization Option"
      structs = [
        {
          name = "character Customization Option Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "decal Color Options"
              type    = "DecalColorOptions"
              display = "none"
            },
            {
              name    = "decal Options"
              type    = "DecalOptions"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "eye Color Options"
              type    = "EyeColorOptions"
              display = "none"
            },
            {
              name    = "face Option Cinematic Host Pattern Ids"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "face Options"
              type    = "FaceOptions"
              display = "none"
            },
            {
              name    = "feature Color Options"
              type    = "FeatureColorOptions"
              display = "none"
            },
            {
              name    = "feature Options"
              type    = "FeatureOptions"
              display = "none"
            },
            {
              name    = "gender Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "hair Color Options"
              type    = "HairColorOptions"
              display = "none"
            },
            {
              name    = "hair Options"
              type    = "HairOptions"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "helmet Preferences"
              type    = "HelmetPreferences"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "lip Color Options"
              type    = "LipColorOptions"
              display = "none"
            },
            {
              name    = "personality Options"
              type    = "PersonalityOptions"
              display = "none"
            },
            {
              name    = "race Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "skin Color Options"
              type    = "SkinColorOptions"
              display = "none"
            },
          ]
        },
        {
          name = "float Slice Options"
          attrs = [
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "value"
              type    = "[]float64"
              display = "none"
            },
          ]
        },
        {
          name = "int Options"
          attrs = [
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "value"
              type    = "int64"
              display = "none"
            },
          ]
        },
        {
          name = "decal Color Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]IntOptions"
              display = "none"
            },
          ]
        },
        {
          name = "decal Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]IntOptions"
              display = "none"
            },
          ]
        },
        {
          name = "eye Color Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]IntOptions"
              display = "none"
            },
          ]
        },
        {
          name = "face Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]IntOptions"
              display = "none"
            },
          ]
        },
        {
          name = "feature Color Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]interface{}"
              display = "none"
            },
          ]
        },
        {
          name = "feature Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]IntOptions"
              display = "none"
            },
          ]
        },
        {
          name = "hair Color Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "float64"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]FloatSliceOptions"
              display = "none"
            },
          ]
        },
        {
          name = "hair Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]IntOptions"
              display = "none"
            },
          ]
        },
        {
          name = "helmet Preferences"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]interface{}"
              display = "none"
            },
          ]
        },
        {
          name = "lip Color Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]IntOptions"
              display = "none"
            },
          ]
        },
        {
          name = "personality Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]interface{}"
              display = "none"
            },
          ]
        },
        {
          name = "skin Color Options"
          attrs = [
            {
              name    = "customization Category Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "display Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "options"
              type    = "[]IntOptions"
              display = "none"
            },
          ]
        },
      ]
    }

    checklist = {
      name = "checklist"
      structs = [
        {
          name = "checklist Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "entries"
              type    = "[]ChecklistEntries"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "scope"
              type    = "int"
              display = "none"
            },
            {
              name    = "view Action String"
              type    = "string"
              display = "default"
            },
          ]
        },
        {
          name = "checklist Entries"
          attrs = [
            {
              name    = "bubble Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "destination Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "location Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "scope"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    class = {
      name = "class"
      structs = [
        {
          name = "class Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "class Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "gendered Class Names"
              type    = "GenderedClassNames"
              display = "none"
            },
            {
              name    = "gendered Class Names By Gender Hash"
              type    = "map[string]interface{}"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "gendered class Names"
          attrs = [
            {
              name    = "female"
              type    = "string"
              display = "none"
            },
            {
              name    = "male"
              type    = "string"
              display = "none"
            },
          ]
        }
      ]
    }

    collectible = {
      name = "collectible"
      structs = [
        {
          name = "collectible Definition"
          attrs = [
            {
              name    = "acquisition Info"
              type    = "AcquisitionInfo"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "item Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "presentation Info"
              type    = "CollectiblePresentationInfo"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "scope"
              type    = "int"
              display = "none"
            },
            {
              name    = "source Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "source String"
              type    = "string"
              display = "default"
            },
            {
              name    = "state Info"
              type    = "CollectibleStateInfo"
              display = "none"
            },
          ]
        },
        {
          name = "acquisition Info"
          attrs = [
            {
              name    = "acquire Material Requirement Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "run Only Acquisition Reward Site"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "collectible Presentation Info"
          attrs = [
            {
              name    = "display Style"
              type    = "int"
              display = "none"
            },
            {
              name    = "parent Presentation Node Hashes"
              type    = "[]int"
              display = "none"
            },
            {
              name    = "presentation Node Type"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "collectible Requirements"
          attrs = [
            {
              name    = "entitlement Unavailable Message"
              type    = "string"
              display = "none"
            },
          ]
        },
        {
          name = "collectible State Info"
          attrs = [
            {
              name    = "requirements"
              type    = "CollectibleRequirements"
              display = "none"
            },
          ]
        }
      ]
    }

    damage_type = {
      name = "damage Type"
      structs = [
        {
          name = "damage Type Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "enum Value"
              type    = "int"
              display = "default"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "show Icon"
              type    = "bool"
              display = "default"
            },
            {
              name    = "transparent Icon Path"
              type    = "string"
              display = "icon"
            },
          ]
        }
      ]
    }

    destination = {
      name = "destination"
      structs = [
        {
          name = "destination Definition"
          attrs = [
            {
              name    = "activity Graph Entries"
              type    = "[]interface{} "
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "bubble Settings"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "bubbles"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "default Freeroam Activity Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "place Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    enemy_race = {
      name = "enemy Race"
      structs = [
        {
          name = "enemy Race Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    energy_type = {
      name = "energy Type"
      structs = [
        {
          name = "energy Type Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "capacity Stat Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "cost Stat Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "enum Value"
              type    = "int"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "show Icon"
              type    = "bool"
              display = "none"
            },
            {
              name    = "transparent Icon Path"
              type    = "string"
              display = "none"
            },
          ]
        }
      ]
    }

    entitlement_offer = {
      name = "entitlement Offer"
      structs = [
        {
          name = "entitlement Offer Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "offer Key"
              type    = "string"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    equipment_slot = {
      name = "equipment Slot"
      structs = [
        {
          name = "equipment Slot Definition"
          attrs = [
            {
              name    = "apply Custom Art Dyes"
              type    = "bool"
              display = "none"
            },
            {
              name    = "art Dye Channels"
              type    = "[]ArtDyeChannels"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "bucket Type Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "equipment Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "art Dye Channels"
          attrs = [
            {
              name    = "art Dye Channel Hash"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    faction = {
      name = "faction"
      structs = [
        {
          name = "faction Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "progression Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "reward Item Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "reward Vendor Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "token Values"
              type    = "map[string]interface{}"
              display = "none"
            },
            {
              name    = "vendors"
              type    = "[]Vendors"
              display = "none"
            },
          ]
        },
        {
          name = "vendors"
          attrs = [
            {
              name    = "background Image Path"
              type    = "string"
              display = "image"
            },
            {
              name    = "destination Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "vendor Hash"
              type    = "int64"
              display = "none"
            },
          ]
        }
      ]
    }

    gender = {
      name = "gender"
      structs = [
        {
          name = "gender Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "gender Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    inventory_bucket = {
      name = "inventory Bucket"
      structs = [
        {
          name = "inventory Bucket Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "bucketOrder"
              type    = "int"
              display = "none"
            },
            {
              name    = "category"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "enabled"
              type    = "bool"
              display = "default"
            },
            {
              name    = "fifo"
              type    = "bool"
              display = "default"
            },
            {
              name    = "has Transfer Destination"
              type    = "bool"
              display = "default"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "item Count"
              type    = "int"
              display = "default"
            },
            {
              name    = "location"
              type    = "int"
              display = "default"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "scope"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    inventory_item = {
      name = "inventory Item"
      structs = [
        {
          name = "inventory Item Definition"
          attrs = [
            {
              name    = "acquire Reward Site Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "acquire Unlock Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "action"
              type    = "InventoryItemAction"
              display = "nested"
            },
            {
              name    = "allow Actions"
              type    = "bool"
              display = "default"
            },
            {
              name    = "background Color"
              type    = "BackgroundColor"
              display = "nested"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "breaker Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "class Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "default Damage Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "display Source"
              type    = "string"
              display = "none"
            },
            {
              name    = "does Postmaster Pull Have Side Effects"
              type    = "bool"
              display = "none"
            },
            {
              name    = "equippable"
              type    = "bool"
              display = "deault"
            },
            {
              name    = "equipping Block"
              type    = "EquippingBlock"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "inventory"
              type    = "Inventory"
              display = "nested"
            },
            {
              name    = "investment Stats"
              type    = "[]InvestmentStats"
              display = "none"
            },
            {
              name    = "is Wrapper"
              type    = "bool"
              display = "deault"
            },
            {
              name    = "item Category Hashes"
              type    = "[]int"
              display = "none"
            },
            {
              name    = "item Sub Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "item Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "item Type And Tier Display Name"
              type    = "string"
              display = "default"
            },
            {
              name    = "item Type Display Name"
              type    = "string"
              display = "default"
            },
            {
              name    = "non Transferrable"
              type    = "bool"
              display = "default"
            },
            {
              name    = "perks"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "preview"
              type    = "Preview"
              display = "nested"
            },
            {
              name    = "quality"
              type    = "Quality"
              display = "nested"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "secondary Icon"
              type    = "string"
              display = "image"
            },
            {
              name    = "secondary Overlay"
              type    = "string"
              display = "image"
            },
            {
              name    = "secondary Special"
              type    = "string"
              display = "image"
            },
            {
              name    = "screenshot"
              type    = "string"
              display = "image"
            },
            {
              name    = "sockets"
              type    = "Sockets"
              display = "nested"
            },
            {
              name    = "source Data"
              type    = "SourceData"
              display = "none"
            },
            {
              name    = "special Item Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "stats"
              type    = "Stats"
              display = "none"
            },
            {
              name    = "summary Item Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "talent Grid"
              type    = "TalentGrid"
              display = "nested"
            },
            {
              name    = "tooltip Notifications"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "translation Block"
              type    = "TranslationBlock"
              display = "none"
            },
            {
              name    = "ui Item Display Style"
              type    = "string"
              display = "none"
            },
          ]
        },
        {
          name = "inventory Item Action"
          attrs = [
            {
              name    = "action Type Label"
              type    = "string"
              display = "default"
            },
            {
              name    = "consume Entire Stack"
              type    = "bool"
              display = "default"
            },
            {
              name    = "delete On Action"
              type    = "bool"
              display = "default"
            },
            {
              name    = "is Positive"
              type    = "bool"
              display = "default"
            },
            {
              name    = "progression Rewards"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "required Cooldown Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "required Cooldown Seconds"
              type    = "int"
              display = "none"
            },
            {
              name    = "required Items"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "reward Item Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "reward Sheet Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "reward Site Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "use On Acquire"
              type    = "bool"
              display = "default"
            },
            {
              name    = "verb Description"
              type    = "string"
              display = "default"
            },
            {
              name    = "verb Name"
              type    = "string"
              display = "default"
            },
          ]
        },
        {
          name = "background Color"
          attrs = [
            {
              name    = "alpha"
              type    = "int"
              display = "default"
            },
            {
              name    = "blue"
              type    = "int"
              display = "default"
            },
            {
              name    = "colorHash"
              type    = "int"
              display = "default"
            },
            {
              name    = "green"
              type    = "int"
              display = "default"
            },
            {
              name    = "red"
              type    = "int"
              display = "default"
            },
          ]
        },
        {
          name = "equipping Block"
          attrs = [
            {
              name    = "ammo Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "attributes"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Strings"
              type    = "[]string"
              display = "none"
            },
            {
              name    = "equipment Slot Type Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "equipping Sound Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "horn Sound Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "unique Label Hash"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "inventory"
          attrs = [
            {
              name    = "bucket Type Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "expiration Tooltip"
              type    = "string"
              display = "default"
            },
            {
              name    = "expired In Activity Message"
              type    = "string"
              display = "default"
            },
            {
              name    = "expired In Orbit Message"
              type    = "string"
              display = "default"
            },
            {
              name    = "is Instance Item"
              type    = "bool"
              display = "default"
            },
            {
              name    = "max Stack Size"
              type    = "int"
              display = "default"
            },
            {
              name    = "non Transferrable Original"
              type    = "bool"
              display = "default"
            },
            {
              name    = "recovery Bucket Type Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "suppress Expiration When Objectives Complete"
              type    = "bool"
              display = "default"
            },
            {
              name    = "tier Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "tier Type Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "tier Type Name"
              type    = "string"
              display = "default"
            },
          ]
        },
        {
          name = "investment Stats"
          attrs = [
            {
              name    = "is Conditionally Active"
              type    = "bool"
              display = "none"
            },
            {
              name    = "stat Type Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "value"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "preview"
          attrs = [
            {
              name    = "preview Action String"
              type    = "string"
              display = "default"
            },
            {
              name    = "preview Vendor Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "screen Style"
              type    = "string"
              display = "default"
            },
          ]
        },
        {
          name = "quality"
          attrs = [
            {
              name    = "infusion Category Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "infusion Category Hashes"
              type    = "[]int64"
              display = "none"
            },
            {
              name    = "infusion Category Name"
              type    = "string"
              display = "default"
            },
            {
              name    = "item Levels"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "progression Level Requirement Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "quality Level"
              type    = "int"
              display = "default"
            },
          ]
        },
        {
          name = "intrinsic Sockets"
          attrs = [
            {
              name    = "default Visible"
              type    = "bool"
              display = "default"
            },
            {
              name    = "plug Item Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "socket Type Hash"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "socket Categories"
          attrs = [
            {
              name    = "socket Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "socket Indexes"
              type    = "[]int"
              display = "none"
            },
          ]
        },
        {
          name = "socket Entries"
          attrs = [
            {
              name    = "default Visible"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hide Perks In Item Tooltip"
              type    = "bool"
              display = "none"
            },
            {
              name    = "overrides Ui Appearance"
              type    = "bool"
              display = "none"
            },
            {
              name    = "plug Sources"
              type    = "int"
              display = "none"
            },
            {
              name    = "prevent Initialization On Vendor Purchase"
              type    = "bool"
              display = "none"
            },
            {
              name    = "prevent Initialization When Versioning"
              type    = "bool"
              display = "none"
            },
            {
              name    = "reusable Plug Items"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "reusable Plug Set Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "single Initial Item Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "socket Type Hash"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "sockets"
          attrs = [
            {
              name    = "detail"
              type    = "string"
              display = "default"
            },
            {
              name    = "intrinsic Sockets"
              type    = "[]IntrinsicSockets"
              display = "none"
            },
            {
              name    = "socket Categories"
              type    = "[]SocketCategories"
              display = "none"
            },
            {
              name    = "socket Entries"
              type    = "[]SocketEntries"
              display = "none"
            },
          ]
        },
        {
          name = "source Data"
          attrs = [
            {
              name    = "exclusive"
              type    = "int"
              display = "none"
            },
            {
              name    = "vendor Sources"
              type    = "[]interface{}"
              display = "none"
            },
          ]
        },
        {
          name = "stats"
          attrs = [
            {
              name    = "disable Primary Stat Display"
              type    = "bool"
              display = "default"
            },
            {
              name    = "has Displayable Stats"
              type    = "bool"
              display = "default"
            },
            {
              name    = "primary Base Stat Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "stat Group Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "stats"
              type    = "map[string]interface{}"
              display = "none"
            },
          ]
        },
        {
          name = "talent Grid"
          attrs = [
            {
              name    = "hud Damage Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "item Detail String"
              type    = "string"
              display = "default"
            },
            {
              name    = "talent Grid Hash"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    item_category = {
      name = "item Category"
      structs = [
        {
          name = "item Category Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "deprecated"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "grant Destiny Breaker Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "grant Destiny Class"
              type    = "int"
              display = "none"
            },
            {
              name    = "grant Destiny Item Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "grant Destiny Sub Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "group Category Only"
              type    = "bool"
              display = "none"
            },
            {
              name    = "grouped Category Hashes"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "is Plug"
              type    = "bool"
              display = "none"
            },
            {
              name    = "parent Category Hashes"
              type    = "[]int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "short Title"
              type    = "string"
              display = "none"
            },
            {
              name    = "visible"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    item_tier_type = {
      name = "item Tier Type"
      structs = [
        {
          name = "item Tier Type Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "infusionProcess"
              type    = "InfusionProcess"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "infusion Process"
          attrs = [
            {
              name    = "base Quality Transfer Ratio"
              type    = "float64"
              display = "none"
            },
            {
              name    = "minimum Quality Increment"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    location = {
      name = "location"
      structs = [
        {
          name = "location Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "location Releases"
              type    = "[]LocationReleases"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "vendor Hash"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "location Releases"
          attrs = [
            {
              name    = "activity Bubble Name"
              type    = "int64"
              display = "none"
            },
            {
              name    = "activity Graph Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "activity Graph Node Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "activity Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "activity Path Bundle"
              type    = "int"
              display = "none"
            },
            {
              name    = "activity Path Destination"
              type    = "int"
              display = "none"
            },
            {
              name    = "destination Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "nav Point Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "spawn Point"
              type    = "int"
              display = "none"
            },
            {
              name    = "world Position"
              type    = "[]int"
              display = "none"
            },
          ]
        }
      ]
    }

    lore = {
      name = "lore"
      structs = [
        {
          name = "lore Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "subtitle"
              type    = "string"
              display = "none"
            },
          ]
        }
      ]
    }

    material_requirement_set = {
      name = "material Requirement Set"
      structs = [
        {
          name = "material Requirement Set Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "materials"
              type    = "[]Materials"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "materials"
          attrs = [
            {
              name    = "count"
              type    = "int"
              display = "none"
            },
            {
              name    = "delete On Action"
              type    = "bool"
              display = "none"
            },
            {
              name    = "item Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "omit From Requirements"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    medal_tier = {
      name = "medal Tier"
      structs = [
        {
          name = "medal Tier Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "order"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "tierName"
              type    = "string"
              display = "none"
            },
          ]
        }
      ]
    }

    milestone = {
      name = "milestone"
      structs = [
        {
          name = "milestone Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "default Order"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "explore Prioritizes Activity Image"
              type    = "bool"
              display = "none"
            },
            {
              name    = "has Predictable Dates"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "is In Game Milestone"
              type    = "bool"
              display = "none"
            },
            {
              name    = "milestone Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "quests"
              type    = "map[string]Quest"
              display = "none"
            },
            {
              name    = "recruitable"
              type    = "bool"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "show In Explorer"
              type    = "bool"
              display = "none"
            },
            {
              name    = "show In Milestones"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "milestone Items"
          attrs = [
            {
              name    = "item Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "quantity"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "quest Rewards"
          attrs = [
            {
              name    = "items"
              type    = "[]MilestoneItems"
              display = "none"
            },
          ]
        },
        {
          name = "quest"
          attrs = [
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "quest Item Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "quest Rewards"
              type    = "QuestRewards"
              display = "none"
            },
            {
              name    = "tracking Unlock Value Hash"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    node_step_summary = {
      name = "node Step Summary"
      structs = [
        {
          name  = "node Step Summary Definition"
          attrs = []
        }
      ]
    }

    objective = {
      name = "objective"
      structs = [
        {
          name = "objective Definition"
          attrs = [
            {
              name    = "allow Negative Value"
              type    = "bool"
              display = "none"
            },
            {
              name    = "allow Overcompletion"
              type    = "bool"
              display = "none"
            },
            {
              name    = "allow Value Change When Completed"
              type    = "bool"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "completed Value Style"
              type    = "int"
              display = "none"
            },
            {
              name    = "completion Value"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "in Progress Value Style"
              type    = "int"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "is Counting Downward"
              type    = "bool"
              display = "none"
            },
            {
              name    = "is Display Only Objective"
              type    = "bool"
              display = "none"
            },
            {
              name    = "location Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "minimum Visibility Threshold"
              type    = "int"
              display = "none"
            },
            {
              name    = "progress Description"
              type    = "string"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "scope"
              type    = "int"
              display = "none"
            },
            {
              name    = "show Value On Complete"
              type    = "bool"
              display = "none"
            },
            {
              name    = "unlock Value Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "value Style"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    place = {
      name = "place"
      structs = [
        {
          name = "place Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    platform_bucket_mapping = {
      name = "platform Bucket Mapping"
      structs = [
        {
          name = "platform Bucket Mapping Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "bucket Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "membership Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    plug_set = {
      name = "plug Set"
      structs = [
        {
          name = "plug Set Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "reusable Plug Items"
              type    = "[]ReusablePlugItems"
              display = "none"
            },
          ]
        },
        {
          name = "reusable Plug Items"
          attrs = [
            {
              name    = "alternate Weight"
              type    = "float64"
              display = "none"
            },
            {
              name    = "plug Item Hash"
              type    = "float64"
              display = "none"
            },
            {
              name    = "weight"
              type    = "float64"
              display = "none"
            },
          ]
        }
      ]
    }

    presentation_node = {
      name = "presentation Node"
      structs = [
        {
          name = "presentation Node Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "children"
              type    = "Children"
              display = "none"
            },
            {
              name    = "disable Child Subscreen Navigation"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "display Style"
              type    = "int"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "node Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "objective Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "original Icon"
              type    = "string"
              display = "none"
            },
            {
              name    = "parent Node Hashes"
              type    = "[]int64"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "requirements"
              type    = "PresentationRequirements"
              display = "none"
            },
            {
              name    = "root View Icon"
              type    = "string"
              display = "none"
            },
            {
              name    = "scope"
              type    = "int"
              display = "none"
            },
            {
              name    = "screen Style"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "presentation Nodes"
          attrs = [
            {
              name    = "presentation Node Hash"
              type    = "int64"
              display = "none"
            },
          ]
        },
        {
          name = "children"
          attrs = [
            {
              name    = "collectibles"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "presentation Nodes"
              type    = "[]PresentationNodes"
              display = "none"
            },
            {
              name    = "records"
              type    = "[]interface{}"
              display = "none"
            },
          ]
        },
        {
          name = "presentation Requirements"
          attrs = [
            {
              name    = "entitlement Unavailable Message"
              type    = "string"
              display = "none"
            },
          ]
        }
      ]
    }

    progression = {
      name = "progression"
      structs = [
        {
          name = "progression Definition"
          attrs = [
            // FIXME
            //            {
            //              name = "world Display Properties"
            //              type = "DisplayProperties"
            //  json:"BungieNet.Engine.Contract.Destiny.World.Definitions.IDestinyDisplayDefinition.displayProperties,omitempty"
            //
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "current Reset Count Unlock Value Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "progress To Next Step Scaling"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "repeat Last Step"
              type    = "bool"
              display = "none"
            },
            {
              name    = "reward Items"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "scope"
              type    = "int"
              display = "none"
            },
            {
              name    = "steps"
              type    = "[]Steps"
              display = "none"
            },
            {
              name    = "storage Mapping Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "visible"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "steps"
          attrs = [
            {
              name    = "display Effect Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "progress Total"
              type    = "int"
              display = "none"
            },
            {
              name    = "step Name"
              type    = "string"
              display = "none"
            },
          ]
        }
      ]
    }

    progression_level_requirement = {
      name = "progression Level Requirement"
      structs = [
        {
          name = "progression Level Requirement Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "progression Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "requirement Curve"
              type    = "[]RequirementCurve"
              display = "none"
            },
          ]
        },
        {
          name = "requirement Curve"
          attrs = [
            {
              name    = "value"
              type    = "float64"
              display = "none"
            },
            {
              name    = "weight"
              type    = "float64"
              display = "none"
            },
          ]
        }
      ]
    }

    progression_mapping = {
      name = "progression Mapping"
      structs = [
        {
          name = "progression Mapping Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    race = {
      name = "race"
      structs = [
        {
          name = "race Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "gendered Race Names"
              type    = "GenderedRaceNames"
              display = "none"
            },
            {
              name    = "gendered Race Names By Gender Hash"
              type    = "map[string]string"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "race Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "gendered Race Names"
          attrs = [
            {
              name    = "Female"
              type    = "string"
              display = "none"
            },
            {
              name    = "Male"
              type    = "string"
              display = "none"
            },
          ]
        }
      ]
    }

    record = {
      name = "record"
      structs = [
        {
          name = "record Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "completion Info"
              type    = "CompletionInfo"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "expiration Info"
              type    = "ExpirationInfo"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "interval Info"
              type    = "IntervalInfo"
              display = "none"
            },
            {
              name    = "objective Hashes"
              type    = "[]int64"
              display = "none"
            },
            {
              name    = "presentation Info"
              type    = "RecordPresentationInfo"
              display = "none"
            },
            {
              name    = "record Value Style"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "requirements"
              type    = "Requirements"
              display = "none"
            },
            {
              name    = "reward Items"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "scope"
              type    = "int"
              display = "none"
            },
            {
              name    = "state Info"
              type    = "RecordStateInfo"
              display = "none"
            },
            {
              name    = "title Info"
              type    = "TitleInfo"
              display = "none"
            },
          ]
        },
        {
          name = "completion Info"
          attrs = [
            {
              name    = "Score Value"
              type    = "int"
              display = "none"
            },
            {
              name    = "partial Completion Objective Count Threshold"
              type    = "int"
              display = "none"
            },
            {
              name    = "should Fire Toast"
              type    = "bool"
              display = "none"
            },
            {
              name    = "toast Style"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "expiration Info"
          attrs = [
            {
              name    = "description"
              type    = "string"
              display = "none"
            },
            {
              name    = "has Expiration"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "interval Info"
          attrs = [
            {
              name    = "interval Objectives"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "is Interval Versioned From Normal Record"
              type    = "bool"
              display = "none"
            },
            {
              name    = "original Objective Array Insertion Index"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "record Presentation Info"
          attrs = [
            {
              name    = "display Style"
              type    = "int"
              display = "none"
            },
            {
              name    = "parent Presentation Node Hashes"
              type    = "[]int64"
              display = "none"
            },
            {
              name    = "presentation Node Type"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "requirements"
          attrs = [
            {
              name    = "entitlement Unavailable Message"
              type    = "string"
              display = "none"
            },
          ]
        },
        {
          name = "record State Info"
          attrs = [
            {
              name    = "claimed Unlock Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "complete Unlock Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "featured Priority"
              type    = "int64"
              display = "none"
            },
          ]
        },
        {
          name = "title Info"
          attrs = [
            {
              name    = "has Title"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    report_reason_category = {
      name = "report Reason Category"
      structs = [
        {
          name = "report Reason Category Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "reasons"
              type    = "map[string]Reason"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "reason"
          attrs = [
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "reasonHash"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    reward_adjuster_pointer = {
      name = "reward Adjuster Pointer"
      structs = [
        {
          name = "reward Adjuster Pointer Definition"
          attrs = [
            {
              name    = "adjuster Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    reward_adjuster_progression_map = {
      name = "reward Adjuster Progression Map"
      structs = [
        {
          name = "reward Adjuster Progression Map Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "is Additive"
              type    = "bool"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    reward_item_list = {
      name = "reward Item List"
      structs = [
        {
          name = "reward Item List Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    reward_mapping = {
      name = "reward Mapping"
      structs = [
        {
          name = "reward Mapping Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "is Global"
              type    = "bool"
              display = "none"
            },
            {
              name    = "mapping Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "mapping Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    reward_sheet = {
      name = "reward Sheet"
      structs = [
        {
          name = "reward Sheet Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "sheet Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "sheet Index"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    reward_source = {
      name = "reward Source"
      structs = [
        {
          name  = "reward Source Definition"
          attrs = []
        }
      ]
    }

    sack_reward_item_list = {
      name = "sack Reward Item List"
      structs = [
        {
          name = "sack Reward Item List Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    sandbox_pattern = {
      name = "sandbox Pattern"
      structs = [
        {
          name = "sandbox Pattern Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "filters"
              type    = "[]Filter"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "pattern Global Tag Id Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "pattern Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "weapon Content Group Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "weapon Translation Group Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "weapon Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "weapon Type Hash"
              type    = "int64"
              display = "none"
            },
          ]
        },
        {
          name = "filter"
          attrs = [
            {
              name    = "arrangement Index By Stat Value"
              type    = "map[string]int"
              display = "none"
            },
            {
              name    = "art Arrangement Region Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "art Arrangement Region Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "stat Hash"
              type    = "int64"
              display = "none"
            },
          ]
        }
      ]
    }

    sandbox_perk = {
      name = "sandbox Perk"
      structs = [
        {
          name = "sandbox Perk Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "damage Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "damage Type Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "is Displayable"
              type    = "bool"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    season = {
      name = "season"
      structs = [
        {
          name = "season Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "season Number"
              type    = "int"
              display = "none"
            },
            {
              name    = "season Pass Progression Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "season Pass Unlock Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "start Time In Seconds"
              type    = "string"
              display = "none"
            },
          ]
        }
      ]
    }

    season_pass = {
      name = "season Pass"
      structs = [
        {
          name = "season Pass Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "prestige Progression Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "reward Progression Hash"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    socket_category = {
      name = "socket Category"
      structs = [
        {
          name = "socket Category Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "category Style"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "ui Category Style"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    socket_type = {
      name = "socket Type"
      structs = [
        {
          name = "socket Type Definition"
          attrs = [
            {
              name    = "always Randomize Sockets"
              type    = "bool"
              display = "none"
            },
            {
              name    = "avoid Duplicates On Initialization"
              type    = "bool"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "currency Scalars"
              type    = "[]interface{} "
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "hide Duplicate Reusable Plugs"
              type    = "bool"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "insert Action"
              type    = "InsertAction"
              display = "none"
            },
            {
              name    = "is Preview Enabled"
              type    = "bool"
              display = "none"
            },
            {
              name    = "overrides Ui Appearance"
              type    = "bool"
              display = "none"
            },
            {
              name    = "plug Whitelist"
              type    = "[]PlugWhitelist"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "socket Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "visibility"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "insert Action"
          attrs = [
            {
              name    = "action Execute Seconds"
              type    = "int"
              display = "none"
            },
            {
              name    = "action Sound Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "action Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "is Positive Action"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "plug Whitelist"
          attrs = [
            {
              name    = "category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "category Identifier"
              type    = "string"
              display = "none"
            },
          ]
        }
      ]
    }

    stat = {
      name = "stat"
      structs = [
        {
          name = "stat Definition"
          attrs = [
            {
              name    = "aggregation Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "has Computed Block"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "interpolate"
              type    = "bool"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "stat Category"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    stat_group = {
      name = "stat Group"
      structs = [
        {
          name = "stat Group Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "maximum Value"
              type    = "int"
              display = "none"
            },
            {
              name    = "overrides"
              type    = "interface{}"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "scaled Stats"
              type    = "[]ScaledStats"
              display = "none"
            },
            {
              name    = "ui Position"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "display Interpolation"
          attrs = [
            {
              name    = "value"
              type    = "int"
              display = "none"
            },
            {
              name    = "weight"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "scaled Stats"
          attrs = [
            {
              name    = "display As Numeric"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Interpolation"
              type    = "[]DisplayInterpolation"
              display = "none"
            },
            {
              name    = "maximum Value"
              type    = "int"
              display = "none"
            },
            {
              name    = "stat Hash"
              type    = "int64"
              display = "none"
            },
          ]
        }
      ]
    }

    talent_grid = {
      name = "talent Grid"
      structs = [
        {
          name = "talent Grid Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "calc Max Grid Level"
              type    = "int"
              display = "none"
            },
            {
              name    = "calc Progress To Max Level"
              type    = "int"
              display = "none"
            },
            {
              name    = "exclusive Sets"
              type    = "[]ExclusiveSet"
              display = "none"
            },
            {
              name    = "grid Level Per Column"
              type    = "int"
              display = "none"
            },
            {
              name    = "groups"
              type    = "map[string]Group"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "independent Node Indexes"
              type    = "[]float64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "max Grid Level"
              type    = "int"
              display = "none"
            },
            {
              name    = "maximum Random Material Requirements"
              type    = "int"
              display = "none"
            },
            {
              name    = "node Categories"
              type    = "[]NodeCategory"
              display = "none"
            },
            {
              name    = "nodes"
              type    = "[]Node"
              display = "none"
            },
            {
              name    = "progression Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "group"
          attrs = [
            {
              name    = "group Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "node Hashes"
              type    = "[]int64"
              display = "none"
            },
            {
              name    = "opposing Group Hashes"
              type    = "[]int64"
              display = "none"
            },
            {
              name    = "opposing Node Hashes"
              type    = "[]int64"
              display = "none"
            },
          ]
        },
        {
          name  = "exclusive Set"
          attrs = []
        },
        {
          name = "node Category"
          attrs = [
            {
              name    = "identifier"
              type    = "string"
              display = "none"
            },
            {
              name    = "is Lore Driven"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "node Hashes"
              type    = "[]int64"
              display = "none"
            },
          ]
        },
        {
          name = "activation Requirement"
          attrs = [
            {
              name    = "exclusive Set Required Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "grid Level"
              type    = "int"
              display = "none"
            },
            {
              name    = "material Requirement Hashes"
              type    = "[]int64"
              display = "none"
            },
          ]
        },
        {
          name = "node"
          attrs = [
            {
              name    = "auto Unlocks"
              type    = "bool"
              display = "none"
            },
            {
              name    = "binary Pair Node Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "column"
              type    = "int"
              display = "none"
            },
            {
              name    = "exclusive Set Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "exclusive With Node Hashes"
              type    = "[]int64"
              display = "none"
            },
            {
              name    = "group Scope Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "ignore For Completion"
              type    = "bool"
              display = "none"
            },
            {
              name    = "is Random"
              type    = "bool"
              display = "none"
            },
            {
              name    = "is Random Repurchasable"
              type    = "bool"
              display = "none"
            },
            {
              name    = "is Real Step Selection Random"
              type    = "bool"
              display = "none"
            },
            {
              name    = "last Step Repeats"
              type    = "bool"
              display = "none"
            },
            {
              name    = "layout Identifier"
              type    = "string"
              display = "none"
            },
            {
              name    = "node Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "node Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "node Style Identifier"
              type    = "string"
              display = "none"
            },
            {
              name    = "original Node Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "prerequisite Node Indexes"
              type    = "[]int32"
              display = "none"
            },
            {
              name    = "random Start Progression Bar At Progression"
              type    = "int"
              display = "none"
            },
            {
              name    = "row"
              type    = "int"
              display = "none"
            },
            {
              name    = "steps"
              type    = "[]Step"
              display = "none"
            },
            {
              name    = "talent Node Types"
              type    = "int"
              display = "none"
            },
            {
              name    = "group Hash"
              type    = "int64"
              display = "none"
            },
          ]
        },
        {
          name = "step"
          attrs = [
            {
              name    = "activation Requirement"
              type    = "ActivationRequirement"
              display = "none"
            },
            {
              name    = "affects Level"
              type    = "bool"
              display = "none"
            },
            {
              name    = "affects Quality"
              type    = "bool"
              display = "none"
            },
            {
              name    = "can Activate Next Step"
              type    = "bool"
              display = "none"
            },
            {
              name    = "damage Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "interaction Description"
              type    = "string"
              display = "none"
            },
            {
              name    = "is Next Step Random"
              type    = "bool"
              display = "none"
            },
            {
              name    = "next Step Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "node Step Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "perk Hashes"
              type    = "[]int64"
              display = "none"
            },
            {
              name    = "start Progression Bar At Progress"
              type    = "int"
              display = "none"
            },
            {
              name    = "stat Hashes"
              type    = "[]int64"
              display = "none"
            },
            {
              name    = "step Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "true Property Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "true Step Index"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    unlock = {
      name = "unlock"
      structs = [
        {
          name = "unlock Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "scope"
              type    = "int"
              display = "none"
            },
            {
              name    = "unlock Type"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    unlock_count_mapping = {
      name = "unlock Count Mapping"
      structs = [
        {
          name = "unlock Count Mapping Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "unlock Value Hash"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    unlock_event = {
      name = "unlock Event"
      structs = [
        {
          name = "unlock Event Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "new Sequence Reward Site Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "sequence Last Updated Unlock Value Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "sequence Unlock Value Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "unlock Entries"
              type    = "[]UnlockEntries"
              display = "none"
            },
          ]
        },
        {
          name = "unlock Entries"
          attrs = [
            {
              name    = "cleared Value"
              type    = "float64"
              display = "none"
            },
            {
              name    = "selected Value"
              type    = "float64"
              display = "none"
            },
            {
              name    = "unlock Hash"
              type    = "float64"
              display = "none"
            },
            {
              name    = "unlock Value Hash"
              type    = "float64"
              display = "none"
            },
          ]
        }
      ]
    }

    unlock_expression_mapping = {
      name = "unlock Expression Mapping"
      structs = [
        {
          name = "unlock Expression Mapping Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }

    unlock_value = {
      name = "unlock Value"
      structs = [
        {
          name = "unlock Value Definition"
          attrs = [
            {
              name    = "aggregation Type"
              type    = "int"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "mapping Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "scope"
              type    = "int"
              display = "none"
            },
          ]
        }
      ]
    }

    vendor = {
      name = "vendor"
      structs = [
        {
          name = "vendor Definition"
          attrs = [
            // FIXME
            //            {
            //              name = ""
            //              type = "DisplayProperties"
            //            },
            // json:"BungieNet.Engine.Contract.Destiny.World.Definitions.IDestinyDisplayDefinition.displayProperties,omitempty"
            {
              name    = "accepted Items"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "actions"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "categories"
              type    = "[]Categories"
              display = "none"
            },
            {
              name    = "consolidate Categories"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Categories"
              type    = "[]DisplayCategories"
              display = "none"
            },
            {
              name    = "display Item Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "enabled"
              type    = "bool"
              display = "none"
            },
            {
              name    = "faction Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "failure Strings"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "ignore Sale Item Hashes"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "inhibit Buying"
              type    = "bool"
              display = "none"
            },
            {
              name    = "inhibit Selling"
              type    = "bool"
              display = "none"
            },
            {
              name    = "interactions"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "inventory Flyouts"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "item List"
              type    = "[]ItemList"
              display = "none"
            },
            {
              name    = "original Categories"
              type    = "[]OriginalCategories"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "reset Interval Minutes"
              type    = "int"
              display = "none"
            },
            {
              name    = "reset Offset Minutes"
              type    = "int"
              display = "none"
            },
            {
              name    = "return With Vendor Request"
              type    = "bool"
              display = "none"
            },
            {
              name    = "services"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "unlock Ranges"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "unlock Value Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "visible"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "categories"
          attrs = [
            {
              name    = "buy String Override"
              type    = "string"
              display = "none"
            },
            {
              name    = "category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "category Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "disabled Description"
              type    = "string"
              display = "none"
            },
            {
              name    = "hide From Regular Purchase"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hide If No Currency"
              type    = "bool"
              display = "none"
            },
            {
              name    = "is Display Only"
              type    = "bool"
              display = "none"
            },
            {
              name    = "is Preview"
              type    = "bool"
              display = "none"
            },
            {
              name    = "quantity Available"
              type    = "int"
              display = "none"
            },
            {
              name    = "reset Interval Minutes Override"
              type    = "int"
              display = "none"
            },
            {
              name    = "reset Offset Minutes Override"
              type    = "int"
              display = "none"
            },
            {
              name    = "show Unavailable Items"
              type    = "bool"
              display = "none"
            },
            {
              name    = "sort Value"
              type    = "int"
              display = "none"
            },
            {
              name    = "vendor Item Indexes"
              type    = "[]int"
              display = "none"
            },
          ]
        },
        {
          name = "display Categories"
          attrs = [
            {
              name    = "display Category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "display In Banner"
              type    = "bool"
              display = "none"
            },
            {
              name    = "display Properties"
              type    = "DisplayProperties"
              display = "nested"
            },
            {
              name    = "identifier"
              type    = "string"
              display = "none"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "sort Order"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "vendor Action"
          attrs = [
            {
              name    = "execute Seconds"
              type    = "float64"
              display = "none"
            },
            {
              name    = "is Positive"
              type    = "bool"
              display = "none"
            },
          ]
        },
        {
          name = "creation Levels"
          attrs = [
            {
              name    = "level"
              type    = "int"
              display = "none"
            },
          ]
        },
        {
          name = "item List"
          attrs = [
            {
              name    = "action"
              type    = "VendorAction"
              display = "none"
            },
            {
              name    = "category Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "creation Levels"
              type    = "[]CreationLevels"
              display = "none"
            },
            {
              name    = "currencies"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "display Category"
              type    = "string"
              display = "none"
            },
            {
              name    = "display Category Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "exclusivity"
              type    = "int"
              display = "none"
            },
            {
              name    = "expiration Tooltip"
              type    = "string"
              display = "none"
            },
            {
              name    = "failure Indexes"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "inventory Bucket Hash"
              type    = "int64"
              display = "none"
            },
            {
              name    = "item Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "license Unlock Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "maximum Level"
              type    = "int"
              display = "none"
            },
            {
              name    = "minimum Level"
              type    = "int"
              display = "none"
            },
            {
              name    = "original Category Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "price Override Enabled"
              type    = "bool"
              display = "none"
            },
            {
              name    = "purchasable Scope"
              type    = "int"
              display = "none"
            },
            {
              name    = "quantity"
              type    = "int"
              display = "none"
            },
            {
              name    = "redirect To Sale Indexes"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "refund Policy"
              type    = "int"
              display = "none"
            },
            {
              name    = "refund Time Limit"
              type    = "int"
              display = "none"
            },
            {
              name    = "reward Adjustor Pointer Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "seed Override"
              type    = "int"
              display = "none"
            },
            {
              name    = "socket Overrides"
              type    = "[]interface{}"
              display = "none"
            },
            {
              name    = "sort Value"
              type    = "int"
              display = "none"
            },
            {
              name    = "vendor Item Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "visibility Scope"
              type    = "int"
              display = "none"
            },
            {
              name    = "weight"
              type    = "float64"
              display = "none"
            },
          ]
        },
        {
          name = "original Categories"
          attrs = [
            {
              name    = "buy String Override"
              type    = "string"
              display = "none"
            },
            {
              name    = "category Hash"
              type    = "int"
              display = "none"
            },
            {
              name    = "category Index"
              type    = "int"
              display = "none"
            },
            {
              name    = "disabled Description"
              type    = "string"
              display = "none"
            },
            {
              name    = "hide From Regular Purchase"
              type    = "bool"
              display = "none"
            },
            {
              name    = "hide If No Currency"
              type    = "bool"
              display = "none"
            },
            {
              name    = "is Display Only"
              type    = "bool"
              display = "none"
            },
            {
              name    = "is Preview"
              type    = "bool"
              display = "none"
            },
            {
              name    = "quantity Available"
              type    = "int"
              display = "none"
            },
            {
              name    = "reset Interval Minutes Override"
              type    = "int"
              display = "none"
            },
            {
              name    = "reset Offset Minutes Override"
              type    = "int"
              display = "none"
            },
            {
              name    = "show Unavailable Items"
              type    = "bool"
              display = "none"
            },
            {
              name    = "sort Value"
              type    = "int"
              display = "none"
            },
            {
              name    = "vendor Item Indexes"
              type    = "[]int"
              display = "none"
            },
          ]
        }
      ]
    }

    vendor_group = {
      name = "vendor Group"
      structs = [
        {
          name = "vendor Group Definition"
          attrs = [
            {
              name    = "blacklisted"
              type    = "bool"
              display = "none"
            },
            {
              name    = "category Name"
              type    = "string"
              display = "none"
            },
            {
              name    = "hash"
              type    = "int"
              display = "default"
            },
            {
              name    = "index"
              type    = "int"
              display = "none"
            },
            {
              name    = "order"
              type    = "int"
              display = "none"
            },
            {
              name    = "redacted"
              type    = "bool"
              display = "none"
            },
          ]
        }
      ]
    }
  }

  filename = "${var.target}/${replace(lower(each.value.name), " ", "_")}_definition.tf"
  content = templatefile("${path.module}/main.tf.tpl", {
    moduleName = replace(lower(each.value.name), " ", "_")
    name       = each.value.name
    structs    = each.value.structs
  })
}