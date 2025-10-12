# 用于和游戏服务器通讯的类
#
# 重写_request实现通用的通讯失败逻辑

extends HTTPHelper


func _ready() -> void:
	base_url = "http://localhost:3000"


func load_token() -> void:
	var token := TokenStore.get_token()
	base_headers["Authorization"] = "Bearer %s" % token


func signup(params: SignupSchema.SignupParams) -> SignupSchema.SignupResponse:
	var response := SignupSchema.SignupResponse.new()
	await request_post("/signup", params, response)
	return response


func get_me() -> UserSchema.MeResponse:
	var response := UserSchema.MeResponse.new()
	await request_get("/me", response)
	return response


func get_transfer_code() -> TransferCodeSchema.GetTransferCodeResponse:
	var response := TransferCodeSchema.GetTransferCodeResponse.new()
	await request_get("/transfer-code", response)
	return response


func rotate_transfer_code() -> TransferCodeSchema.RotateTransferCodeResponse:
	var response := TransferCodeSchema.RotateTransferCodeResponse.new()
	await request_post("/transfer-code/rotate", null, response)
	return response


func _request(method: int, url: String, params: HTTPParams, response: HTTPResponse) -> void:
	await super._request(method, url, params, response)
	if response.err == OK and response.status_code == 401:
		AuthenticationEvents.unauthorized.emit()
	if response.err != OK:
		ModalEvents.open_modal.emit(
			tr("通信失败。") + " err=%s status_code=%s" % [response.err, response.status_code] + "\n" +
			tr("请检查网络连接或稍后重试。")
		)
		ScreenEvents.force_replace_root_screen.emit(ScreenDatabase.TITLE_SCREEN_ENTRY)
