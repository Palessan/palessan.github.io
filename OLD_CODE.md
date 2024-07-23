<!-- <script>
    if (WebAssembly){
        const go = new Go();
        WebAssembly
        .instantiateStreaming(fetch("main.wasm"), go.importObject)
        .then((result) => {
            go.run(result.instance);
        });
}
</script> -->
<!-- <h1>It takes around 23 seconds normally to load the first time. But sometimes github is slow and takes minutes</h1>
<p>Please note that Safari is not supported</p> -->
<!-- <img src="assets/images/other/loading.gif" alt="Loading"> -->
<!-- <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="wasm_exec.js"></script>
    <script>
        if (WebAssembly){
            const go = new Go();
            WebAssembly
            .instantiateStreaming(fetch("main.wasm"), go.importObject)
            .then((result) => {
                go.run(result.instance);
            });
    }
    </script>
    <title>Hello, World!</title>
</head>
<body>
    <h1>Hello, World!</h1>
    <p>This is a simple Hello, World! website.</p>
</body> -->

<!-- <!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Ebiten WASM Example</title>
    <script src="wasm_exec.js"></script>
    <script>
        async function loadWasm() {
            if (!WebAssembly) {
                console.error('WebAssembly is not supported in this browser');
                return;
            }

            const go = new Go();
            try {
                const result = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject);
                go.run(result.instance);
            } catch (err) {
                console.error('Failed to load WebAssembly module', err);
            }
        }

        window.onload = loadWasm;
    </script>
</head>
<body>
    <h1>Ebiten WASM Example</h1>
</body>
</html> -->

<!-- <script type="text/javascript">
    let audioContext;

    function initializeAudioContext() {
        if (!audioContext) {
            audioContext = new (window.AudioContext || window.webkitAudioContext)();
        }
    }

    document.addEventListener('click', function() {
        initializeAudioContext();
        if (audioContext.state === 'suspended') {
            audioContext.resume();
        }
    });
</script> -->