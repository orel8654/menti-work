const PrefixUrl = "http://127.0.0.1:8010"
async function Login() {
    document.getElementById("loginForm").addEventListener("submit", function(event) {
        event.preventDefault();
        let username = document.getElementById("username").value;
        let password = document.getElementById("password").value;

        fetch(PrefixUrl + "/api/v1/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ username: username, password: password })
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error("Ошибка сети");
                }
                return response.json();
            })
            .then(data => {
                // Обработка ответа от сервер
                console.log(data);
                saveTokenToSessionStorage(data.token)
            })
            .catch(error => {
                console.error("Произошла ошибка:", error);
            });
    });
}

function saveTokenToSessionStorage(token) {
    sessionStorage.setItem('token', token);
}
