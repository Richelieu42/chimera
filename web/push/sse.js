var prefix = "sse_";

var channel = null;
var connectBtn = document.getElementById("connectBtn"),
    disconnectBtn = document.getElementById("disconnectBtn");

connectBtn.onclick = function () {
    var {url, err} = getFinalUrl();
    if (err) {
        alert(err);
        return;
    }
    if (!url.startsWith("http://") && !url.startsWith("https://")) {
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

    channel = new EventSource(url);
    channel.onopen = function (event) {
        println("onopen");
    };
    channel.onmessage = function (e) {
        var origin = e.origin,
            id = e.lastEventId,
            event = e.type,
            data = e.data;

        println("onmessage: origin(" + origin + "), id(" + id + "), event(" + event + "), data(" + data + ")");
    };
    // 自定义事件
    channel.addEventListener('test', function (e) {
        var origin = e.origin,
            id = e.lastEventId,
            event = e.type,
            data = e.data;

        println("ontest: origin(" + origin + "), id(" + id + "), event(" + event + "), data(" + data + ")");
    }, false);
    channel.onmessage = function (e) {
        var origin = e.origin,
            id = e.lastEventId,
            event = e.type,
            data = e.data;

        println("onmessage: origin(" + origin + "), id(" + id + "), event(" + event + "), data(" + data + ")");
    };
    // 如果发生通信错误（比如连接中断），就会触发error事件
    channel.onerror = function (event) {
        println("onerror");
        console.error(event);
    };
}

function disconnect() {
    if (channel == null) {
        return;
    }

    channel.close();
    channel = null;
}

