import { useState, useRef, useEffect } from "react";
import CardForm from "../components/CardForm";
import ClassManager from "../components/ClassManager";
import AttendanceReport from "../components/AttendanceReport";

export default function Admin() {
  const [connected, setConnected] = useState(false);
  const [selectedCard, setSelectedCard] = useState<string | null>(null);
  const [activeTab, setActiveTab] = useState<"scanner" | "classes" | "reports">("scanner");
  const portRef = useRef<SerialPort | null>(null);
  const bufferRef = useRef<string>("");

  useEffect(() => {
    if (!("serial" in navigator)) {
      alert("Browser doesn't support Web Serial API. Use Chrome!");
    }
  }, []);

  const sendMessageToESP32 = async (message: string) => {
    if (!portRef.current?.writable) {
      console.error("Port not writable");
      return;
    }
    const encoder = new TextEncoder();
    const writer = portRef.current.writable.getWriter();
    try {
      await writer.write(encoder.encode(message + "\n"));
      console.log("Sent to ESP32:", message);
    } catch (e) {
      console.error("Write error:", e);
    } finally {
      writer.releaseLock();
    }
  };

  const requestConnection = async () => {
    try {
      const port = await navigator.serial.requestPort();
      await port.open({ baudRate: 115200 });
      portRef.current = port;
      setConnected(true);
      await sendMessageToESP32("MESSAGE: Connected!");
    } catch (e) {
      const error = e instanceof Error ? e.message : String(e);
      alert("Connection failed: " + error);
    }
    readData();
  };

  const readData = async () => {
    if (!portRef.current?.readable) return;
    const decoder = new TextDecoderStream();
    portRef.current.readable.pipeTo(decoder.writable as WritableStream<Uint8Array>);
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
    <div className="min-h-screen bg-zinc-950 text-white">
      <div className="fixed inset-0 bg-[linear-gradient(rgba(255,255,255,.02)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,.02)_1px,transparent_1px)] bg-[size:64px_64px] pointer-events-none" />
      
      <div className="relative max-w-4xl mx-auto px-6 py-12">
        <header className="mb-12">
          <div className="flex items-center gap-3 mb-2">
            <div className={`w-2 h-2 rounded-full ${connected ? "bg-emerald-500 animate-pulse" : "bg-zinc-600"}`} />
            <span className="text-xs uppercase tracking-widest text-zinc-500">
              {connected ? "Device Connected" : "Offline"}
            </span>
          </div>
          <h1 className="text-3xl font-light tracking-tight">
            RFID <span className="text-zinc-500">Admin</span>
          </h1>
        </header>

        <button
          onClick={requestConnection}
          disabled={connected}
          className={`w-full mb-8 py-4 border transition-all duration-300 ${
            connected
              ? "border-emerald-500/30 bg-emerald-500/5 text-emerald-500"
              : "border-zinc-800 hover:border-zinc-600 bg-zinc-900/50 hover:bg-zinc-900 text-zinc-400 hover:text-white"
          }`}
        >
          <div className="flex items-center justify-center gap-3">
            {connected ? (
              <>
                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
                </svg>
                <span className="font-medium">Connected to ESP32</span>
              </>
            ) : (
              <>
                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
                <span className="font-medium">Connect Device</span>
              </>
            )}
          </div>
        </button>

        <div className="flex gap-1 mb-8 p-1 bg-zinc-900/50 border border-zinc-800/50">
          <button
            onClick={() => setActiveTab("scanner")}
            className={`flex-1 py-3 text-sm font-medium transition-all duration-200 ${
              activeTab === "scanner"
                ? "bg-zinc-800 text-white"
                : "text-zinc-500 hover:text-zinc-300"
            }`}
          >
            Card Scanner
          </button>
          <button
            onClick={() => setActiveTab("classes")}
            className={`flex-1 py-3 text-sm font-medium transition-all duration-200 ${
              activeTab === "classes"
                ? "bg-zinc-800 text-white"
                : "text-zinc-500 hover:text-zinc-300"
            }`}
          >
            Manage Classes
          </button>
          <button
            onClick={() => setActiveTab("reports")}
            className={`flex-1 py-3 text-sm font-medium transition-all duration-200 ${
              activeTab === "reports"
                ? "bg-zinc-800 text-white"
                : "text-zinc-500 hover:text-zinc-300"
            }`}
          >
            Attendance Reports
          </button>
        </div>

        {activeTab === "scanner" && (
          <div className="border border-zinc-800/50 bg-zinc-900/30">
            <div className="p-6 border-b border-zinc-800/50">
              <div className="flex items-center justify-between">
                <div className="flex items-center gap-3">
                  <div className="w-10 h-10 bg-zinc-800 flex items-center justify-center">
                    <svg className="w-5 h-5 text-zinc-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 9a2 2 0 10-4 0v5a2 2 0 01-2 2h6m-6-4h4m8 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <div>
                    <h2 className="font-medium">Card Scanner</h2>
                    <p className="text-sm text-zinc-500">Scan a card to view or edit user</p>
                  </div>
                </div>
                {selectedCard && (
                  <button
                    onClick={clearCard}
                    className="text-xs text-zinc-500 hover:text-zinc-300 transition-colors"
                  >
                    Clear
                  </button>
                )}
              </div>
            </div>

            <div className="p-6">
              {selectedCard ? (
                <div>
                  <div className="flex items-center gap-4 mb-6">
                    <div className="flex-1 bg-zinc-800/50 p-4 border-l-4 border-emerald-500">
                      <span className="text-xs text-zinc-500 uppercase tracking-wider">Card ID</span>
                      <p className="font-mono text-lg text-white mt-1">{selectedCard}</p>
                    </div>
                  </div>
                  <CardForm cardId={selectedCard} />
                </div>
              ) : (
                <div className="text-center py-16">
                  <div className="w-16 h-16 mx-auto mb-6 border-2 border-dashed border-zinc-800 rounded-full flex items-center justify-center">
                    <svg className="w-8 h-8 text-zinc-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M15 9a2 2 0 10-4 0v5a2 2 0 01-2 2h6m-6-4h4m8 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <p className="text-zinc-500 mb-2">
                    {connected ? "Waiting for card scan..." : "Connect device to start scanning"}
                  </p>
                  <p className="text-xs text-zinc-600">
                    Hold an RFID card near the reader
                  </p>
                </div>
              )}
            </div>
          </div>
        )}

        {activeTab === "classes" && (
          <div className="border border-zinc-800/50 bg-zinc-900/30 p-6">
            <ClassManager />
          </div>
        )}

        {activeTab === "reports" && (
          <div className="border border-zinc-800/50 bg-zinc-900/30 p-6">
            <AttendanceReport />
          </div>
        )}

        <footer className="mt-12 text-center">
          <p className="text-xs text-zinc-600">
            RFID Attendance System • {new Date().getFullYear()}
          </p>
        </footer>
      </div>
    </div>
  );
}