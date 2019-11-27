// Code generated; DANGER ZONE FOR EDITS

module entitlement_offer_definition {
  source = "./definition"

  target = var.target

  name = "entitlement Offer"
  structs_content = join("\n", [
    module.entitlement_offer_entitlement_offer_definition.rendered,
  ])

}

module entitlement_offer_entitlement_offer_definition {
  source = "./struct"

  name = "entitlement Offer Definition"
  attributes_content = join("\n", [
    module.entitlement_offer_entitlement_offer_definition_blacklisted.rendered,
    module.entitlement_offer_entitlement_offer_definition_hash.rendered,
    module.entitlement_offer_entitlement_offer_definition_index.rendered,
    module.entitlement_offer_entitlement_offer_definition_offer_key.rendered,
    module.entitlement_offer_entitlement_offer_definition_redacted.rendered,
  ])
}

module entitlement_offer_entitlement_offer_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module entitlement_offer_entitlement_offer_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module entitlement_offer_entitlement_offer_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module entitlement_offer_entitlement_offer_definition_offer_key {
  source = "./attribute"

  name = "offer Key"
  type = "string"
}
module entitlement_offer_entitlement_offer_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


