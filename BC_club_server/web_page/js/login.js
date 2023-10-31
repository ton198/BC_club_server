
function submitForm() {
    let username = document.getElementById('username');
    let password = document.getElementById('password');
    let check = document.getElementById('rememberMe');

    let request = new XMLHttpRequest();
    request.open('POST', '/login');
    let data = {
        'method': 'login',
        'username': username.value,
        'password': password.value
    };
    request.addEventListener('readystatechange', () => {
        if (request.readyState === 4 && request.status === 200) {
            const result = request.responseText;
            if (result === 'RIGHT') {
                if (check.checked) {
                    setCookie('username', data.username, 365, '/');
                    setCookie('password', data.password, 365, '/');
                } else {
                    setTempCookie('username', data.username, '/');
                    setTempCookie('password', data.password, '/');
                }
                window.location.href = '/';
            } else if (result === 'WRONG') {
                let alert = document.getElementById('alert_box');
                alert.innerHTML = '<strong>用户名或密码错误</strong>'
                alert.style.visibility = 'visible';
            } else if (result === 'NOT_EXIST') {
                let alert = document.getElementById('alert_box');
                alert.innerHTML = '<strong>用户名不存在</strong>'
                alert.style.visibility = 'visible';
            } else {
                let alert = document.getElementById('alert_box');
                alert.innerHTML = '<strong>服务器发生未知错误，正在紧急修复</strong>'
                alert.style.visibility = 'visible';
            }
        }
    });
    request.send(JSON.stringify(data));
}
