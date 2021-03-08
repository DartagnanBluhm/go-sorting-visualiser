function populateVisualisation(length) {
    container = document.getElementById("bar-container")
    width = parseInt(window.getComputedStyle(container, null).getPropertyValue('width'))
    height = parseInt(window.getComputedStyle(container, null).getPropertyValue('height'))
    if (length > width/2) {
        length = Math.floor(width/2);
        console.log(length)
    }
    console.log(width, height)
    container.innerHTML = ""
    for (i = 0; i < length; i++) {
        value = Math.floor(Math.random() * height + 1)
        var div = document.createElement("div")
        div.id = "bar-" + i
        div.className = "bar"
        div.style.height = value + "px"
        div.style.backgroundColor = "white"
        div.style.flex = "1 1 auto"
        container.appendChild(div)
    }
}