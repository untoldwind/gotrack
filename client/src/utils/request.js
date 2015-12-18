class RequestError extends Error {
    constructor(method, uri, status, text) {
        super('Request ' + method + ' ' + uri + ' failed with ' + status + ': ' + text, 'RequestError')
        this.method  = method
        this.status  = status
        this.text    = text
    }
}

function buildQuery(params) {
    var query = []
    for (const key in params) {
        const value = params[key]
        query.push(encodeURIComponent(key), "=", encodeURIComponent(value))
    }
    if (query.length > 0) {
        return "?" + query.join("&")
    }
    return ""
}

function createRequest() {
    if (window.ActiveXObject) {
        return new ActiveXObject('Microsoft.XMLHTTP')
    } else if (window.XMLHttpRequest) {
        return new XMLHttpRequest()
    }
    return false
}

export function jsonRequest(method, uri, params, data) {
    const request = createRequest()
    const result  = new Promise((resolve, reject) => {
        request.onreadystatechange = () => {
            if (request.readyState == 4) {
                if (request.status >= 200 && request.status < 300) {
                    resolve(request.responseText ? JSON.parse(request.responseText) : null)
                } else {
                    reject(new RequestError(method, uri, request.status, request.responseText))
                }
            }
        }
        request.open(method, uri + buildQuery(params), true)
        request.setRequestHeader('Accept', 'application/json')
        if (data) {
            request.setRequestHeader('Content-Type', 'application/json; charset=utf-8')
            request.send(JSON.stringify(data))
        } else {
            request.send(null)
        }
    })
    result.abort = () => {
        request.abort()
    }
    return result
}
