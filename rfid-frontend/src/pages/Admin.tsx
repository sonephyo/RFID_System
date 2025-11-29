import { useState, useRef, useEffect } from "react";
import CardForm from "../components/CardForm";

export default function Admin() {
  const [connected, setConnected] = useState(false);
  const [selectedCard, setSelectedCard] = useState<string | null>(null);
  const portRef = useRef<SerialPort | null>(null);
  const bufferRef = useRef<string>("");

  useEffect(() => {
    if (!("serial" in navigator)) {
      alert("Browser doesn't support Web Serial API. Use Chrome!");
    }
  }, []);

  const requestConnection = async () => {
    try {
      const port = await navigator.serial.requestPort();
      await port.open({ baudRate: 115200 });
      portRef.current = port;
      setConnected(true);
      alert("Connected!");
      readData();
    } catch (e) {
      const error = e instanceof Error ? e.message : String(e);
      alert("Failed: " + error);
    }
  };

  const readData = async () => {
    if (!portRef.current?.readable) return;

    const decoder = new TextDecoderStream();
    portRef.current.readable.pipeTo(decoder.writable);
    const reader = decoder.readable.getReader();

    try {
      while (true) {
        const { value, done } = await reader.read();
        if (done) break;

        bufferRef.current += value;
        const lines = bufferRef.current.split("\n");

        while (lines.length > 1) {
          const line = lines.shift()?.trim();
          if (line && line.startsWith("Card:")) {
            const cardId = line.substring(5).trim();
            setSelectedCard(cardId);
            console.log("Card scanned:", cardId);
          }
        }
        bufferRef.current = lines[0] || "";
      }
    } catch (e) {
      console.error("Read error:", e);
    }
  };

  const clearCard = () => {
    setSelectedCard(null);
  };

  return (
    <div className="p-8 max-w-2xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">RFID Admin Panel</h1>

      <button
        onClick={requestConnection}
        disabled={connected}
        className="bg-blue-500 text-white px-6 py-2 rounded mb-4 disabled:bg-gray-400"
      >
        {connected ? "Connected ✓" : "Connect ESP32"}
      </button>

      <div className="bg-gray-100 p-4 rounded">
        <div className="flex justify-between items-center mb-2">
          <p className="font-bold">Scanned Card:</p>
          {selectedCard && (
            <button
              onClick={clearCard}
              className="text-sm text-gray-500 underline"
            >
              Clear
            </button>
          )}
        </div>

        {selectedCard ? (
          <div className="bg-white p-4 rounded shadow">
            <p className="text-sm text-gray-500">Card ID:</p>
            <span className="font-mono font-bold text-lg">{selectedCard}</span>
            <CardForm cardId={selectedCard} />
          </div>
        ) : (
          <p className="text-gray-500">
            {connected
              ? "Waiting for card scan..."
              : "Connect ESP32 and scan a card"}
          </p>
        )}
      </div>
    </div>
  );
}