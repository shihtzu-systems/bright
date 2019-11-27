// Code generated; DANGER ZONE FOR EDITS

module ${replace(lower(name)," ","_")}_definition {
  source = "./definition"

  target = var.target

  name = "${name}"
  structs_content = join("\n", [
%{ for struct in structs ~}
    module.${replace(lower(name)," ","_")}_${replace(lower(struct.name)," ","_")}.rendered,
%{ endfor ~}
  ])

}

%{ for struct in structs ~}
module ${replace(lower(name)," ","_")}_${replace(lower(struct.name)," ","_")} {
  source = "./struct"

  name = "${struct.name}"
  attributes_content = join("\n", [
%{ for attr in struct.attrs ~}
    module.${replace(lower(name)," ","_")}_${replace(lower(struct.name)," ","_")}_${replace(lower(attr.name)," ","_")}.rendered,
%{ endfor ~}
  ])
}

%{ for attr in struct.attrs ~}
module ${replace(lower(name)," ","_")}_${replace(lower(struct.name)," ","_")}_${replace(lower(attr.name)," ","_")} {
  source = "./attribute"

  name = "${attr.name}"
  type = "${attr.type}"
}
%{ endfor ~}


%{ endfor ~}