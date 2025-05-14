import { useEffect, useState } from 'react';
import { getVictims } from '../services/api';

interface Victim {
  id: number;
  full_name: string;
  cause?: string;
  details?: string;
  created_at: string;
  image_url: string;
}

function DeathList() {
  const [victims, setVictims] = useState<Victim[]>([]);

  useEffect(() => {
    const loadVictims = async () => {
      try {
        const response = await getVictims();
        setVictims(response.data);
      } catch (error) {
        console.error("Error cargando víctimas:", error);
      }
    };
    loadVictims();
  }, []);

  return (
    <div className="flex flex-col items-center justify-center min-h-screen text-center px-4 py-20">
      <h2 className="text-4xl font-bold mb-8">
        <div className="bg-white text-black p-4 rounded-lg shadow-md">
          Lista de personas registradas
        </div>
      </h2>
      
      <div className="w-full max-w-2xl bg-black bg-opacity-70 p-6 rounded-lg">
        {victims.map((victim) => (
          <div key={victim.id} className="bg-gray-800 p-4 rounded-lg mb-4">
            <h3 className="text-xl font-semibold">{victim.full_name}</h3>
            <p className="text-gray-400">
              Causa: {victim.cause || 'Ataque al corazón'}
            </p>
            {victim.details && (
              <p className="text-gray-500 mt-2">{victim.details}</p>
            )}
            <p className="text-sm text-gray-600 mt-2">
              Registrado el: {new Date(victim.created_at).toLocaleDateString()}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default DeathList;