const getNow = Date.now || function() {
    return new Date().getTime()
}

export function throttle(func, wait, options) {
    var args = null, context = null, result = null, timeout = null
    var previous = 0
    if (!options) {
        options = {};
    }
    var later = () => {
        previous = options.leading === false ? 0 : getNow()
        timeout  = null
        result   = func.apply(context, args)
        if (!timeout) {
            return context = args = null
        }
    }
    return function() {
        var now, remaining
        now = getNow()
        if (!previous && options.leading === false) {
            previous = now
        }
        remaining = wait - (now - previous)
        context = this
        args = arguments
        if (remaining <= 0 || remaining > wait) {
            clearTimeout(timeout)
            timeout = null
            previous = now
            result = func.apply(context, args)
            if (!timeout) {
                context = args = null;
            }
        } else if (!timeout && options.trailing !== false) {
            timeout = setTimeout(later, remaining)
        }
        return result
    }
}