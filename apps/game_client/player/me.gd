class_name Me


var name: String


static func fetch() -> Me:
	var res := await GameServer.me()
	if res.status_code != 200:
		ModalEvents.open_modal.emit("获取用户信息失败: err=%s status_code=%s" % [res.err, res.status_code])
		# TODO: 制作常用通讯失败处理，返回Title
		return null

	var me := Me.new()
	me.name = res.name()
	return me
