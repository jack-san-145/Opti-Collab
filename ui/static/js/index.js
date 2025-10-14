let editor;

document.addEventListener("DOMContentLoaded", function () {
  require.config({
    paths: { 'vs': 'https://cdn.jsdelivr.net/npm/monaco-editor@0.38.0/min/vs' }
  });

  require(['vs/editor/editor.main'], function () {
    editor = monaco.editor.create(document.getElementById('editor'), {
      value: `// Start coding here...\n`,
      language: 'javascript',
      theme: 'vs-dark',
      automaticLayout: true
    });

    // Handle language switching
    document.getElementById("language-select").addEventListener("change", (e) => {
      const lang = e.target.value;
      monaco.editor.setModelLanguage(editor.getModel(), lang);
    });
  });
});

// Run code button
document.getElementById("run-btn").addEventListener("click", function () {
  const code = editor.getValue();
  const lang = document.getElementById("language-select").value;

  console.log("▶️ Running code in:", lang);
  console.log("💻 Code:\n", code);
  Send_code_to_server(code,lang)

  // Optional: Send to your backend for execution
  // fetch("/opti-collab/run-code", {
  //   method: "POST",
  //   headers: { "Content-Type": "application/json" },
  //   body: JSON.stringify({ language: lang, code: code }),
  // })
  // .then(res => res.text())
  // .then(output => console.log("Output:", output))
  // .catch(err => console.error("Error:", err));
});


async function Send_code_to_server(code,lang){
  const response = await fetch("/opti-collab/run-code",{
    method : "POST",
    content:"application/json",
    body : JSON.stringify({"code":code,"language":lang})
  })
  const data =await response.json()
  console.log("data - ",data)

}


// document.addEventListener("DOMContentLoaded", function () {
//   require.config({
//     paths: { 'vs': 'https://cdn.jsdelivr.net/npm/monaco-editor@0.38.0/min/vs' }
//   });

//   require(['vs/editor/editor.main'], function () {
//     const editor = monaco.editor.create(document.getElementById('editor'), {
//       value: `// Start coding here...\n`,
//       language: 'javascript',
//       theme: 'vs-dark',
//       automaticLayout: true
//     });


//     // Connect WebSocket
//   //   const roomId = prompt("Enter Room ID") || "default";
//   //   const ws = new WebSocket(`ws://localhost:8989/opti-collab/ws?room=${roomId}`);

//   //   ws.onopen = () => console.log("✅ Connected to WebSocket room:", roomId);

//   //   let suppress = false;

//   //   ws.onmessage = (event) => {
//   //     const data = event.data;
//   //     if (editor.getValue() !== data) {
//   //       suppress = true;
//   //       editor.setValue(data);
//   //       suppress = false;
//   //     }
//   //   };

//   //   editor.onDidChangeModelContent(() => {
//   //     if (!suppress) {
//   //       ws.send(editor.getValue());
//   //     }
//   //   });

//   //   // Dynamic language switching
//   //   document.getElementById("language-select").addEventListener("change", (e) => {
//   //     const lang = e.target.value;
//   //     monaco.editor.setModelLanguage(editor.getModel(), lang);
//   //   });
//   });
// });


// function Run_code(){
//   const code=document.getElementById('editor').value
//   console.log("code - ",code)

// }

