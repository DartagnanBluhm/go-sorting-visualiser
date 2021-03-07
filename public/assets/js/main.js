function populateVisualisation(length) {
    height = document.getElementById("bar-container").clientHeight
    console.log(height)
    document.getElementById("bar-container").innerHTML = ""
    for (i = 0; i < length; i++) {
        value = Math.floor(Math.random() * height + 1)
        var div = document.createElement("div")
        div.id = "bar-" + i
        div.className = "bar"
        div.style.height = value + "px"
        div.style.backgroundColor = "white"
        div.style.flex = "1 1 auto"
        document.getElementById("bar-container").appendChild(div)
        console.log(div)
    }
}