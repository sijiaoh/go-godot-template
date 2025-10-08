extends Screen


@export var signup_container: Container
@export var user_name_input: LineEdit


func signup():
	await Authentication.signup(user_name_input.text)
	if Authentication.is_logged_in:
		ScreenEvents.replace_screen.emit(ScreenDatabase.GAME_SCREEN_ENTRY)


func _on_covered_button_pressed() -> void:
	if Authentication.is_logged_in:
		ScreenEvents.replace_screen.emit(ScreenDatabase.GAME_SCREEN_ENTRY)
	else:
		signup_container.visible = true
