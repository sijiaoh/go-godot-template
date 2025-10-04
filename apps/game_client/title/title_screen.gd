extends Screen


@export var user_name_input: LineEdit


func signup():
	await Authentication.signup(user_name_input.text)
	if Authentication.is_logged_in:
		ScreenEvents.replace_screen.emit(ScreenDatabase.GAME_SCREEN_ENTRY)
