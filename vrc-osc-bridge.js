function sendosc(address, message, callback) {
    var server = "127.0.0.1";
    var port = "59000";

    var xhr = new XMLHttpRequest();
    xhr.open('POST', "http://" + server + ":" + port + address);
    xhr.onload = function () {
        callback(xhr.status);
    };
    xhr.onerror = function () {
        callback(xhr.status);
    };
    xhr.send(message);
}
