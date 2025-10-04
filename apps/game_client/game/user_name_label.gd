extends Label


func _ready() -> void:
	var me := await PlayerState.instance(self).fetch_me()
	text = me.name
