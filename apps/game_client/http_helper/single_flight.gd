# 确保同一时间只有一个相同key的任务在运行，避免重复请求
extends Node


signal completed(key: String, result: Variant)

var flights: Dictionary[String, Callable] = {}


func do(key: String, func_to_run: Callable) -> Variant:
	if flights.has(key):
		# [1]: key, [2]: func_to_run result
		var signal_result: Array = ["", null]
		while signal_result[0] != key:
			signal_result = await completed
		return signal_result[1]

	flights[key] = func_to_run
	var result: Variant = await func_to_run.call()
	flights.erase(key)
	return result
