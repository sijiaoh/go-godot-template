extends Screen


@export var user_name_input: LineEdit


func signup():
	Authentication.signup(user_name_input.text)
