var bars = new Array();

function shuffleBars() {
    populateVisualisation(document.getElementById("sorting-bar-number").value);
}

function displayValue() {
    document.getElementById("sort-speed-label").innerHTML =
        "action speed (time(" +
        document.getElementById("sort-speed").value +
        "ms) between actions)";
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
        value = Math.floor(Math.random() * height + 1);
        var div = document.createElement("div");
        div.id = "bar-" + i;
        div.className = "bar not-selected";
        div.style.height = value + "px";
        div.style.backgroundColor = "white";
        div.style.flex = "1 1 auto";
        container.appendChild(div);
        bars.push(div.id)
    }
}

function insertionsort(arr, sleepTime) {
    for (i = 1; i < arr.length; i++) {
        keyDiv = arr[i]
        key = document.getElementById(keyDiv).clientHeight
        document.getElementById(keyDiv).classList.replace("not-selected", "selected1")
        for (j = i - 1; j >= 0 && document.getElementById(arr[j]).clientHeight > key; j--) {
            document.getElementById(arr[j]).classList.replace("not-selected", "selected2")
            sleep(sleepTime)
            document.getElementById(arr[j+1]).style.height = document.getElementById(arr[j]).clientHeight + "px"
            arr[j + 1] = arr[j];
            sleep(sleepTime)
            document.getElementById(arr[j]).classList.replace("selected2", "not-selected")
            sleep(sleepTime)
        }
        document.getElementById(keyDiv).classList.replace("selected1", "not-selected")
        document.getElementById(arr[j+1]).style.height = key + "px"
        arr[j + 1] = keyDiv;
        sleep(sleepTime)
    }
    console.log(bars)
}

function startSorting() {
    algType = document.getElementById("alg-dropdown").value
    timeDelay = document.getElementById("sort-speed").value
    console.log(algType, timeDelay)
    switch (algType) {
        case "insertion":
            insertionsort(bars, timeDelay)
            break;
    }
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}