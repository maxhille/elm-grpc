var service = require("./build/example_pb_service");
var pb = require("./build/example_pb");
import bind from "./binder";
import elm from "./build/elm";

var app = elm.Main.fullscreen();
var host = window.location.protocol + "//" + window.location.host;

bind(host, service.SearchService.Search,
  app.ports.search, pb.SearchRequest,
  app.ports.respond, pb.SearchResponse);
