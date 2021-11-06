// look in project-back/connection/Querry.go

function log(){
    fetch("http://localhost:8080/connectSafe", {
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
    setTimeout(function () {
        document.getElementById('formContent').classList.remove("shake");
        document.getElementById('wrong-login').setAttribute('style', 'display:block')
    }, 250);
}