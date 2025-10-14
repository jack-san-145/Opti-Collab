const roomId = prompt("Enter Room ID") || "default";

const ws = new WebSocket(`ws://localhost:8989/opti-collab/ws?room=${roomId}`);

ws.onopen = () => {
  console.log("Connected to OptiCollab WebSocket");
};

ws.onmessage = (event) => {
  const data = event.data;
  // Update code editor content (example for a textarea)
  const editor = document.getElementById("editor");
  if (editor.value !== data) {
    editor.value = data;
  }
};

// Send changes to server
const editor = document.getElementById("editor");
editor.addEventListener("input", (e) => {
  ws.send(editor.value);
});
