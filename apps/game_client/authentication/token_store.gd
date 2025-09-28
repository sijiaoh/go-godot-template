class_name TokenStore


static func set_token(token: String) -> void:
	var file := FileAccess.open("user://token.txt", FileAccess.WRITE)
	file.store_string(token)
	file.close()

static func get_token() -> String:
	var file := FileAccess.open("user://token.txt", FileAccess.READ)
	var token := file.get_as_text().strip_edges()
	file.close()
	return token
