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

const emptyClass = {
  name: "",
  startTime: "",
  endTime: "",
  monday: false,
  tuesday: false,
  wednesday: false,
  thursday: false,
  friday: false,
  saturday: false,
  sunday: false,
};

const days = [
  { key: "monday", label: "M" },
  { key: "tuesday", label: "T" },
  { key: "wednesday", label: "W" },
  { key: "thursday", label: "Th" },
  { key: "friday", label: "F" },
  { key: "saturday", label: "Sa" },
  { key: "sunday", label: "Su" },
];

export default function ClassManager() {
  const [classes, setClasses] = useState<Class[]>([]);
  const [loading, setLoading] = useState(true);
  const [editing, setEditing] = useState<number | null>(null);
  const [creating, setCreating] = useState(false);
  const [formData, setFormData] = useState(emptyClass);
  const [saving, setSaving] = useState(false);

  const API_URL = import.meta.env.VITE_API_URL || "";

  useEffect(() => {
    fetchClasses();
  }, []);

  const fetchClasses = async () => {
    try {
      const response = await fetch(`${API_URL}/api/classes/`);
      const data = await response.json();
      setClasses(data.data || []);
    } catch (e) {
      console.error("Failed to fetch classes:", e);
    } finally {
      setLoading(false);
    }
  };

  const handleCreate = async () => {
    setSaving(true);
    try {
      const response = await fetch(`${API_URL}/api/classes/`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });
      if (response.ok) {
        setCreating(false);
        setFormData(emptyClass);
        fetchClasses();
      }
    } catch (e) {
      console.error("Failed to create class:", e);
    } finally {
      setSaving(false);
    }
  };

  const handleUpdate = async (id: number) => {
    setSaving(true);
    try {
      const response = await fetch(`${API_URL}/api/classes/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });
      if (response.ok) {
        setEditing(null);
        setFormData(emptyClass);
        fetchClasses();
      }
    } catch (e) {
      console.error("Failed to update class:", e);
    } finally {
      setSaving(false);
    }
  };

  const handleDelete = async (id: number) => {
    if (!confirm("Delete this class?")) return;
    try {
      const response = await fetch(`${API_URL}/api/classes/${id}`, {
        method: "DELETE",
      });
      if (response.ok) {
        fetchClasses();
      }
    } catch (e) {
      console.error("Failed to delete class:", e);
    }
  };

  const startEdit = (cls: Class) => {
    setEditing(cls.ID);
    setCreating(false);
    setFormData({
      name: cls.name,
      startTime: cls.startTime,
      endTime: cls.endTime,
      monday: cls.monday,
      tuesday: cls.tuesday,
      wednesday: cls.wednesday,
      thursday: cls.thursday,
      friday: cls.friday,
      saturday: cls.saturday,
      sunday: cls.sunday,
    });
  };

  const cancelEdit = () => {
    setEditing(null);
    setCreating(false);
    setFormData(emptyClass);
  };

  const startCreate = () => {
    setCreating(true);
    setEditing(null);
    setFormData(emptyClass);
  };

  const toggleDay = (day: string) => {
    setFormData((prev) => ({ ...prev, [day]: !prev[day as keyof typeof prev] }));
  };

  const updateField = (field: string, value: string) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
  };

  if (loading) {
    return (
      <div className="flex items-center gap-3 text-zinc-500">
        <div className="w-4 h-4 border-2 border-zinc-600 border-t-emerald-500 rounded-full animate-spin" />
        Loading classes...
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h2 className="text-lg font-medium text-white tracking-tight">Classes</h2>
          <p className="text-sm text-zinc-500 mt-1">{classes.length} registered</p>
        </div>
        {!creating && !editing && (
          <button
            onClick={startCreate}
            className="group flex items-center gap-2 bg-zinc-800 hover:bg-emerald-500 text-zinc-400 hover:text-black px-4 py-2 transition-all duration-200"
          >
            <svg
              className="w-4 h-4"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
            </svg>
            <span className="text-sm font-medium">Add Class</span>
          </button>
        )}
      </div>

      {creating && (
        <div className="border-l-4 border-emerald-500 bg-zinc-800/30 p-6 space-y-5">
          <div className="flex items-center gap-3 mb-2">
            <div className="w-2 h-2 bg-emerald-500 rounded-full animate-pulse" />
            <span className="text-xs uppercase tracking-widest text-zinc-500">New Class</span>
          </div>

          <div>
            <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-2">
              Class Name
            </label>
            <input
              type="text"
              value={formData.name}
              onChange={(e) => updateField("name", e.target.value)}
              className="w-full bg-zinc-800 border border-zinc-700 px-4 py-3 text-white placeholder-zinc-600 focus:outline-none focus:border-emerald-500 transition-colors"
              placeholder="e.g. CSC322"
            />
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-2">
                Start
              </label>
              <input
                type="time"
                value={formData.startTime}
                onChange={(e) => updateField("startTime", e.target.value)}
                className="w-full bg-zinc-800 border border-zinc-700 px-4 py-3 text-white focus:outline-none focus:border-emerald-500 transition-colors"
              />
            </div>
            <div>
              <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-2">
                End
              </label>
              <input
                type="time"
                value={formData.endTime}
                onChange={(e) => updateField("endTime", e.target.value)}
                className="w-full bg-zinc-800 border border-zinc-700 px-4 py-3 text-white focus:outline-none focus:border-emerald-500 transition-colors"
              />
            </div>
          </div>

          <div>
            <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-3">
              Schedule
            </label>
            <div className="flex gap-2">
              {days.map(({ key, label }) => (
                <button
                  key={key}
                  type="button"
                  onClick={() => toggleDay(key)}
                  className={`w-10 h-10 text-sm font-medium transition-all duration-200 ${
                    formData[key as keyof typeof formData]
                      ? "bg-emerald-500 text-black"
                      : "bg-zinc-800 text-zinc-500 hover:bg-zinc-700 hover:text-zinc-300"
                  }`}
                >
                  {label}
                </button>
              ))}
            </div>
          </div>

          <div className="flex gap-3 pt-2">
            <button
              onClick={handleCreate}
              disabled={saving || !formData.name}
              className="flex-1 bg-emerald-500 text-black font-medium py-3 hover:bg-emerald-400 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {saving ? "Saving..." : "Create Class"}
            </button>
            <button
              onClick={cancelEdit}
              className="px-6 py-3 bg-zinc-800 text-zinc-400 hover:bg-zinc-700 hover:text-white transition-colors"
            >
              Cancel
            </button>
          </div>
        </div>
      )}

      <div className="space-y-1">
        {classes.map((cls) => (
          <div key={cls.ID}>
            {editing === cls.ID ? (
              <div className="border-l-4 border-emerald-500 bg-zinc-800/30 p-6 space-y-5">
                <div className="flex items-center gap-3 mb-2">
                  <div className="w-2 h-2 bg-emerald-500 rounded-full animate-pulse" />
                  <span className="text-xs uppercase tracking-widest text-zinc-500">Editing</span>
                </div>

                <div>
                  <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-2">
                    Class Name
                  </label>
                  <input
                    type="text"
                    value={formData.name}
                    onChange={(e) => updateField("name", e.target.value)}
                    className="w-full bg-zinc-800 border border-zinc-700 px-4 py-3 text-white placeholder-zinc-600 focus:outline-none focus:border-emerald-500 transition-colors"
                    placeholder="e.g. CSC322"
                  />
                </div>

                <div className="grid grid-cols-2 gap-4">
                  <div>
                    <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-2">
                      Start
                    </label>
                    <input
                      type="time"
                      value={formData.startTime}
                      onChange={(e) => updateField("startTime", e.target.value)}
                      className="w-full bg-zinc-800 border border-zinc-700 px-4 py-3 text-white focus:outline-none focus:border-emerald-500 transition-colors"
                    />
                  </div>
                  <div>
                    <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-2">
                      End
                    </label>
                    <input
                      type="time"
                      value={formData.endTime}
                      onChange={(e) => updateField("endTime", e.target.value)}
                      className="w-full bg-zinc-800 border border-zinc-700 px-4 py-3 text-white focus:outline-none focus:border-emerald-500 transition-colors"
                    />
                  </div>
                </div>

                <div>
                  <label className="block text-xs uppercase tracking-wider text-zinc-500 mb-3">
                    Schedule
                  </label>
                  <div className="flex gap-2">
                    {days.map(({ key, label }) => (
                      <button
                        key={key}
                        type="button"
                        onClick={() => toggleDay(key)}
                        className={`w-10 h-10 text-sm font-medium transition-all duration-200 ${
                          formData[key as keyof typeof formData]
                            ? "bg-emerald-500 text-black"
                            : "bg-zinc-800 text-zinc-500 hover:bg-zinc-700 hover:text-zinc-300"
                        }`}
                      >
                        {label}
                      </button>
                    ))}
                  </div>
                </div>

                <div className="flex gap-3 pt-2">
                  <button
                    onClick={() => handleUpdate(cls.ID)}
                    disabled={saving || !formData.name}
                    className="flex-1 bg-emerald-500 text-black font-medium py-3 hover:bg-emerald-400 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    {saving ? "Saving..." : "Save Changes"}
                  </button>
                  <button
                    onClick={cancelEdit}
                    className="px-6 py-3 bg-zinc-800 text-zinc-400 hover:bg-zinc-700 hover:text-white transition-colors"
                  >
                    Cancel
                  </button>
                </div>
              </div>
            ) : (
              <div className="group bg-zinc-900/30 hover:bg-zinc-900/50 border border-zinc-800/50 hover:border-zinc-700/50 p-4 transition-all duration-200">
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-4">
                    <div className="w-1 h-12 bg-gradient-to-b from-emerald-500 to-emerald-700" />
                    <div>
                      <h3 className="font-medium text-white">{cls.name}</h3>
                      <div className="flex items-center gap-3 mt-1">
                        <span className="text-sm text-zinc-500">
                          {cls.startTime} → {cls.endTime}
                        </span>
                        <div className="flex gap-1">
                          {days.map(({ key, label }) => (
                            <span
                              key={key}
                              className={`w-5 h-5 flex items-center justify-center text-[10px] ${
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
                    </div>
                  </div>
                  <div className="flex gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                    <button
                      onClick={() => startEdit(cls)}
                      className="p-2 text-zinc-500 hover:text-white hover:bg-zinc-800 transition-colors"
                    >
                      <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                      </svg>
                    </button>
                    <button
                      onClick={() => handleDelete(cls.ID)}
                      className="p-2 text-zinc-500 hover:text-red-400 hover:bg-zinc-800 transition-colors"
                    >
                      <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                      </svg>
                    </button>
                  </div>
                </div>
              </div>
            )}
          </div>
        ))}
      </div>

      {classes.length === 0 && !creating && (
        <div className="text-center py-12 border border-dashed border-zinc-800">
          <div className="w-12 h-12 mx-auto mb-4 bg-zinc-900 flex items-center justify-center">
            <svg className="w-6 h-6 text-zinc-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
          </div>
          <p className="text-zinc-500 mb-4">No classes configured</p>
          <button
            onClick={startCreate}
            className="text-emerald-500 hover:text-emerald-400 text-sm font-medium"
          >
            Create your first class →
          </button>
        </div>
      )}
    </div>
  );
}