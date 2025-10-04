class_name PlayerState extends Node


static var _instance: PlayerState = null


static func instance(ctx: Node) -> PlayerState:
	_instance = ctx.get_tree().root.find_child("PlayerState", true, false)
	assert(_instance != null)
	return _instance


var _me: Me


func fetch_me() -> Me:
	if _me == null:
		_me = await Me.fetch()
	return _me
