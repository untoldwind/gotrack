
export function formatRate(rate) {
    var value = rate.toString()
    var unit = "bytes/s"

    if (rate >= 100 * 1024) {
        value = (rate / 1024.0 / 1024.0).toFixed(2)
        unit = "MiB/s"
    } else if (rate >= 100) {
        value = (rate / 1024.0).toFixed(2)
        unit = "kiB/s"
    }

    return { value: value, unit: unit }
}

export function formatTotal(bytes) {
    var value = bytes.toString()
    var unit = "B"

    if (bytes >= 100 * 1024) {
        value = (bytes / 1024.0 / 1024.0).toFixed(2)
        unit = "MiB"
    } else if (bytes >= 100) {
        value = (bytes / 1024.0).toFixed(2)
        unit = "kiB"
    }

    return { value: value, unit: unit }
}

export function formatTotalString(bytes) {
    const {value, unit} = formatTotal(bytes)

    return value + " " + unit
}