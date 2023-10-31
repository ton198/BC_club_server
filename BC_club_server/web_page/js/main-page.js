
function clearAccountInfo() {
  deleteCookie('username', '/')
  deleteCookie('password', '/')
}

function loginOnClick() {
  window.location.href = "/login"
}

function logoutOnClick() {
  clearAccountInfo()
  document.getElementById('account_div').innerHTML = "  <button style=\"color: black\" onclick=\"loginOnClick()\">登录</button>"
}

function generateLoginElements(username) {
  return "  <p style=\"display: table-cell; border-right: 20px solid rgba(0,0,0,0%)\">\n" +
  "    你好，" + username + "\n" +
  "  </p>\n" +
  "  <button style=\"display: table-cell; color: black\" onclick='logoutOnClick()'>登出</button>"
}

function APOnClick() {
  window.location.href = "/AP"
}

function CompetitionOnClick() {
  window.location.href = "/competition"
}

const username = getCookie('username');
if (username !== '') {
  const password = getCookie('password');
  const loginRequest = new XMLHttpRequest()
  loginRequest.open('POST', '/login');
  const data = {
    'method': 'login',
    'username': username,
    'password': password
  }
  loginRequest.addEventListener('readystatechange', () => {
    if (loginRequest.readyState === 4 && loginRequest.status === 200) {
      const result = loginRequest.responseText
      if (result === 'RIGHT') {
        let accountDiv = document.getElementById('account_div')
        accountDiv.innerHTML = generateLoginElements(username);
      } else {
        clearAccountInfo();
      }
    }
  })
  loginRequest.send(JSON.stringify(data));
}
