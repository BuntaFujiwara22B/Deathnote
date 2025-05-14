import { Link } from 'react-router-dom'

const Navbar = () => (
  <nav className="bg-black text-white p-4">
    <ul className="flex space-x-4">
      <li><Link to="/" className="hover:text-gray-400">Inicio</Link></li>
      <li><Link to="/list" className="hover:text-gray-400">Mis Notas</Link></li> {/* Cambiar a "/list" */}
      <li><Link to="/register" className="hover:text-gray-400">Registrar Persona</Link></li> {/* Cambiar a "/register" */}
    </ul>
  </nav>
)

export default Navbar
