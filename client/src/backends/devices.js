
export function getDevices() {
    return fetch("/v1/devices").then((response) => {
        if (response.status >= 400) {
            throw new Error("Bad response from server")
        }
        return response.json()
    })
}

export function getDeviceDetails(deviceIp) {
    return fetch("/v1/devices/" + deviceIp).then((response) => {
        if (response.status >= 400) {
            throw new Error("Bad response from server")
        }
        return response.json()
    })
}

export function getDeviceSpan(deviceIp) {
    return fetch("/v1/devices/" + deviceIp + "/span").then((response) => {
        if (response.status >= 400) {
            throw new Error("Bad response from server")
        }
        return response.json()
    })
}