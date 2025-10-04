extends Screen

func _ready() -> void:
	super._ready()
	if not Authentication.is_logged_in:
		push_error("User is not logged in, redirecting to title screen.")
		ScreenEvents.replace_screen.emit(ScreenDatabase.TITLE_SCREEN_ENTRY)
