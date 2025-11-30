import { useState, useEffect } from "react";

interface Class {
  ID: number;
  Name: string;
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

  const API_URL = import.meta.env.VITE_API_URL || "";

  const fetchUser = async () => {
    setLoading(true);
    setError(null);

    try {
      const response = await fetch(`${API_URL}/api/users/card/${cardId}`);

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
      const response = await fetch("${API_URL}/api/classes");
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
      const userResponse = await fetch(`${API_URL}/api/users/${user.ID}`, {
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

      const classResponse = await fetch(`${API_URL}/api/users/${user.ID}/classes`, {
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
    return <div className="mt-4 text-gray-500">Loading user info...</div>;
  }

  if (error) {
    return (
      <div className="mt-4">
        <p className="text-red-500 mb-2">{error}</p>
        <button
          onClick={() => (window.location.href = `/register?cardId=${cardId}`)}
          className="bg-green-500 text-white px-4 py-2 rounded text-sm"
        >
          Register New User
        </button>
      </div>
    );
  }

  if (!user) return null;

  return (
    <div className="mt-4 space-y-4">
      <div className="border-t pt-4">
        {editing ? (
          <div className="space-y-3">
            <div>
              <label className="block text-sm text-gray-600 mb-1">Name</label>
              <input
                type="text"
                value={editName}
                onChange={(e) => setEditName(e.target.value)}
                className="border rounded px-3 py-2 w-full"
              />
            </div>
            <div>
              <label className="block text-sm text-gray-600 mb-1">Age</label>
              <input
                type="number"
                value={editAge}
                onChange={(e) => setEditAge(parseInt(e.target.value) || 0)}
                className="border rounded px-3 py-2 w-full"
              />
            </div>
            <div>
              <label className="block text-sm text-gray-600 mb-1">Classes</label>
              <div className="space-y-2">
                {allClasses.map((cls) => (
                  <label key={cls.ID} className="flex items-center gap-2">
                    <input
                      type="checkbox"
                      checked={editClassIds.includes(cls.ID)}
                      onChange={() => toggleClass(cls.ID)}
                      className="w-4 h-4"
                    />
                    <span>{cls.Name}</span>
                  </label>
                ))}
              </div>
            </div>
            <div className="flex gap-2">
              <button
                onClick={handleSave}
                disabled={saving}
                className="bg-green-500 text-white px-4 py-2 rounded text-sm disabled:bg-gray-400"
              >
                {saving ? "Saving..." : "Save"}
              </button>
              <button
                onClick={handleCancel}
                className="bg-gray-300 text-gray-700 px-4 py-2 rounded text-sm"
              >
                Cancel
              </button>
            </div>
          </div>
        ) : (
          <div>
            <h3 className="font-bold text-lg">{user.Name}</h3>
            <p className="text-gray-600 text-sm">Age: {user.Age}</p>
            <button
              onClick={() => setEditing(true)}
              className="mt-2 text-blue-500 text-sm underline"
            >
              Edit
            </button>
          </div>
        )}
      </div>

      <div>
        <h4 className="font-semibold mb-2">
          Registered Classes ({user.Classes?.length || 0})
        </h4>

        {!user.Classes || user.Classes.length === 0 ? (
          <p className="text-gray-500 text-sm">No classes registered</p>
        ) : (
          <ul className="space-y-2">
            {user.Classes.map((cls) => (
              <li
                key={cls.ID}
                className="bg-gray-50 p-3 rounded border text-sm"
              >
                <p className="font-medium">{cls.Name}</p>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}