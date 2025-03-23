document.getElementById("btnHello").addEventListener("click", function() {
	fetchEndpoint('/hello', 'GET');
});
document.getElementById("btnTime").addEventListener("click", function() {
	fetchEndpoint('/time', 'GET');
});
document.getElementById("btnPost").addEventListener("click", sendPost);


function fetchEndpoint(endpoint, method) {
    fetch(endpoint, { method: method })
        .then(response => response.text())
        .then(data => {
            document.getElementById("output").textContent = data;
        })
        .catch(err => {
            document.getElementById("output").textContent = "Error: " + err;
        });
}

function sendPost() {
const msg = document.getElementById("postMessage").value;
fetch('/post', {
    method: 'POST',
    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
    body: 'message=' + encodeURIComponent(msg)
})
.then(response => response.text())
.then(data => {
    document.getElementById("output").textContent = data;
})
.catch(err => {
    document.getElementById("output").textContent = "POST Error: " + err;
});
}