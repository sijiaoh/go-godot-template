@tool
extends EditorTranslationParserPlugin


const Config = preload("res://addons/i18n_tools/config.gd")


func _parse_file(path) -> Array[PackedStringArray]:
	if path.ends_with(".tscn"):
		return _parse_scene_file(path)
	elif path.ends_with(".tres"):
		return _parse_resource_file(path)
	return []


func _get_recognized_extensions():
	return ["tscn", "tres"]


func _parse_scene_file(path: String) -> Array[PackedStringArray]:
	var resource: PackedScene = ResourceLoader.load(path)
	var root: Node = resource.instantiate()
	return _scan_node(root)


func _parse_resource_file(path: String) -> Array[PackedStringArray]:
	var result: Array[PackedStringArray] = []
	var resource: Resource = ResourceLoader.load(path)

	var target_property_names := Config.RESOURCE_PROPERTY_NAMES

	for property_name in target_property_names:
		if property_name in resource:
			var value = resource.get(property_name)
			if value is String and value.length() > 0:
				var entry: PackedStringArray = PackedStringArray()
				entry.append(value)
				result.append(entry)

	return result


func _scan_node(node: Node) -> Array[PackedStringArray]:
	var result: Array[PackedStringArray] = []

	var target_property_names := Config.NODE_PROPERTY_NAMES

	for property in node.get_property_list():
		for target_property_name in target_property_names:
			if property.name == target_property_name:
				var value: String = node.get(property.name)
				if value.length() > 0:
					var entry: PackedStringArray = PackedStringArray()
					entry.append(value)
					result.append(entry)

	for child in node.get_children():
		result += _scan_node(child)

	return result
