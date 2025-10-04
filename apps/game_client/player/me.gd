class_name Me


var name: String


static func fetch() -> Me:
	var res := await GameServer.me()
	if res.status_code != 200:
		return null

	var me := Me.new()
	me.name = res.name()
	return me
