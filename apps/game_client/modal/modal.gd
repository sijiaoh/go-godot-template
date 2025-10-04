extends Node2D


@export var message_label: Label

@onready var canvas_layer: CanvasLayer = $CanvasLayer


func open_modal(message: String) -> void:
	message_label.text = message
	canvas_layer.visible = true

func _on_confirm_button_pressed() -> void:
	canvas_layer.visible = false
	ModalEvents.modal_closed.emit()

func _ready() -> void:
	canvas_layer.visible = false
	ModalEvents.open_modal.connect(open_modal)
