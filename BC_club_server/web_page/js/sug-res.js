
let subject = document.getElementById("subject");
let dataContainer = document.getElementById("data");
let loading = document.getElementById("loading");

let requestInfo = "";

switch (subject.innerHTML) {
    case "生物AP考前建议":
        requestInfo = "BIO_SUGGESTION";
        break;
    case "化学AP考前建议":
        requestInfo = "CHE_SUGGESTION";
        break;
    case "生物AP准备资料":
        requestInfo = "BIO_RESOURCE";
        break;
    case "化学AP准备资料":
        requestInfo = "CHE_RESOURCE";
        break;
}

let request = new XMLHttpRequest();
request.open("POST", "/res_sug");
request.addEventListener('readystatechange', () => {
    if (request.readyState === 4 && request.status === 200) {
        let parent = loading.parentNode;
        parent.removeChild(loading);
        dataContainer.innerHTML += request.responseText;
    }
});
request.send(requestInfo);
