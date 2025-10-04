class_name ScreenManager extends Node


@export var _screens_node: Node

# 除了最新的场景之外都会暂停
# 播放对话时会吧null放进来
var _screen_queue: Array[ScreenQueueItem] = []


func _ready() -> void:
	assert(_screens_node != null)

	ScreenEvents.replace_screen.connect(_replace_screen)
	ScreenEvents.append_screen.connect(_append_screen)
	ScreenEvents.exit_screen.connect(_exit_screen)
	ScreenEvents.exited_screen.connect(_on_exited_screen)


func _replace_screen(screen_entry: ScreenEntry) -> void:
	if _screen_queue.size() == 0:
		_append_screen(screen_entry)
		return

	ScreenEvents.call_deferred("emit_signal", "exit_screen")
	await ScreenEvents.exited_screen

	_append_screen(screen_entry)


func _append_screen(screen_entry: ScreenEntry) -> void:
	if _screen_queue.size() > 0:
		_screen_queue[_screen_queue.size() - 1].screen.process_mode = Node.PROCESS_MODE_DISABLED

	var screen := load(screen_entry.path).instantiate() as Screen
	_screens_node.add_child(screen)

	_screen_queue.append(ScreenQueueItem.new(screen_entry, screen))


func _exit_screen() -> void:
	assert(_screen_queue.size() > 0, "No screen to exit")
	_screen_queue[_screen_queue.size() - 1].screen.exit_screen.emit()

func _on_exited_screen(screen_entry: ScreenEntry) -> void:
	assert(_screen_queue.size() > 0, "No screen to remove")
	assert(_screen_queue[_screen_queue.size() - 1].screen_entry == screen_entry, "The exited screen is not the top screen")

	_screen_queue.pop_back()
	if _screen_queue.size() > 0:
		_screen_queue[_screen_queue.size() - 1].screen.process_mode = Node.PROCESS_MODE_INHERIT
