
export function formatRate(rate) {
    var value = rate.toString()
    var unit = "bytes/s"

    if (rate >= 100 * 1024) {
        value = (rate / 1024.0 / 1024.0).toFixed(2)
        unit = "MB/s"
    } else if (rate >= 100) {
        value = (rate / 1024.0).toFixed(2)
        unit = "kB/s"
    }

    return { value: value, unit: unit }
}