let mDiv = document.querySelector("#mainDiv");
class pasteClientside {
    constructor(id, content, timestamp) {
        this.id = id;
        this.timestamp = timestamp;
        this.content = content;
    }
    createElem() {
        let node = document.createElement("div");
        let timeDiv = document.createElement("div");
        let contentDiv = document.createElement("div");

        node.id = "paste" + this.id;
        timeDiv.id = "timeDiv" + this.id;
        contentDiv.id = "contentDiv" + this.id;

        timeDiv.textContent = this.timestamp;
        contentDiv.textContent = this.content;

        node.appendChild(timeDiv);
        node.appendChild(contentDiv);
        return node;
    }
}

function onload() {
    mDiv.textContent = "Hello there, js says hi innit";
}


async function getPaste(id) {
    let paste = await fetch("/api/paste/" + id)
        .then(response => response.json());
    console.log(paste);
    let thisPaste = new pasteClientside(
        paste.id,
        paste.content,
        paste.timestamp);
    mDiv.appendChild(thisPaste.createElem());
}