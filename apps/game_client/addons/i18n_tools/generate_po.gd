# From: https://github.com/godotengine/godot-proposals/issues/10986#issuecomment-2419914451
#
# NOTE: 以下PR合并后，可以更简单的从CLI生成POT
# https://github.com/godotengine/godot/pull/98422

@tool
extends EditorPlugin


const Config = preload("res://addons/i18n_tools/config.gd")

var _button: Button


func _enter_tree() -> void:
	_button = Button.new()
	_button.text = "Generate PO"
	_button.flat = true
	_button.pressed.connect(_run)
	add_control_to_container(EditorPlugin.CONTAINER_CANVAS_EDITOR_MENU, _button)


func _exit_tree() -> void:
	remove_control_from_container(EditorPlugin.CONTAINER_CANVAS_EDITOR_MENU, _button)
	_button.queue_free()
	_button = null


func _run() -> void:
	var generate_properties_pot = preload("res://addons/i18n_tools/generate_properties_pot.gd").new()
	add_translation_parser_plugin(generate_properties_pot)

	# 自动添加额外文件
	var original_pot_files = _add_extra_files()

	var localization = EditorInterface.get_base_control().find_child("*Localization*", true, false)
	var file_dialog: EditorFileDialog = localization.get_child(5)
	file_dialog.file_selected.emit("res://locale/generated.pot")

	# 恢复原始设置
	ProjectSettings.set_setting("internationalization/locale/translations_pot_files", original_pot_files)

	remove_translation_parser_plugin(generate_properties_pot)

	var pot_path = ProjectSettings.globalize_path("res://locale/generated.pot")
	var po_dir = ProjectSettings.globalize_path("res://locale")

	if not DirAccess.dir_exists_absolute(po_dir):
		DirAccess.make_dir_recursive_absolute(po_dir)

	for language in Config.languages:
		var po_file = "%s/%s.po" % [po_dir, language]
		if not FileAccess.file_exists(po_file):
			var output: Array[String] = []
			OS.execute("msginit", ["--no-translator", "--locale", language, "--input", pot_path, "--output-file", po_file], output, true)
			print("\n".join(output))
		else:
			var output: Array[String] = []
			OS.execute("msgmerge", ["--update", po_file, pot_path], output, true)
			print("\n".join(output))

	_register_po_files()


func _add_extra_files() -> Array:
	var original_pot_files = ProjectSettings.get_setting("internationalization/locale/translations_pot_files", [])
	var updated_pot_files = original_pot_files.duplicate()

	for rule in Config.EXTRA_SCAN_RULES:
		var files = _scan_directory(rule.dir, rule.suffix)
		for file_path in files:
			if not file_path in updated_pot_files:
				updated_pot_files.append(file_path)

	ProjectSettings.set_setting("internationalization/locale/translations_pot_files", updated_pot_files)
	return original_pot_files


func _scan_directory(dir_path: String, suffix: String) -> Array[String]:
	var result: Array[String] = []
	var dir = DirAccess.open(dir_path)
	assert(dir, "Failed to open directory: %s" % dir_path)
	_scan_directory_recursive(dir, dir_path, suffix, result)
	return result


func _scan_directory_recursive(dir: DirAccess, current_path: String, suffix: String, result: Array[String]) -> void:
	dir.list_dir_begin()

	var file_name := dir.get_next()
	while file_name != "":
		var full_path := current_path
		if not full_path.ends_with("/"):
			full_path += "/"
		full_path += file_name

		if dir.current_is_dir() and not file_name.begins_with("."):
			var sub_dir = DirAccess.open(full_path)
			if sub_dir:
				_scan_directory_recursive(sub_dir, full_path, suffix, result)
		elif file_name.ends_with(suffix):
			result.append(full_path)

		file_name = dir.get_next()

	dir.list_dir_end()


func _register_po_files() -> void:
	var po_paths := PackedStringArray()
	for language in Config.languages:
		var po_path = "res://locale/%s.po" % language
		po_paths.append(po_path)
	ProjectSettings.set_setting("internationalization/locale/translations", po_paths)
	ProjectSettings.save()
