class_name TransferCodeSchema


class GetTransferCodeResponse extends HTTPResponse:
	func code() -> String:
		return body["code"]
