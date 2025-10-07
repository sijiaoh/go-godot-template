extends Node


var is_logged_in: bool = false


func _ready() -> void:
	var token := TokenStore.get_token()
	is_logged_in = token != ""
	if is_logged_in:
		GameServer.load_token()
	AuthenticationEvents.unauthorized.connect(_on_unauthorized)


func signup(user_name: String) -> void:
	var params = SignupSchema.SignupParams.new(user_name)
	var res := await GameServer.signup(params)

	if res.status_code != 201:
		ModalEvents.open_modal.emit(tr("登录失败") + ": err=%s status_code=%s" % [res.err, res.status_code])
		return

	TokenStore.set_token(res.token())
	GameServer.load_token()
	is_logged_in = true


func _on_unauthorized() -> void:
	TokenStore.clear_token()
	is_logged_in = false
	ModalEvents.open_modal.emit(tr("认证失败，请重新登录。"))
	ScreenEvents.force_replace_root_screen.emit(ScreenDatabase.TITLE_SCREEN_ENTRY)
