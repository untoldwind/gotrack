
export function getRates() {
    return fetch("/v1/totals/rates").then((response) => {
        if (response.status >= 400) {
            throw new Error("Bad response from server")
        }
        return response.json()
    })
}

export function getSpan() {
    return fetch("/v1/totals/span").then((response) => {
        if (response.status >= 400) {
            throw new Error("Bad response from server")
        }
        return response.json()
    })
}