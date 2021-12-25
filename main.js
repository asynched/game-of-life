const main = async () => {
  const goRuntime = new Go()
  const source = await WebAssembly.instantiateStreaming(
    fetch('main.wasm'),
    goRuntime.importObject
  )
  goRuntime.run(source.instance)
}

main()
