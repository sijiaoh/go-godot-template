class_name UserSchema


class MeResponse extends HTTPResponse:
	func name() -> String:
		return body["name"]
