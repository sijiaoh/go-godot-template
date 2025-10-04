extends Node2D


func _ready() -> void:
	if not Authentication.is_logged_in:
		ScreenEvents.replace_screen.emit(ScreenDatabase.TITLE_SCREEN_ENTRY)
	else:
		ScreenEvents.replace_screen.emit(ScreenDatabase.GAME_SCREEN_ENTRY)
