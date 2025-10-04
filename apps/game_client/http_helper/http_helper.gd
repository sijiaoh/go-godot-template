class_name HTTPHelper extends Node


var base_url: String
var base_headers: Dictionary = {}


func build_url(path: String) -> String:
	return "%s/%s" % [base_url.trim_suffix("/"), path.trim_prefix("/")]

func request_get(path: String, response: HTTPResponse) -> void:
	await _request(HTTPClient.METHOD_GET, build_url(path), null, response)

func request_post(path: String, params: HTTPParams, response: HTTPResponse) -> void:
	await _request(HTTPClient.METHOD_POST, build_url(path), params, response)

func request_put(path: String, params: HTTPParams, response: HTTPResponse) -> void:
	await _request(HTTPClient.METHOD_PUT, build_url(path), params, response)

func request_delete(path: String, response: HTTPResponse) -> void:
	await _request(HTTPClient.METHOD_DELETE, build_url(path), null, response)

func _request(method: int, url: String, params: HTTPParams, response: HTTPResponse) -> void:
	var http_request := HTTPRequest.new()
	add_child(http_request)

	var headers := ["Content-Type: application/json", "Accept: application/json"]
	for key in base_headers.keys():
		headers.append("%s: %s" % [key, base_headers[key]])

	var request_data := ""
	if params != null:
		request_data = JSON.stringify(params.body)

	var err := http_request.request(url, headers, method, request_data)
	if err != OK:
		http_request.queue_free()
		# TODO: 更好的Message
		push_error("HTTP request error: %s" % err)
		response.err = err
		return

	var result = await http_request.request_completed
	http_request.queue_free()

	response.err = result[0]
	if response.err != OK:
		# TODO: 更好的Message
		push_error("HTTP request error: %s" % response.err)
		return

	response.status_code = result[1]
	var response_body_json: String = result[3].get_string_from_utf8()
	response.body = JSON.parse_string(response_body_json)
