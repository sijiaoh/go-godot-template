extends Node2D


func _ready() -> void:
	ScreenEvents.replace_screen.emit(ScreenDatabase.TITLE_SCREEN_ENTRY)
