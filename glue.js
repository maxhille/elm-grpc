var service = require("./build/example_pb_service")
var pb = require("./build/example_pb")
var jspb = require("google-protobuf")
var grpc = require("grpc-web-client")
import elm from "./build/elm"

var app = elm.Main.fullscreen()

var host = window.location.protocol + "//" + window.location.host;
app.ports.search.subscribe(function(msg) {
  var req = new pb.SearchRequest();
  req.setQuery(msg.query);
  grpc.grpc.unary(service.SearchService.Search, {
    request: req,
    host: host,
    onEnd: function(res) {
      // TODO log error?;
    }
  });
});
