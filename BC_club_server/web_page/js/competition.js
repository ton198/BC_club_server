
let bio_timeline = document.getElementById("bio_timeline");
let che_timeline = document.getElementById("che_timeline");

let getBioCompetitionInfo = new XMLHttpRequest();
getBioCompetitionInfo.open("POST", "/competition", false);
getBioCompetitionInfo.send("GET_COMPETITION");
let competitions = JSON.parse(getBioCompetitionInfo.responseText);
for (let data of competitions.bio_info) {
    bio_timeline.innerHTML += generateComment(data.time, data.name, data.detail);
}

for (let data of competitions.che_info) {
    che_timeline.innerHTML += generateComment(data.time, data.name, data.detail);
}

function generateComment(time, name, detail) {
    return "                    <div class=\"timeline-post\">\n" +
        "                        <div class=\"timeline-date\">" + time + "</div>\n" +
        "                        <div class=\"timeline-icon-con\">\n" +
        "                            <div class=\"timeline-icon\">\n" +
        "                                <div class=\"timeline-icon-inner\"></div>\n" +
        "                            </div>\n" +
        "                        </div>\n" +
        "                        <div class=\"timeline-content\">\n" +
        "                            <h3>" + name + "</h3>\n" +
        "                            <p>\n" +
        "                               " + detail + "\n" +
        "                            </p>\n" +
        "                        </div>\n" +
        "                    </div>"
}