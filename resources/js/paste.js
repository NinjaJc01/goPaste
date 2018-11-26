let mDiv = document.querySelector("#mainDiv");
let urlParams = new URLSearchParams(window.location.search);

class pasteClientside {
    constructor(id, content, timestamp) {
        this.id = id;
        this.timestamp = timestamp;
        this.content = content;
    }
    createElem() {
        let node = document.createElement("div");
        let timeDiv = document.createElement("div");
        let contentDiv = document.createElement("code");
        let pre = document.createElement("pre");
        pre.appendChild(contentDiv);
        node.id = "paste" + this.id;
        timeDiv.id = "timeDiv" + this.id;
        contentDiv.id = "contentDiv" + this.id;

        timeDiv.textContent = this.timestamp;
        contentDiv.textContent = this.content;
        contentDiv.classList.add("bg-dark");
        contentDiv.classList.add("text-light");

        node.appendChild(timeDiv);
        node.appendChild(pre);
        return node;
    }
}

function onload() {
    console.log(urlParams.get("id"));
    if (urlParams.get("id") !== null) {
        getPaste(urlParams.get("id"));
    } else console.log("null")
}


async function getPaste(id) {
    const response = await fetch("/api/paste/" + id);
    if (response.ok) {
        const paste = await response.json();
        console.log(paste);
        let thisPaste = new pasteClientside(
            paste.id,
            paste.content,
            paste.timestamp);
        mDiv.appendChild(thisPaste.createElem());
    } else {
        if (response.status == 404) {
            window.location.replace("/client/404"); //Not so good, breaks browser history
        }
        console.log(response);
    }
}

async function getPastes() {
    // let pastes = await fetch("/api/paste/list")
    //     .then(response => response.json());
    const response = await fetch('/api/paste/list');
    console.log(response);
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