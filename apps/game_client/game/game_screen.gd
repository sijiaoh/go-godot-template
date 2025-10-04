extends Screen

func _ready() -> void:
	super._ready()
	if not Authentication.is_logged_in:
		push_error("User is not logged in, redirecting to title screen.")
		ScreenEvents.replace_screen.emit(ScreenDatabase.TITLE_SCREEN_ENTRY)

	# TODO: 任何时候接收401都返回Title
	var me := await PlayerState.instance(self).fetch_me()
	if me == null:
		ScreenEvents.replace_screen.emit(ScreenDatabase.TITLE_SCREEN_ENTRY)
