async function shortenURL() {
    const url = document.getElementById("urlInput").value;
    const response = await fetch("/shorten", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ original: url })
    });
    const data = await response.json();
    document.getElementById("result").innerHTML = 
        `Short URL: <a href="${data.short}" target="_blank">${data.short}</a>`;
}
