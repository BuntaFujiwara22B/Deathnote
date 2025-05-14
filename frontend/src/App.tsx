import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Navbar from './components/Navbar'
import Home from './pages/Home'
import RegisterPerson from './pages/RegisterPerson'
import DeathList from './pages/DeathList'

function App() {
  return (
    <BrowserRouter>
      <div
        className="min-h-screen bg-cover bg-center"
        style={{ backgroundImage: 'url(/fondoInicio.jpeg)' }}
      >
        <Navbar />
        <div className="flex items-center justify-center h-full">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/register" element={<RegisterPerson />} />
            <Route path="/list" element={<DeathList />} />
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  )
}

export default App
