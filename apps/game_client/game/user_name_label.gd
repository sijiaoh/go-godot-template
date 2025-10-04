extends Label


func _ready() -> void:
	var me := await PlayerState.instance(self).fetch_me()
	if me == null: return
	text = me.name
