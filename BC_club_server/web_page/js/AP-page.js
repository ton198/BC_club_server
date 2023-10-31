
let bioElements = document.getElementsByClassName("bio");

for (let element of bioElements) {
    element.addEventListener('click', onBioClick);
}

let cheElements = document.getElementsByClassName("che");
for (let element of cheElements) {
    element.addEventListener('click', onCheClick);
}

function onBioClick(event) {
    switch (event.target.innerHTML) {
        case "建议":
            window.location.href = '/AP/bio/sug';
            break;
        case "资料":
            window.location.href = '/AP/bio/res';
            break;
    }
}

function onCheClick(event) {
    switch (event.target.innerHTML) {
        case "建议":
            window.location.href = '/AP/che/sug';
            break;
        case "资料":
            window.location.href = '/AP/che/res';
            break;
    }
}