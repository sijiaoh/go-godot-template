class_name AuthenticationSchema


class SignupParams extends HTTPParams:
	func _init(user_name: String):
		body["userName"] = user_name


class SignupResponse extends HTTPResponse:
	func token() -> String:
		return body["token"]


class LoginParams extends HTTPParams:
	func _init(transfer_code: String):
		body["transferCode"] = transfer_code


class LoginResponse extends HTTPResponse:
	func token() -> String:
		return body["token"]
