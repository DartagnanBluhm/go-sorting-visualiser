var bars = new Array()
var stop = false
var sorting = false
var changesIndex = 0
var jsonArrayChanges

function shuffleBars() {
    stop = true
    populateVisualisation(document.getElementById("sorting-bar-number").value)
    changesIndex = 0
    window.setTimeout(function () { stop = false; }, 50);
    sorting = false
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
    if (!sorting) {
        if (!stop) {
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
        } else {
            stop = false
            animate(jsonArrayChanges)
        }
    }
}

function queueStop() {
    stop = true
}

async function animate(changes) {
    sorting = true
    if (changes != null) {
        for (; changesIndex < changes.length; changesIndex++) {
            if (stop) {
                break
            }
            first = changes[changesIndex]["first-index"]
            second = changes[changesIndex]["second-index"]
            document.getElementById(bars[first]).classList.replace("not-selected", "selected1")
            await sleep(document.getElementById("sort-speed").value)
            document.getElementById(bars[second]).classList.replace("not-selected", "selected2")
            await sleep(document.getElementById("sort-speed").value)
            document.getElementById(bars[first]).style.height = changes[changesIndex]["first-value"] + "px"
            document.getElementById(bars[first]).classList.replace("selected1", "selected2")
            document.getElementById(bars[second]).style.height = changes[changesIndex]["second-value"] + "px"
            document.getElementById(bars[second]).classList.replace("selected2", "selected1")
            await sleep(document.getElementById("sort-speed").value)
            document.getElementById(bars[first]).classList.replace("selected2", "not-selected")
            await sleep(document.getElementById("sort-speed").value)
            document.getElementById(bars[second]).classList.replace("selected1", "not-selected")
        }
    }
    if (!stop) {
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
    sorting = false
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}