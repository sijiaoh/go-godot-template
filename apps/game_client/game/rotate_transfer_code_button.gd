extends Button


func _ready() -> void:
	pressed.connect(func() -> void:
		var tc := await PlayerState.instance(self).fetch_transfer_code()
		tc.rotate()
	)
