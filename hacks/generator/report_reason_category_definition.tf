// Code generated; DANGER ZONE FOR EDITS

module report_reason_category_definition {
  source = "./definition"

  target = var.target

  name = "report Reason Category"
  structs_content = join("\n", [
    module.report_reason_category_report_reason_category_definition.rendered,
    module.report_reason_category_reason.rendered,
  ])

}

module report_reason_category_report_reason_category_definition {
  source = "./struct"

  name = "report Reason Category Definition"
  attributes_content = join("\n", [
    module.report_reason_category_report_reason_category_definition_blacklisted.rendered,
    module.report_reason_category_report_reason_category_definition_display_properties.rendered,
    module.report_reason_category_report_reason_category_definition_hash.rendered,
    module.report_reason_category_report_reason_category_definition_index.rendered,
    module.report_reason_category_report_reason_category_definition_reasons.rendered,
    module.report_reason_category_report_reason_category_definition_redacted.rendered,
  ])
}

module report_reason_category_report_reason_category_definition_blacklisted {
  source = "./attribute"

  name = "blacklisted"
  type = "bool"
}
module report_reason_category_report_reason_category_definition_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module report_reason_category_report_reason_category_definition_hash {
  source = "./attribute"

  name = "hash"
  type = "int"
}
module report_reason_category_report_reason_category_definition_index {
  source = "./attribute"

  name = "index"
  type = "int"
}
module report_reason_category_report_reason_category_definition_reasons {
  source = "./attribute"

  name = "reasons"
  type = "map[string]Reason"
}
module report_reason_category_report_reason_category_definition_redacted {
  source = "./attribute"

  name = "redacted"
  type = "bool"
}


module report_reason_category_reason {
  source = "./struct"

  name = "reason"
  attributes_content = join("\n", [
    module.report_reason_category_reason_display_properties.rendered,
    module.report_reason_category_reason_reasonhash.rendered,
  ])
}

module report_reason_category_reason_display_properties {
  source = "./attribute"

  name = "display Properties"
  type = "DisplayProperties"
}
module report_reason_category_reason_reasonhash {
  source = "./attribute"

  name = "reasonHash"
  type = "int"
}


