var prefix = "ws_";

var channel = null;
var connectBtn = document.getElementById("connectBtn"),
    disconnectBtn = document.getElementById("disconnectBtn");

connectBtn.onclick = function () {
    var {url, err} = getFinalUrl();
    if (err) {
        alert(err);
        return;
    }
    if (!url.startsWith("ws://") && !url.startsWith("wss://")) {
        alert("Invalid url!!!");
        return;
    }

    println("[建立连接]");

    connect(url);
    println(`url: [${url}].`);
};

disconnectBtn.onclick = function () {
    println("[断开连接]");

    disconnect();
};

/**
 * PS: EventSource 没有onclose事件.
 */
function connect(url) {
    disconnect();

    channel = new WebSocket(url);
    channel.onopen = function () {
        println("onopen");
    };
    channel.onmessage = function (e) {
        var data = e.data;

        if (data instanceof ArrayBuffer) {
            let blob = new Blob([data]), reader = new FileReader();

            reader.readAsText(blob, "UTF-8");
            reader.onload = () => {
                var text = reader.result;
                println("on binary message: " + text);
            };
        } else if (data instanceof Blob) {
            var reader = new FileReader();

            reader.readAsText(data, "UTF-8");
            reader.onload = () => {
                var text = reader.result;
                println("on binary message: " + text);
            };
        } else if (typeof data === "string") {
            var text = e.data;
            println("on text message: " + text);
        }
    };
    channel.onerror = function (e) {
        println("onerror");
        console.error(e);
    };
    channel.onclose = function (e) {
        println("onclose: code(" + e.code + "), reason(" + e.reason + "), wasClean(" + e.wasClean + ")");
        console.error(e);
    };
}

function disconnect() {
    if (channel == null) {
        return;
    }

    channel.close();
    channel = null;
}
