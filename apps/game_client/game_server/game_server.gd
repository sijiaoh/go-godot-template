extends HTTPHelper


func _ready() -> void:
	base_url = "http://localhost:3000"

func signup(params: SignupAPI.SignupParams) -> SignupAPI.SignupResponse:
	var response := SignupAPI.SignupResponse.new()
	await request_post("/signup", params, response)
	return response
