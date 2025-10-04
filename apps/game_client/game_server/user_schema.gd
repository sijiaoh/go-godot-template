class_name UserSchema


class UserResponse extends HTTPResponse:
	func name() -> String:
		return body["name"]
