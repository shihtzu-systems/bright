output rendered {
  value = templatefile("${path.module}/struct.go.tpl", {
    name         = replace(title(var.name), " ", "")
    attrsContent = var.attributes_content
  })
}