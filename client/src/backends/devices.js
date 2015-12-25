
export function getDevices() {
    return fetch("/v1/devices").then((response) => {
        if (response.status >= 400) {
            throw new Error("Bad response from server")
        }
        return response.json()
    })
}