var bars = new Array()
var stop = false

function shuffleBars() {
    populateVisualisation(document.getElementById("sorting-bar-number").value)
}

function displayValue() {
    document.getElementById("sort-speed-label").innerHTML =
        "action speed (" +
        document.getElementById("sort-speed").value +
        "ms between actions)"
}

function populateVisualisation(length) {
    bars = new Array();
    container = document.getElementById("bar-container");
    width = parseInt(
        window.getComputedStyle(container, null).getPropertyValue("width")
    );
    height = parseInt(
        window.getComputedStyle(container, null).getPropertyValue("height")
    );
    if (length > width / 2) {
        length = Math.floor(width / 2);
    }
    container.innerHTML = "";
    for (i = 0; i < length; i++) {
        value = Math.floor(Math.random() * height + 1)
        var div = document.createElement("div")
        div.id = "bar-" + i
        div.className = "bar not-selected"
        div.style.height = value + "px"
        div.style.flex = "1 1 auto"
        container.appendChild(div)
        bars.push(div.id)
    }
}

function startSorting() {
    algType = document.getElementById("alg-dropdown").value
    var xhr = new XMLHttpRequest()
    xhr.open("POST", "http://localhost:8000/api/v1/sort", true)
    xhr.setRequestHeader("Content-Type", "application/json")
    var arr = new Array()
    for (i = 0; i < bars.length; i++) {
        arr.push({ value: document.getElementById(bars[i]).clientHeight })
    }
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                jsonArrayChanges = JSON.parse(xhr.responseText)
                animate(jsonArrayChanges)
            }
        }
    }
    xhr.send(JSON.stringify({ algorithmType: algType, arrayLength: bars.length, array: arr }))
}

function queueStop(){
    stop = true
}

function checkStopped() {

}

async function animate(changes) {
    if (changes != null) {
        for (i = 0; i < changes.length; i++) {
            if (stop) {
                window.setTimeout(checkStopped, 100)
            }
            first = changes[i]["first-index"]
            second = changes[i]["second-index"]
            document.getElementById(bars[first]).classList.replace("not-selected", "selected1")
            await sleep(document.getElementById("sort-speed").value)
            document.getElementById(bars[second]).classList.replace("not-selected", "selected2")
            await sleep(document.getElementById("sort-speed").value)
            document.getElementById(bars[first]).style.height = changes[i]["first-value"] + "px"
            document.getElementById(bars[first]).classList.replace("selected1", "selected2")
            document.getElementById(bars[second]).style.height = changes[i]["second-value"] + "px"
            document.getElementById(bars[second]).classList.replace("selected2", "selected1")
            await sleep(document.getElementById("sort-speed").value)
            document.getElementById(bars[first]).classList.replace("selected2", "not-selected")
            await sleep(document.getElementById("sort-speed").value)
            document.getElementById(bars[second]).classList.replace("selected1", "not-selected")
        }
    }

    for (i = bars.length - 1; i > 0; i--) {
        if (document.getElementById(bars[i]).clientHeight >= document.getElementById(bars[i - 1]).clientHeight) {
            document.getElementById(bars[i]).classList.replace("not-selected", "selected1")
        } else {
            document.getElementById(bars[i]).classList.replace("not-selected", "selected2")
        }
        await sleep(document.getElementById("sort-speed").value)
    }
    document.getElementById(bars[0]).classList.replace("not-selected", "selected1")
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}