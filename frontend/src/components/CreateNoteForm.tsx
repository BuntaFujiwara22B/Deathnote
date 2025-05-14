import { useState, useEffect } from 'react';
import { registerVictim, updateCause, updateDetails } from '../services/api';

const CreateNoteForm = () => {
  const [name, setName] = useState('');
  const [cause, setCause] = useState('');
  const [details, setDetails] = useState('');
  const [imageUrl, setImageUrl] = useState('');
  const [phase, setPhase] = useState<'idle' | 'cause_timer' | 'details_timer'>('idle');
  const [timer, setTimer] = useState(40);
  const [victimId, setVictimId] = useState<number | null>(null);

  useEffect(() => {
    let interval: number;

    const tick = () => {
      setTimer(prev => {
        if (prev <= 1) {
          if (phase === 'cause_timer') handleTimeExpired();
          else if (phase === 'details_timer') handleDetailsTimeExpired();
          return 0;
        }
        return prev - 1;
      });
    };

    if (phase !== 'idle') {
      interval = window.setInterval(tick, 1000);
    }

    return () => window.clearInterval(interval);
  }, [phase]);

  const handleTimeExpired = async () => {
    if (victimId && !cause) {
      try {
        await updateCause(victimId, 'Ataque al corazón');
      } catch (error) {
        console.error(error);
      }
    }
    setPhase('details_timer');
    setTimer(400);
  };

  const handleDetailsTimeExpired = async () => {
    if (victimId && details) {
      try {
        await updateDetails(victimId, details);
      } catch (error) {
        console.error(error);
      }
    }
    resetForm();
  };

  const startCauseTimer = () => {
    if (name.trim() && imageUrl.trim()) {
      setPhase('cause_timer');
      setTimer(40);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    try {
      if (phase === 'cause_timer') {
        const response = await registerVictim({
          full_name: name,
          image_url: imageUrl
        });
        setVictimId(response.data.id);
        
        if (cause) {
          await updateCause(response.data.id, cause);
        }
        setPhase('details_timer');
        setTimer(400);
      } else if (phase === 'details_timer' && victimId && details) {
        await updateDetails(victimId, details);
        resetForm();
      }
    } catch (error) {
      console.error(error);
    }
  };

  const resetForm = () => {
    setName('');
    setCause('');
    setDetails('');
    setImageUrl('');
    setPhase('idle');
    setTimer(40);
    setVictimId(null);
  };

  const handleImageUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = (event) => {
        setImageUrl(event.target?.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  return (
    <div className="p-6 max-w-md mx-auto bg-gray-900 text-white rounded-lg shadow-lg">
      <h2 className="text-2xl font-bold mb-6 text-center">Death Note</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label htmlFor="name" className="block text-sm font-medium mb-1">
            Nombre Completo
          </label>
          <input
            id="name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            onFocus={startCauseTimer}
            placeholder="Escribe el nombre completo"
            className="w-full p-3 rounded-md bg-gray-800 border border-gray-700 focus:border-red-500 focus:ring-red-500"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="image" className="block text-sm font-medium mb-1">
            Rostro de la persona
          </label>
          <input
            id="image"
            type="file"
            accept="image/*"
            onChange={handleImageUpload}
            className="w-full text-sm text-gray-400 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-sm file:font-semibold file:bg-red-700 file:text-white hover:file:bg-red-600"
            required
          />
          {imageUrl && (
            <div className="mt-2">
              <img 
                src={imageUrl} 
                alt="Preview" 
                className="h-20 w-20 object-cover rounded-md border border-gray-700"
              />
            </div>
          )}
        </div>

        {phase === 'cause_timer' && (
          <div className="mb-4">
            <label htmlFor="cause" className="block text-sm font-medium mb-1">
              Causa de Muerte ({timer}s restantes)
            </label>
            <textarea
              id="cause"
              value={cause}
              onChange={(e) => setCause(e.target.value)}
              className="w-full p-3 rounded-md bg-gray-800 border border-gray-700 min-h-[100px]"
              placeholder="Describe la causa de muerte..."
            />
          </div>
        )}

        {phase === 'details_timer' && (
          <div className="mb-4">
            <label htmlFor="details" className="block text-sm font-medium mb-1">
              Detalles Específicos ({timer}s restantes)
            </label>
            <textarea
              id="details"
              value={details}
              onChange={(e) => setDetails(e.target.value)}
              className="w-full p-3 rounded-md bg-gray-800 border border-gray-700 min-h-[120px]"
              placeholder="Describe los detalles específicos..."
            />
          </div>
        )}

        <button
          type="submit"
          className="w-full py-3 px-4 rounded-md font-bold bg-red-700 hover:bg-red-600 transition-colors"
        >
          {phase === 'cause_timer' ? 'Registrar Causa' : 
          phase === 'details_timer' ? 'Registrar Detalles' : 
          'Comenzar Registro'}
        </button>
      </form>
    </div>
  );
};

export default CreateNoteForm;