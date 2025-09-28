extends HTTPHelper


func _ready() -> void:
	base_url = "http://localhost:3000"

func load_token() -> void:
	var token := TokenStore.get_token()
	base_headers["Authorization"] = "Bearer %s" % token

func signup(params: SignupAPI.SignupParams) -> SignupAPI.SignupResponse:
	var response := SignupAPI.SignupResponse.new()
	await request_post("/signup", params, response)
	return response
