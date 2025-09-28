extends Node


var is_logged_in: bool = false


func _ready() -> void:
	var token := TokenStore.get_token()
	is_logged_in = token != ""
	if is_logged_in:
		GameServer.load_token()

func signup(user_name: String) -> void:
	var params = SignupAPI.SignupParams.new(user_name)
	var res := await GameServer.signup(params)

	if res.status_code != 201:
		# TODO: 显示登录失败
		assert(false)
		return

	TokenStore.set_token(res.token())
	GameServer.load_token()
	is_logged_in = true
