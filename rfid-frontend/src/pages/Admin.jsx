import { useState, useRef, useEffect } from "react";
import React from "react";

export default function Admin() {
  const [connected, setConnected] = useState(false);
  const [scannedCards, setScannedCards] = useState([]);
  const portRef = useRef(null);

  useEffect(() => {
    if (!("serial" in navigator)) {
      alert("Browser doesn't support Web Serial API. Use Chrome!");
    }
  }, []);

  const requestConnection = async () => {
    try {
      const port = await navigator.serial.requestPort();
      await port.open({ baudRate: 115200 }); // ESP32 uses 115200
      portRef.current = port;
      setConnected(true);
      alert("Connected!");
      readData();
    } catch (e) {
      alert("Failed: " + e.message);
    }
  };

  const readData = async () => {
    const decoder = new TextDecoderStream();
    portRef.current.readable.pipeTo(decoder.writable);
    const reader = decoder.readable.getReader();

    while (true) {
      const { value, done } = await reader.read();
      if (done) break;
      
      const trimmed = value.trim();
      if (trimmed) {
        setScannedCards(prev => [...prev, {
          id: trimmed,
          time: new Date().toLocaleTimeString()
        }]);
      }
    }
  };

  return (
    <div className="p-8">
      <h1 className="text-2xl font-bold mb-4">RFID Scanner Test</h1>
      
      <button 
        onClick={requestConnection}
        disabled={connected}
        className="bg-blue-500 text-white px-6 py-2 rounded mb-4 disabled:bg-gray-400"
      >
        {connected ? "Connected ✓" : "Connect ESP32"}
      </button>

      <div className="bg-gray-100 p-4 rounded">
        <p className="font-bold mb-2">Scanned Cards:</p>
        {scannedCards.length === 0 && <p className="text-gray-500">No cards scanned yet...</p>}
        {scannedCards.map((card, i) => (
          <div key={i} className="bg-white p-2 mb-2 rounded shadow">
            <span className="font-mono font-bold">{card.id}</span>
            <span className="text-gray-500 ml-4 text-sm">{card.time}</span>
          </div>
        ))}
      </div>
    </div>
  );
}