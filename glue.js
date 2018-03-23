var service = require("./build/example_pb_service")
var pb = require("./build/example_pb")
var jspb = require("google-protobuf")
var grpc = require("grpc-web-client")
import elm from "./build/elm"

var app = elm.Main.fullscreen()

function connect(app) {
  var host = window.location.protocol + "//" + window.location.host;
  app.ports.addBank.subscribe(function(msg) {
    var bank = new pb.AddBankRequest();
    bank.setId(msg.id);
    bank.setPin(msg.pin);
    bank.setBlz(msg.blz);
    grpc.grpc.unary(service.AccountService.AddBank, {
      request: bank,
      host: host,
      onEnd: function(res) {
        // TODO log error?;
      }
    });
  });

  app.ports.getBanks.subscribe(function(msg) {
    var empty = new pb.Empty();
    grpc.grpc.unary(service.AccountService.GetBanks, {
      request: empty,
      host: host,
      onEnd: function(res) {
        if (res.status == 2) {
          // TODO error
          console.log(res.statusMessage);
        } else {
          var msg = {
            "banks": res.message.getBanksList().map(function(bank) {
              return {
                "id": bank.getId(),
                "blz": bank.getBlz(),
                "updating": bank.getUpdating(),
                "accounts": bank.getAccountsList().map(function(
                  account) {
                  return {
                    "id": account.getId(),
                    "name": account.getName()
                  }
                })
              }
            })
          };

          app.ports.setBanks.send(msg);
        }
      }
    });
  });

  app.ports.refresh.subscribe(function(msg) {
    var req = new pb.RefreshRequest();
    req.setId(msg.id)
    grpc.grpc.unary(service.AccountService.RefreshBank, {
      request: req,
      host: host,
      onEnd: function(res) {
        if (res.status == 2) {
          // TODO error
          console.log(res.statusMessage);
        }
      }
    });
  });

  app.ports.getUserEmpty.subscribe(function(msg) {
    var empty = new pb.Empty();
    grpc.grpc.unary(service.UserService.GetUser, {
      request: empty,
      host: host,
      onEnd: function(res) {
        if (res.status == 2) {
          // TODO error
          console.log(res.statusMessage);
        } else {
          app.ports.setUser.send(res.message.toObject());
        }
      }
    });
  });
}
