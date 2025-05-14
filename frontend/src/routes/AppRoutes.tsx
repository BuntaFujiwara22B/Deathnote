import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Home from '../pages/Home'
import RegisterPerson from '../pages/RegisterPerson'
import DeathList from '../pages/DeathList'

function AppRoutes() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/register" element={<RegisterPerson />} />
        <Route path="/list" element={<DeathList />} />
      </Routes>
    </BrowserRouter>
  )
}

export default AppRoutes
