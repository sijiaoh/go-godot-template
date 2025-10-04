class_name Screen extends Node2D


signal exit_screen


@export var screen_entry: ScreenEntry


# 重写此函数开始退出动画
func _on_exit_screen() -> void:
	ScreenEvents.exited_screen.emit(screen_entry)
	queue_free()


func _ready() -> void:
	assert(screen_entry != null)

	exit_screen.connect(_on_exit_screen)
