// Code generated; DANGER ZONE FOR EDITS

module medal_tier_definition {
  source = "./definition"

  target = var.target

  name = "medal Tier"
  structs_content = join("\n", [
    module.medal_tier_medal_tier_definition.rendered,
  ])

}

module medal_tier_medal_tier_definition {
  source = "./struct"

  name = "medal Tier Definition"
  attributes_content = join("\n", [
    module.medal_tier_medal_tier_definition_blacklisted.rendered,
    module.medal_tier_medal_tier_definition_hash.rendered,
    module.medal_tier_medal_tier_definition_index.rendered,
    module.medal_tier_medal_tier_definition_order.rendered,
    module.medal_tier_medal_tier_definition_redacted.rendered,
    module.medal_tier_medal_tier_definition_tiername.rendered,
  ])
}

module medal_tier_medal_tier_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module medal_tier_medal_tier_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module medal_tier_medal_tier_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module medal_tier_medal_tier_definition_order {
  source = "./attribute"

  name = "order"
  type = "int"
}
module medal_tier_medal_tier_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}
module medal_tier_medal_tier_definition_tiername {
  source = "./attribute"

  name = "tierName"
  type = "string"
}


