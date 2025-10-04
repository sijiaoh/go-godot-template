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
