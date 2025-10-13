extends Button


func _ready():
	pressed.connect(
		func():
			ScreenEvents.replace_screen.emit(ScreenDatabase.TITLE_SCREEN_ENTRY)
	)
