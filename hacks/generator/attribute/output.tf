output rendered {
  value = templatefile("${path.module}/attribute.go.tpl", {
    name      = replace(title(var.name), " ", "")
    type      = var.type
    json_name = replace(var.name, " ", "")
    yaml_name = replace(var.name, " ", "")
  })
}