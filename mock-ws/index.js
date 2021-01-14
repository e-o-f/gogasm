const WebSocket = require("ws");
const fs = require("fs");

const ws = new WebSocket("ws://localhost:5001");
const wss = new WebSocket.Server({ port: 5001 });

const dirListContent = fs.readFileSync("dir-list-small.txt", "utf8");
const dirListLines = dirListContent.split("\n");

const wait = () => {
  return new Promise((resolve) => setTimeout(resolve, 300));
};

wss.on("connection", async function connection(ws) {
  ws.on("message", function incoming(message) {
    console.log("received: %s", message);
  });

  let i = 0;
  while (true) {
    const line = dirListLines[i % dirListLines.length];
    if (Boolean(line) && line[0] !== "#") {
      ws.send(line);
      console.log("Sent line: ", line);
      await wait();
    }
    i++;
  }
});
