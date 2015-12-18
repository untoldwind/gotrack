
export function getRates() {
    return fetch("/v1/data/totals/rates").then((response) => {
        if (response.status >= 400) {
            throw new Error("Bad response from server")
        }
        return response.json()
    })
}