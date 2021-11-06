// look in project-back/connection/NotSecuredQuerry.go
function log(){
    document.getElementById("wrong-login").innerHTML = "";
    fetch("http://localhost:8080/connect", {
            method: "POST",
            body: JSON.stringify({
                password: document.getElementById('password').value
            })
        }).then(async res => {
            if (res.ok) {
                document.getElementById('formContent').innerHTML = await res.text();
            } else {
                shakeWindow();
            }
        }).catch(() => {
            shakeWindow();
        })
}

function shakeWindow() {
    document.getElementById('formContent').classList.add("shake");
    let password = document.getElementById('password').value;
    setTimeout(function () {
        document.getElementById('formContent').classList.remove("shake");
        document.getElementById("wrong-login").innerHTML += "<p>" + password + " is wrong" + "</p>";
        document.getElementById('wrong-login').setAttribute('style', 'display:block')
    }, 250);
}