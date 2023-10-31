
function submitForm() {
    let username = document.getElementById('username');
    let password = document.getElementById('password');
    let confirm = document.getElementById('confirm');

    if (password.value !== confirm.value) {
        document.getElementById('info_box').innerHTML = '密码不一致';
        document.getElementById('alert_box').style.visibility = 'visible';
        return
    }

    let request = new XMLHttpRequest();
    request.open("POST", "register");

    let data = {
        'method': 'register',
        'username': username.value,
        'password': password.value
    }

    request.addEventListener('readystatechange', () => {
        if (request.readyState === 4 && request.status === 200) {
            const result = request.responseText;
            if (result === "OK") {
                setCookie('username', data.username, 365, '/login');
                setCookie('password', data.password, 365, '/login');
                window.location.href = '/';
            } else if (result === "OCCUPIED") {
                document.getElementById('info_box').innerHTML = '用户名已被占用';
                document.getElementById('alert_box').style.visibility = 'visible';
            }
        }
    });

    request.send(JSON.stringify(data))
}
