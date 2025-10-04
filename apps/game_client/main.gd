extends Node2D


func _ready() -> void:
	AuthenticationEvents.unauthorized.connect(_on_authentication_unauthorized)

	if not Authentication.is_logged_in:
		ScreenEvents.replace_screen.emit(ScreenDatabase.TITLE_SCREEN_ENTRY)
	else:
		ScreenEvents.replace_screen.emit(ScreenDatabase.GAME_SCREEN_ENTRY)


func _on_authentication_unauthorized() -> void:
		ScreenEvents.replace_screen.emit(ScreenDatabase.TITLE_SCREEN_ENTRY)
