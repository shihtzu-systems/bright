
resource local_file this {
  filename = "${var.target}/${replace(lower(var.name), " ", "_")}_definition.go"
  content = templatefile("${path.module}/definition.go.tpl", {
    prefix         = replace(title(var.name), " ", "")
    name           = replace(lower(var.name), " ", "-")
    structsContent = var.structs_content
  })
}