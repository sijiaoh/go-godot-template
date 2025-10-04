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

func me() -> UserSchema.MeResponse:
	var response := UserSchema.MeResponse.new()
	await request_get("/me", response)
	return response

func _request(method: int, url: String, params: HTTPParams, response: HTTPResponse) -> void:
	await super._request(method, url, params, response)
	if response.err == OK and response.status_code == 401:
		AuthenticationEvents.unauthorized.emit()
