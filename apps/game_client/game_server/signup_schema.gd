class_name SignupSchema


class SignupParams extends HTTPParams:
	func _init(user_name: String):
		body["userName"] = user_name


class SignupResponse extends HTTPResponse:
	func token() -> String:
		return body["token"]
