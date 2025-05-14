function Home() {
    return (
      <div className="flex flex-col items-center justify-center h-screen text-center px-4">
        <h1 className="text-5xl font-bold mb-4">
          <div className="bg-white text-black p-4 rounded-lg shadow-md">
            Bienvenido a la Death Note
          </div>
        </h1>
        <p className="text-lg text-gray-700 dark:text-gray-300 max-w-xl">
          Esta es una herramienta inspirada en el universo de Death Note para registrar nombres. Usa el menú superior para navegar y comenzar.
        </p>
      </div>
    )
  }
  
  export default Home
  