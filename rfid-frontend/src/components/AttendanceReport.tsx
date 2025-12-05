import { useState, useEffect } from "react";

interface Class {
  ID: number;
  name: string;
}

interface UserStats {
  id: number;
  name: string;
  cardId: string;
  attendedCount: number;
  totalSessions: number;
  attendanceRate: number;
}

interface ReportData {
  class: {
    id: number;
    name: string;
    totalSessions: number;
    averageRate: number;
    enrolledCount: number;
  };
  users: UserStats[];
}

const DATE_RANGES = [
  { label: "Last 7 days", days: 7 },
  { label: "Last 30 days", days: 30 },
  { label: "This semester", days: 120 },
  { label: "All time", days: 0 },
];

export default function AttendanceReport() {
  const [classes, setClasses] = useState<Class[]>([]);
  const [selectedClassId, setSelectedClassId] = useState<number | null>(null);
  const [selectedRange, setSelectedRange] = useState(30);
  const [report, setReport] = useState<ReportData | null>(null);
  const [loading, setLoading] = useState(false);
  const [loadingClasses, setLoadingClasses] = useState(true);

  useEffect(() => {
    fetchClasses();
  }, []);

  useEffect(() => {
    if (selectedClassId) {
      fetchReport();
    }
  }, [selectedClassId, selectedRange]);

  const fetchClasses = async () => {
    try {
      const response = await fetch("/api/classes/");
      const data = await response.json();
      setClasses(data.data || []);
    } catch (e) {
      console.error("Failed to fetch classes:", e);
    } finally {
      setLoadingClasses(false);
    }
  };

  const fetchReport = async () => {
    if (!selectedClassId) return;
    setLoading(true);

    let url = `/api/attendance/report/${selectedClassId}`;
    if (selectedRange > 0) {
      const endDate = new Date().toISOString().split("T")[0];
      const startDate = new Date(Date.now() - selectedRange * 24 * 60 * 60 * 1000)
        .toISOString()
        .split("T")[0];
      url += `?startDate=${startDate}&endDate=${endDate}`;
    }

    try {
      const response = await fetch(url);
      const data = await response.json();
      setReport(data);
    } catch (e) {
      console.error("Failed to fetch report:", e);
    } finally {
      setLoading(false);
    }
  };

  const getAttendanceColor = (rate: number) => {
    if (rate >= 80) return "text-emerald-400";
    if (rate >= 60) return "text-yellow-400";
    return "text-red-400";
  };

  const getBarColor = (rate: number) => {
    if (rate >= 80) return "bg-emerald-500";
    if (rate >= 60) return "bg-yellow-500";
    return "bg-red-500";
  };

  if (loadingClasses) {
    return (
      <div className="flex items-center justify-center py-16">
        <div className="w-5 h-5 border-2 border-zinc-700 border-t-emerald-500 rounded-full animate-spin" />
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row gap-4">
        <div className="flex-1">
          <label className="block text-xs text-zinc-500 uppercase tracking-wider mb-2">
            Select Class
          </label>
          <select
            value={selectedClassId || ""}
            onChange={(e) => setSelectedClassId(Number(e.target.value) || null)}
            className="w-full bg-zinc-900 border border-zinc-800 text-white px-4 py-3 focus:outline-none focus:border-zinc-600 transition-colors"
          >
            <option value="">Choose a class...</option>
            {classes.map((cls) => (
              <option key={cls.ID} value={cls.ID}>
                {cls.name}
              </option>
            ))}
          </select>
        </div>

        <div className="sm:w-48">
          <label className="block text-xs text-zinc-500 uppercase tracking-wider mb-2">
            Date Range
          </label>
          <select
            value={selectedRange}
            onChange={(e) => setSelectedRange(Number(e.target.value))}
            className="w-full bg-zinc-900 border border-zinc-800 text-white px-4 py-3 focus:outline-none focus:border-zinc-600 transition-colors"
          >
            {DATE_RANGES.map((range) => (
              <option key={range.days} value={range.days}>
                {range.label}
              </option>
            ))}
          </select>
        </div>
      </div>

      {!selectedClassId && (
        <div className="text-center py-16 border border-dashed border-zinc-800">
          <div className="w-16 h-16 mx-auto mb-6 border-2 border-dashed border-zinc-800 rounded-full flex items-center justify-center">
            <svg className="w-8 h-8 text-zinc-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
          <p className="text-zinc-500 mb-2">Select a class to view attendance report</p>
          <p className="text-xs text-zinc-600">Choose from the dropdown above</p>
        </div>
      )}

      {loading && (
        <div className="flex items-center justify-center py-16">
          <div className="w-5 h-5 border-2 border-zinc-700 border-t-emerald-500 rounded-full animate-spin" />
        </div>
      )}

      {!loading && report && (
        <>
          <div className="grid grid-cols-3 gap-4">
            <div className="bg-zinc-900/50 border border-zinc-800/50 p-4">
              <p className="text-xs text-zinc-500 uppercase tracking-wider mb-1">Class Average</p>
              <p className={`text-2xl font-light ${getAttendanceColor(report.class.averageRate)}`}>
                {report.class.averageRate.toFixed(1)}%
              </p>
            </div>
            <div className="bg-zinc-900/50 border border-zinc-800/50 p-4">
              <p className="text-xs text-zinc-500 uppercase tracking-wider mb-1">Total Sessions</p>
              <p className="text-2xl font-light text-white">{report.class.totalSessions}</p>
            </div>
            <div className="bg-zinc-900/50 border border-zinc-800/50 p-4">
              <p className="text-xs text-zinc-500 uppercase tracking-wider mb-1">Enrolled</p>
              <p className="text-2xl font-light text-white">{report.class.enrolledCount}</p>
            </div>
          </div>

          <div className="border border-zinc-800/50">
            <div className="bg-zinc-900/50 px-4 py-3 border-b border-zinc-800/50">
              <h3 className="text-sm font-medium text-zinc-300">Student Attendance</h3>
            </div>

            {report.users && report.users.length > 0 ? (
              <div className="divide-y divide-zinc-800/50">
                {report.users.map((user) => (
                  <div key={user.id} className="p-4 hover:bg-zinc-900/30 transition-colors">
                    <div className="flex items-center justify-between mb-3">
                      <div>
                        <p className="font-medium text-white">{user.name}</p>
                        <p className="text-xs text-zinc-600 font-mono">{user.cardId}</p>
                      </div>
                      <div className="text-right">
                        <p className={`text-lg font-light ${getAttendanceColor(user.attendanceRate)}`}>
                          {user.attendanceRate.toFixed(1)}%
                        </p>
                        <p className="text-xs text-zinc-500">
                          {user.attendedCount} / {user.totalSessions} sessions
                        </p>
                      </div>
                    </div>
                    <div className="h-1.5 bg-zinc-800 overflow-hidden">
                      <div
                        className={`h-full ${getBarColor(user.attendanceRate)} transition-all duration-500`}
                        style={{ width: `${Math.min(user.attendanceRate, 100)}%` }}
                      />
                    </div>
                  </div>
                ))}
              </div>
            ) : (
              <div className="p-8 text-center text-zinc-500">
                No students enrolled in this class
              </div>
            )}
          </div>
        </>
      )}
    </div>
  );
}