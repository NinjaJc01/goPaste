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
        contentDiv.classList.add("bg-dark");
        contentDiv.classList.add("text-light");
        node.appendChild(timeDiv);
        node.appendChild(contentDiv);
        return node;
    }
}

function onload() {
    getPastes
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

async function getPastes() {
    // let pastes = await fetch("/api/paste/list")
    //     .then(response => response.json());
    const response = await fetch('/api/paste/list');
    const pastes = await response.json();
    if (pastes !== null) {
        console.log(pastes);
        pastes.forEach(element => {
            let thisPaste = new pasteClientside(
                element.id,
                element.content,
                element.timestamp);
            mDiv.appendChild(thisPaste.createElem());
        });
    } else {
        console.log("null");
        mDiv.textContent = "No pastes yet! Check back soon."
    }
}