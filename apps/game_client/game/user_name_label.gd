extends Label


func _process(_delta: float) -> void:
	# TODO: 需要用signal优化
	var player_state := PlayerState.instance(self)
	if player_state != null && player_state.me != null:
		text = player_state.me.name
