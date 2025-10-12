class_name TransferCodeSchema


class GetTransferCodeResponse extends HTTPResponse:
	func code() -> String:
		return body["code"]

class RotateTransferCodeResponse extends HTTPResponse:
	func code() -> String:
		return body["code"]
