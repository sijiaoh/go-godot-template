extends Screen


var user_name_input: LineEdit


func signup():
	var params = SignupAPI.SignupParams.new(user_name_input.text)
	var res := GameServer.signup(params)

	if res.status_code != 201:
		# TODO: 显示登录失败
		assert(false)

	TokenStore.set_token(res.token())
