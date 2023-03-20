
function validateForm() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    if (username == "admin" && password == "admin") {
        alert("Inicio de sesión exitoso");//dashboard admin
    } else {
        alert("El nombre de usuario o la contraseña son incorrectos");//dashboard usuario   
    }
}