class_name TransferCode


signal rotated

var code: String


static func fetch() -> TransferCode:
	var res := await GameServer.get_transfer_code()
	if res.status_code != 200:
		return null

	var tc := TransferCode.new()
	tc.code = res.code()
	return tc

func rotate() -> void:
	var res := await GameServer.rotate_transfer_code()
	if res.status_code != 200:
		return

	code = res.code()
	rotated.emit()
