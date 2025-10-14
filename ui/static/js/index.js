require.config({ paths: { 'vs': 'https://cdn.jsdelivr.net/npm/monaco-editor@0.38.0/min/vs' }});

require(['vs/editor/editor.main'], function () {
  const editor = monaco.editor.create(document.getElementById('editor'), {
    value: `// Start coding here...\n`,
    language: 'javascript',
    theme: 'vs-dark',
    automaticLayout: true
  });

  // Connect WebSocket
  const roomId = prompt("Enter Room ID") || "default";
  const ws = new WebSocket(`ws://localhost:8989/opti-collab/ws?room=${roomId}`);

  ws.onopen = () => console.log("Connected to OptiCollab WebSocket");

  ws.onmessage = (event) => {
    const data = event.data;
    if (editor.getValue() !== data) {
      editor.setValue(data);
    }
  };

  // Send changes on every edit
  editor.onDidChangeModelContent((event) => {
    ws.send(editor.getValue());
  });
});
