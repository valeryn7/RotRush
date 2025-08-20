components {
  id: "item"
  component: "/main/item.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"cup\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/items.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.3
    y: 0.3
  }
}
embedded_components {
  id: "debug"
  type: "sprite"
  data: "default_animation: \"circle_debug\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/items.atlas\"\n"
  "}\n"
  ""
  position {
    z: 1.0
  }
}
embedded_components {
  id: "fx"
  type: "sprite"
  data: "default_animation: \"explosion\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/items.atlas\"\n"
  "}\n"
  ""
  position {
    z: 1.0
  }
}
