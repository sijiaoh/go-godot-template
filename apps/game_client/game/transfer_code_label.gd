extends Label


func _ready() -> void:
	var tc := await PlayerState.instance(self).fetch_transfer_code()
	if tc == null: return

	tc.rotated.connect(func() -> void:
		text = tc.code
	)
	text = tc.code
