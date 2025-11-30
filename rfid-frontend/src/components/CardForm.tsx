import { useState, useEffect } from "react";

interface Class {
  ID: number;
  name: string;
  startTime: string;
  endTime: string;
  monday: boolean;
  tuesday: boolean;
  wednesday: boolean;
  thursday: boolean;
  friday: boolean;
  saturday: boolean;
  sunday: boolean;
}

interface UserData {
  ID: number;
  Name: string;
  Age: number;
  CardID: string;
  Classes: Class[];
}

interface CardFormProps {
  cardId: string;
}

const days = [
  { key: "monday", label: "M" },
  { key: "tuesday", label: "T" },
  { key: "wednesday", label: "W" },
  { key: "thursday", label: "Th" },
  { key: "friday", label: "F" },
  { key: "saturday", label: "Sa" },
  { key: "sunday", label: "Su" },
];

export default function CardForm({ cardId }: CardFormProps) {
  const [user, setUser] = useState<UserData | null>(null);
  const [allClasses, setAllClasses] = useState<Class[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [editing, setEditing] = useState(false);
  const [saving, setSaving] = useState(false);

  const [editName, setEditName] = useState("");
  const [editAge, setEditAge] = useState(0);
  const [editClassIds, setEditClassIds] = useState<number[]>([]);

  useEffect(() => {
    fetchUser();
    fetchClasses();
  }, [cardId]);

  const fetchUser = async () => {
    setLoading(true);
    setError(null);

    try {
      const response = await fetch(`/api/users/card/${cardId}`);

      if (!response.ok) {
        if (response.status === 404) {
          setError("User not found - card not registered");
        } else {
          throw new Error(`HTTP ${response.status}`);
        }
        setUser(null);
        return;
      }

      const data: UserData = await response.json();
      setUser(data);
      setEditName(data.Name);
      setEditAge(data.Age);
      setEditClassIds(data.Classes?.map((c) => c.ID) || []);
    } catch (e) {
      const msg = e instanceof Error ? e.message : "Failed to fetch user";
      setError(msg);
      setUser(null);
    } finally {
      setLoading(false);
    }
  };

  const fetchClasses = async () => {
    try {
      const response = await fetch("/api/classes/");
      const data = await response.json();
      setAllClasses(data.data || []);
    } catch (e) {
      console.error("Failed to fetch classes:", e);
    }
  };

  const toggleClass = (classId: number) => {
    setEditClassIds((prev) =>
      prev.includes(classId)
        ? prev.filter((id) => id !== classId)
        : [...prev, classId]
    );
  };

  const handleSave = async () => {
    if (!user) return;
    setSaving(true);

    try {
      const userResponse = await fetch(`/api/users/${user.ID}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          Name: editName,
          Age: editAge,
        }),
      });

      if (!userResponse.ok) {
        throw new Error(`HTTP ${userResponse.status}`);
      }

      const classResponse = await fetch(`/api/users/${user.ID}/classes`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          classIds: editClassIds,
        }),
      });

      if (!classResponse.ok) {
        throw new Error(`HTTP ${classResponse.status}`);
      }

      const updatedUser = await classResponse.json();
      setUser(updatedUser);
      setEditing(false);
    } catch (e) {
      const msg = e instanceof Error ? e.message : "Failed to save";
      alert("Error saving: " + msg);
    } finally {
      setSaving(false);
    }
  };

  const handleCancel = () => {
    if (user) {
      setEditName(user.Name);
      setEditAge(user.Age);
      setEditClassIds(user.Classes?.map((c) => c.ID) || []);
    }
    setEditing(false);
  };

  if (loading) {
    return (
      <div className="flex items-center gap-3 text-zinc-500 py-8">
        <div className="w-4 h-4 border-2 border-zinc-600 border-t-emerald-500 rounded-full animate-spin" />
        Loading user info...
      </div>
    );
  }

  if (error) {
    return (
      <div className="py-8">
        <div className="flex items-center gap-3 text-amber-500 mb-4">
          <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
          <span>{error}</span>
        </div>
        <button
          onClick={() => (window.location.href = `/register?cardId=${cardId}`)}
          className="bg-emerald-500 text-black font-medium px-6 py-3 hover:bg-emerald-400 transition-colors"
        >
          Register New User
        </button>
      </div>
    );
  }

  if (!user) return null;

  return (
    <div className="space-y-6">
      {editing ? (
        <div className="border-l-4 border-emerald-500 bg-zinc-800/30 p-6 space-y-5">
          <div className="flex items-center gap-3 mb-2">
            <div className="w-2 h-2 bg-emerald-500 rounded-full animate-pulse" />
            <span className="text-xs uppercase tracking-widest text-zinc-500">Editing User</span>
          </div>

          <div>
            <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-2">Name</label>
            <input
              type="text"
              value={editName}
              onChange={(e) => setEditName(e.target.value)}
              className="w-full bg-zinc-800 border border-zinc-700 px-4 py-3 text-white placeholder-zinc-600 focus:outline-none focus:border-emerald-500 transition-colors"
            />
          </div>

          <div>
            <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-2">Age</label>
            <input
              type="number"
              value={editAge}
              onChange={(e) => setEditAge(parseInt(e.target.value) || 0)}
              className="w-full bg-zinc-800 border border-zinc-700 px-4 py-3 text-white placeholder-zinc-600 focus:outline-none focus:border-emerald-500 transition-colors"
            />
          </div>

          <div>
            <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-3">Classes</label>
            <div className="space-y-2">
              {allClasses.map((cls) => (
                <label
                  key={cls.ID}
                  className={`flex items-center gap-3 p-3 cursor-pointer transition-colors ${
                    editClassIds.includes(cls.ID)
                      ? "bg-emerald-500/10 border border-emerald-500/30"
                      : "bg-zinc-800/50 border border-zinc-700/50 hover:border-zinc-600"
                  }`}
                >
                  <input
                    type="checkbox"
                    checked={editClassIds.includes(cls.ID)}
                    onChange={() => toggleClass(cls.ID)}
                    className="sr-only"
                  />
                  <div
                    className={`w-5 h-5 border flex items-center justify-center transition-colors ${
                      editClassIds.includes(cls.ID)
                        ? "bg-emerald-500 border-emerald-500"
                        : "border-zinc-600"
                    }`}
                  >
                    {editClassIds.includes(cls.ID) && (
                      <svg className="w-3 h-3 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={3} d="M5 13l4 4L19 7" />
                      </svg>
                    )}
                  </div>
                  <div className="flex-1">
                    <span className="text-white">{cls.name}</span>
                    <span className="text-zinc-500 text-sm ml-2">
                      {cls.startTime} - {cls.endTime}
                    </span>
                  </div>
                </label>
              ))}
            </div>
          </div>

          <div className="flex gap-3 pt-2">
            <button
              onClick={handleSave}
              disabled={saving}
              className="flex-1 bg-emerald-500 text-black font-medium py-3 hover:bg-emerald-400 transition-colors disabled:opacity-50"
            >
              {saving ? "Saving..." : "Save Changes"}
            </button>
            <button
              onClick={handleCancel}
              className="px-6 py-3 bg-zinc-800 text-zinc-400 hover:bg-zinc-700 hover:text-white transition-colors"
            >
              Cancel
            </button>
          </div>
        </div>
      ) : (
        <>
          <div className="flex items-start justify-between">
            <div className="flex items-center gap-4">
              <div className="w-12 h-12 bg-gradient-to-br from-emerald-500 to-emerald-700 flex items-center justify-center text-black font-bold text-lg">
                {user.Name.charAt(0).toUpperCase()}
              </div>
              <div>
                <h3 className="text-xl font-medium text-white">{user.Name}</h3>
                <p className="text-zinc-500 text-sm">Age: {user.Age}</p>
              </div>
            </div>
            <button
              onClick={() => setEditing(true)}
              className="flex items-center gap-2 text-zinc-500 hover:text-white transition-colors"
            >
              <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
              </svg>
              <span className="text-sm">Edit</span>
            </button>
          </div>

          <div>
            <h4 className="text-xs uppercase tracking-wider text-zinc-500 mb-3">
              Enrolled Classes ({user.Classes?.length || 0})
            </h4>

            {!user.Classes || user.Classes.length === 0 ? (
              <p className="text-zinc-600 text-sm py-4">No classes enrolled</p>
            ) : (
              <div className="space-y-2">
                {user.Classes.map((cls) => (
                  <div
                    key={cls.ID}
                    className="flex items-center justify-between bg-zinc-800/30 border border-zinc-800/50 p-4"
                  >
                    <div className="flex items-center gap-3">
                      <div className="w-1 h-8 bg-emerald-500" />
                      <div>
                        <p className="font-medium text-white">{cls.name}</p>
                        <p className="text-sm text-zinc-500">
                          {cls.startTime} - {cls.endTime}
                        </p>
                      </div>
                    </div>
                    <div className="flex gap-1">
                      {days.map(({ key, label }) => (
                        <span
                          key={key}
                          className={`w-6 h-6 flex items-center justify-center text-[10px] ${
                            cls[key as keyof Class]
                              ? "bg-emerald-500/20 text-emerald-400"
                              : "bg-zinc-800/50 text-zinc-600"
                          }`}
                        >
                          {label}
                        </span>
                      ))}
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>
        </>
      )}
    </div>
  );
}