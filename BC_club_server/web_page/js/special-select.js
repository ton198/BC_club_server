
function intoElement(event) {
    event.target.style.borderColor = 'rgba(0, 0, 0, 100%)'
    event.target.style.backgroundColor = 'rgba(0, 0, 0, 50%)'
}

function outElement(event) {
    event.target.style.borderColor = 'rgba(0, 0, 0, 0%)'
    event.target.style.backgroundColor = 'rgba(0, 0, 0, 0%)'
}

const divs = document.getElementsByClassName('select_style')

for (const div of divs) {
    div.addEventListener('mouseout', outElement)
    div.addEventListener('mouseover', intoElement)
}
