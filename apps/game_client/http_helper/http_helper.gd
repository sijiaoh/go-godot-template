class_name HTTPHelper extends Node


var base_url: String


func build_url(path: String) -> String:
	return "%s/%s" % [base_url.trim_suffix("/"), path.trim_prefix("/")]

func request_get(path: String) -> HTTPResponse:
	return await _request(HTTPClient.METHOD_GET, build_url(path))

func request_post(path: String, request_body: Dictionary = {}) -> HTTPResponse:
	return await _request(HTTPClient.METHOD_POST, build_url(path), request_body)

func request_put(path: String, request_body: Dictionary = {}) -> HTTPResponse:
	return await _request(HTTPClient.METHOD_PUT, build_url(path), request_body)

func request_delete(path: String) -> HTTPResponse:
	return await _request(HTTPClient.METHOD_DELETE, build_url(path))

func _request(method: int, url: String, request_body: Dictionary = {}) -> HTTPResponse:
	var http_response := HTTPResponse.new()

	var http_request := HTTPRequest.new()
	add_child(http_request)

	var header := ["Content-Type: application/json", "Accept: application/json"]
	var request_body_json := JSON.stringify(request_body)

	var err := http_request.request(url, header, method, request_body_json)
	if err != OK:
		http_request.queue_free()
		push_error("HTTP request error: %s" % err)
		http_response.err = err
		return http_response

	var result = await http_request.request_completed
	http_request.queue_free()

	http_response.err = result[0]
	if http_response.err != OK:
		push_error("HTTP request error: %s" % http_response.err)
		return http_response

	http_response.status_code = result[1]
	var response_body_json: String = result[3]
	http_response.body = JSON.parse_string(response_body_json)
	return http_response
