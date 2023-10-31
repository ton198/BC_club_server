
function setCookie(name, value, time, path) {
    const d = new Date()
    d.setTime(d.getTime() + (time * 24 * 60 * 60 * 1000))
    const expires = 'expires=' + d.toUTCString()
    document.cookie = name + '=' + value + '; ' + expires + ';path=' + path
}

function setTempCookie(name, value, path) {
    document.cookie = name + '=' + value + ';path=' + path
}

function getCookie(name) {
    const cname = name + '='
    const ca = document.cookie.split(';')
    for (const element of ca) {
        const c = element.trim()
        if (c.indexOf(cname) === 0) {
            return c.substring(cname.length, c.length)
        }
    }
    return ''
}

function deleteCookie(name, path) {
    setCookie(name, '', -1, path)
}