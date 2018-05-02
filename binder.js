var grpc = require("grpc-web-client");

export default function bind(host, service, reqPort, reqType, resPort,
  resType) {
  reqPort.subscribe(function(msg) {
    grpc.grpc.unary(service, {
      request: pbize(msg, reqType),
      host: host,
      onEnd: function(res) {
        resPort.send(jsonize(res.message));
      }
    });
  });
}

function pbize(json, type) {
  var pb = new type()

  for (var p in json) {
    var method1 = p.substring(0, 1).toUpperCase();
    var methodN = p.substring(1);
    var value = json[p];

    var method;
    if (typeof value === "array") {
      method = "set" + method1 + methodN + "List";
    } else {
      method = "set" + method1 + methodN;
    }

    pb[method](value)
  }

  return pb;
}

function jsonize(pbObject) {
  var json = {};
  for (var p in pbObject) {
    if (typeof p != "string") {
      continue;
    }

    if (p.indexOf("get") != 0) {
      continue;
    }

    if (p == "getExtension" || p == "getJsPbMessageId") {
      continue;
    }

    var value = pbObject[p]();
    var key1 = p.substring(3, 4).toLowerCase();
    var keyN = p.substring(4);

    var key;
    if (keyN.substring(keyN.length - 4) == "List") {
      key = key1 + keyN.substring(0, keyN.length - 4);
    } else {
      key = key1 + keyN;
    }

    json[key] = value;
  }
  return json;
}
